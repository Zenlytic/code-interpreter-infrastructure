# A Note From Zenlytic

Thanks to the E2B team for their hard work on this project. Our primary goal in forking [e2b-dev/infra](https://github.com/e2b-dev/infra) is to adapt it to our own hosting requirements, namely by enabling deployment on AWS. We release our derived work under the Apache 2.0 license.

# E2B Infrastructure

[E2B](https://e2b.dev) is an open-source infrastructure for AI code interpreting. In our main repository [e2b-dev/e2b](https://github.com/e2b-dev/E2B) we are giving you SDKs and CLI to customize and manage environments and run your AI agents in the cloud.

This repository contains the infrastructure that powers the E2B platform.

## Self-hosting

Read the [self-hosting guide](./self-host.md) to learn how to set up the infrastructure on your own. The infrastructure is deployed using Terraform. The process is not perfect yet, but we are working on it.

Supported cloud providers:
- 🟢 GCP
- 🚧 AWS
- [ ] Azure
- [ ] General linux machine
