import {App, Fn, Stack} from 'aws-cdk-lib'
import {
    AwsLogDriver,
    Cluster,
    Compatibility,
    ContainerImage,
    FargateService,
    Secret,
    TaskDefinition
} from 'aws-cdk-lib/aws-ecs'
import {resolve} from "path";
import {Vpc} from "aws-cdk-lib/aws-ec2";
import {Platform} from "aws-cdk-lib/aws-ecr-assets";
import * as secretsmanager from "aws-cdk-lib/aws-secretsmanager";
import {
    Effect,
    ManagedPolicy,
    PolicyDocument,
    PolicyStatement,
    Role,
    ServicePrincipal
} from "aws-cdk-lib/aws-iam";

const TARGET_LAMBDA_NAME = "foo";
const CACHE_NAME = "default-cache";
const TOPIC_NAME = "test-topic";

const app = new App();
const stack = new Stack(app, 'topic-lambda-invoker');

const momentoAuthToken = secretsmanager.Secret.fromSecretNameV2(stack,
    "momento-auth-token-secret",
    "/momento/authToken"
);

const vpc = new Vpc(stack, 'topic-invoker-vpc', {maxAzs: 2});

const fargateCluster = new Cluster(stack, 'fargate-cluster', {vpc});

const topicLambdaInvokerTaskDef = new TaskDefinition(stack, "invoker-task", {
    compatibility: Compatibility.FARGATE,
    cpu: "256",
    memoryMiB: "512",
    taskRole: new Role(stack, 'task-role', {
        assumedBy: new ServicePrincipal("ecs-tasks.amazonaws.com"),
        inlinePolicies: {
            "lambdaInvokePolicy": new PolicyDocument({
                statements: [
                    new PolicyStatement({
                        actions: ['lambda:InvokeFunction'],
                        effect: Effect.ALLOW,
                        resources: [Fn.sub(
                            `arn:aws:lambda:$\{AWS::Region}:$\{AWS::AccountId}:function:${TARGET_LAMBDA_NAME}*`
                        )]
                    })
                ]
            })
        },
    })
});

const logging = new AwsLogDriver({
    streamPrefix: "topic-lambda-invoker",
})

topicLambdaInvokerTaskDef.addContainer("invoker-container", {
    containerName: 'topic-lambda-invoker',
    image: ContainerImage.fromAsset(resolve(__dirname, '..'), {platform: Platform.LINUX_AMD64}),
    environment: {
        "CACHE_NAME": CACHE_NAME,
        "TOPIC_NAME": TOPIC_NAME,
        "FUNCTION_TARGET_NAME": TARGET_LAMBDA_NAME,
    },
    logging,
    secrets: {
        "MOMENTO_AUTH_TOKEN": Secret.fromSecretsManager(momentoAuthToken)
    }
});

new FargateService(stack, "lambda-invoker-fargate-service", {
    cluster: fargateCluster,
    taskDefinition: topicLambdaInvokerTaskDef,
    desiredCount: 1
});

app.synth();
