# Welcome to your CDK TypeScript project

This is reference CDK code for deploying the `topic-lambda-invoker` appliance.

To change the cache, topic, or target lambda it will subscribe to and then invoke
please edit the following config files at the top of `main.ts` and then deploy the 
project into your account.

```typescript
const TARGET_LAMBDA_NAME = "foo";
const CACHE_NAME = "default-cache";
const TOPIC_NAME = "test-topic";
```

## Useful commands

* `npm run build`   compile typescript to js
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
