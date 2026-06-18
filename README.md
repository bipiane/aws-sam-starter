# AWS SAM Starter

🚀 AWS SAM Starter is a hands-on learning project that demonstrates how to build, deploy, and operate serverless
applications on AWS using AWS SAM.

## Requirements

- AWS SAM CLI
- Go 1.22 o superior

## Commands

```bash
# Build
sam build
```

```bash
# Run local server
sam local start-api
curl http://127.0.0.1:3000/api/qr
```

```bash
# Invoke local lambda
sam local invoke QrFunction --event events/apigateway-list-qr.json
```

```bash
# Deploy
sam build && sam deploy --profile your-aws-profile
```
