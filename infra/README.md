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

This project also currently assumes you have your `MOMENTO_AUTH_TOKEN` stored in AWS secrets manager with the
name `/momento/authToken` as a simple string secret.

## Useful commands

* `npm run build`   compile typescript to js
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
