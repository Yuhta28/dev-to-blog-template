---
title: GitHub Actions CI/CD without AWS Credential information
published: true
description:
tags: terraform, aws, github
---

## Introduction

October 2021, GitHub [announced](https://github.blog/changelog/2021-10-27-github-actions-secure-cloud-deployments-with-openid-connect/) that GitHub Actions supports OpenID Connect (OIDC) for secure deployments to cloud, which uses short-lived tokens that are automatically rotated for each deployment.

This feature helps me build CI/CD on GitHub Actions by AWS without AWS Credential such as IAM access key ID or secret access key.

I'll try to implement CI/CD for building EC2 instances by incorporating this feature into CI/CD using a combination of Terraform and GitHub Actions, and passing the IAM role to GitHub Actions.

## GitHub Actions CI/CD with terraform

Previously, it is necessary to register IAM access key ID and secret access key in the environment variables in advance when deploying AWS services with terraform and building CI/CD with GitHub Actions or Circle CI. ![GitHub Actions Setting](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/p0xa9fnqponrdxenq6qa.png)

```yaml
- name: Configure AWS Credentials
  uses: aws-actions/configure-aws-credentials@v1
  with:
    aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }} #access key ID
    aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }} #secret access key
```

There are some problems.

- Create IAM user to publish Access keys
- Re-register Access keys in case of AWS multi account operation

To solve them, I add OpenID Connect provider. I build CloudFormation stack template to refer to this [article](https://dev.classmethod.jp/articles/github-actions-without-permanent-credential/) and deploy IAM role & OIDC provider.

## Deploying IAM role & OIDC provider

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: 'IAM Role for GHA'

Parameters:
  RepoName:
    Type: String
    Default: Yuhta28/terraform-githubaction-ci

Resources:
  Role:
    Type: AWS::IAM::Role
    Properties:
      RoleName: ExampleGithubRole
      AssumeRolePolicyDocument:
        Statement:
          - Effect: Allow
            Action: sts:AssumeRoleWithWebIdentity
            Principal:
              Federated: !Ref GithubOidc
            Condition:
              StringLike:
                token.actions.githubusercontent.com:sub: !Sub repo:${RepoName}:*

  Policy:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: test-gha
      Roles:
        - !Ref Role
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - 'ec2:*'
              - 'sts:GetCallerIdentity'
              - 's3:*'
            Resource: '*'

  GithubOidc:
    Type: AWS::IAM::OIDCProvider
    Properties:
      Url: https://token.actions.githubusercontent.com
      ClientIdList: [sigstore]
      ThumbprintList: [a031c46782e6e6c662c2c87c76da9aa62ccabd8e]

Outputs:
  Role:
    Value: !GetAtt Role.Arn
```

After launch stack, IAM Role and OIDC are displayed. ![IAM Role](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/lpvzscustjh8hp4zd557.png) ![OIDC Provider](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/5l9daxy6o367r11zeikn.png)

## Create GHA Workflow

Create GitHub Actions workflow file. There is a template workflow file for terraform by HashiCorp and choose it. ![workflow](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/xr5kw9fe9fl5kch0exbx.png)

```yaml
name: 'Terraform'
on:
  push:
    branches:
      - main
  pull_request:

jobs:
  terraformCICD:
    runs-on: ubuntu-latest
    permissions:
      id-token: write
      contents: read
    steps:
      - run: sleep 5

      - uses: actions/checkout@v2

      - name: Configure AWS
        run: |
          export AWS_ROLE_ARN=arn:aws:iam::<AWS_AccountID>:role/ExampleGithubRole
          export AWS_WEB_IDENTITY_TOKEN_FILE=/tmp/awscreds
          export AWS_DEFAULT_REGION=ap-northeast-1

          echo AWS_WEB_IDENTITY_TOKEN_FILE=$AWS_WEB_IDENTITY_TOKEN_FILE >> $GITHUB_ENV
          echo AWS_ROLE_ARN=$AWS_ROLE_ARN >> $GITHUB_ENV
          echo AWS_DEFAULT_REGION=$AWS_DEFAULT_REGION >> $GITHUB_ENV
          curl -H "Authorization: bearer $ACTIONS_ID_TOKEN_REQUEST_TOKEN" "$ACTIONS_ID_TOKEN_REQUEST_URL&audience=sigstore" | jq -r '.value' > $AWS_WEB_IDENTITY_TOKEN_FILE

      # Install the latest version of Terraform CLI and configure the Terraform CLI configuration file with a Terraform Cloud user API token
      - name: Setup Terraform
        uses: aws-actions/configure-aws-credentials@master
        with:
          role-to-assume: '${{ env.AWS_ROLE_ARN }}'
          web-identity-token-file: '${{ env.AWS_WEB_IDENTITY_TOKEN_FILE }}'
          aws-region: '${{ env.AWS_DEFAULT_REGION }}'
          role-duration-seconds: 900
          role-session-name: GitHubActionsTerraformCICD

      # Checks that all Terraform configuration files adhere to a canonical format
      - name: Terraform Format
        run: terraform fmt -check -diff

      # Initialize a new or existing Terraform working directory by creating initial files, loading any remote state, downloading modules, etc.
      - name: Terraform Init
        run: terraform init

      - name: Terraform Validate
        run: terraform validate -no-color

      # Generates an execution plan for Terraform
      - name: Terraform Plan
        if: github.event_name == 'pull_request'
        run: terraform plan -no-color
        continue-on-error: true

        # On push to main, build or change infrastructure according to Terraform configuration files
        # Note: It is recommended to set up a required "strict" status check in your repository for "Terraform Cloud". See the documentation on "strict" required status checks for more information: https://help.github.com/en/github/administering-a-repository/types-of-required-status-checks
      - name: Terraform Apply
        if: github.ref == 'refs/heads/main' && github.event_name == 'push'
        run: terraform apply -auto-approve
```

It is important for that step `Configure AWS`.

I set three environment variables on GitHub.

- AWS_ROLE_ARN
- AWS_WEB_IDENTITY_TOKEN_FILE
- AWS_DEFAULT_REGION

### AWS_ROLE_ARN

Specify Assume Role I just create.

### AWS_WEB_IDENTITY_TOKEN_FILE

A path to Web ID token file.

### AWS_DEFAULT_REGION

A default region. I select Tokyo region,`ap-northeast-1`.

With `curl`, two parameters, `ACTIONS_ID_TOKEN_REQUEST_TOKEN` and `ACTIONS_ID_TOKEN_REQUEST_URL` is passed to Web ID token file.

After that, web ID token file is passed to `configure-aws-credentials`, one of actions.

```yaml
uses: aws-actions/configure-aws-credentials@master
with:
  role-to-assume: '${{ env.AWS_ROLE_ARN }}'
  web-identity-token-file: '${{ env.AWS_WEB_IDENTITY_TOKEN_FILE }}'
  aws-region: '${{ env.AWS_DEFAULT_REGION }}'
  role-duration-seconds: 900
  role-session-name: GitHubActionsTerraformCICD
```

Compared with workflow file described at the beginning of this section, IAM role and OIDC token are set instead of Access keys of environment variables. It is recommendation that `role-session-name` is set to examine the logs.

## Build Terraform

Create one EC2 instance and store `terraform.tfstate` in S3.

### Directory structure

```bash
|
├── main.tf
└── variables.tf
```

### main.tf

```hcl
terraform {
  backend "s3" {
    bucket = "specify bucket name"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"

  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_instance" "app_server_yuta" {
  ami                    = "ami-00cb500575fd9f9be"
  instance_type          = "t2.micro"
  vpc_security_group_ids = ["sg-XXXXXXXXXXXXXXXXX", "sg-XXXXXXXXXXXXXXXXX"]
  subnet_id              = "subnet-XXXXXXXXXXXXXXXXX"

  tags = {
    Name = var.instance_name
  }
}
```

### variables.tf

```hcl
variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
  default     = "Yuta-ServerInstance"
}
```

After coding terraform file, create branch and PR. GitHub Action launches. ![GitHub Action](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/2kr4y3cugsxuwrojv4or.png)

If `terraform plan`step is executed without any problems, merging will be possible. Once the main branch is merged, `terraform apply`step will run and the EC2 instance will be created. ![EC2](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/oev4wf6l3mlwlqnfj9hy.png)

## Conclusion

I try to build GitHub Actions CI/CD without AWS Credential information. This feature is new and there are few reference materials. However, I think the feature is very useful and would like that everyone use it.

## Original

https://zenn.dev/yuta28/articles/terraform-gha
