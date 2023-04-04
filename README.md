<img src="https://docs.momentohq.com/img/logo.svg" alt="logo" width="400"/>

[![project status](https://momentohq.github.io/standards-and-practices/badges/project-status-official.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)
[![project stability](https://momentohq.github.io/standards-and-practices/badges/project-stability-experimental.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)


# topic-lambda-invoker

A container appliance that can be deployed as a standalone fargate ecs service in your AWS account to subscribe 
to a momento topic and trigger lambda functions asynchronously in your account.

When deployed you will end up with an architecture like this.

![image](./img/TopicInvoker.png)
