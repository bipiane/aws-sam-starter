# AWS SAM Starter

🚀 AWS SAM Starter is a hands-on learning project that demonstrates how to build, deploy, and operate serverless
applications on AWS using AWS SAM.

## Requirements

- AWS SAM CLI
- Go 1.22 o superior

## Run

```bash
# Build SAM
sam build
```

```bash
# Start Lambda at http://127.0.0.1:3001
sam local start-api

# Local
sam local invoke QrFunction --event events/apigateway-list-qr.json

# Remote
sam remote invoke QrFunction --profile your-aws-profile --region us-east-1 --stack-name aws-sam-starter --event-file './events/apigateway-list-qr.json'
```

## Deploy

```bash
# Deploy
sam build && sam deploy --profile your-aws-profile
```
