---
title: 'Be attention to replace a key pair for EC2'
published: true
description:
tags: aws, ec2, ssh
---

## Introduction

I will share points to note about replace a key pair for EC2.

## About key pair

When you launch the instance, you select the key pair to connect with SSH. ![image1](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/dn2ypob1sn2x6ugmpne0.png)

EC2 instance cannot allow to be connected to with SSH by password authentication in initial settings.

```bash
# To disable tunneled clear text passwords, change to no here!
#PasswordAuthentication yes
#PermitEmptyPasswords no
PasswordAuthentication no
```

Once you select a key pair and launch EC2 instance, you can connect to it with public key authentication.

```bash
ssh -i ./NewWindows.pem ec2-user@18.182.24.156
Last login: Mon Aug  2 13:36:42 2021

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
```

## Replace a key pair

You create an Amazon Machine Image(AMI) based on running EC2 instance and you launch EC2 instance from that AMI. Then, select a key pair. ![Image2](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/5rdd7g3irkdhlotloxvd.png)

There is a case that you need to replace a key pair attached EC2 instance because of security, review of operations and so on.

In this case, you change a new key pair from existing key pair. Also, you can connect to EC2 instance with another secret key.

```bash
ssh -i .\tepkey.pem ec2-user@18.183.44.42
Last login: Mon Aug  2 13:45:37 2021

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
```

At first, I expected that only new public key exists in the file, `~/.ssh/authorized_keys`. But, it was wrong and an old public key existed there, too.

```bash
ssh-rsa ~~~~~~~~~~~~~~~~~~~XXXX~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ NewWindows
ssh-rsa ~~~~~~~~~~~~~~~~~~~XXXX~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ tepkey
```

Actually, I confirm that I can connect to EC2 instance with SSH by the former secret key.

```bash
ssh -i .\NewWindows.pem ec2-user@18.183.44.42
Last login: Mon Aug  2 15:07:23 2021

       __|  __|_  )
       _|  (     /   Amazon Linux 2 AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-2/
```

According to the AWS documentation, you have to remove an old public key in `~/.ssh/authorized_keys` manually when you replace a key pair.

https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#replacing-key-pair

## Conclusion

I research replace a key pair for EC2 instance. I'm surprised at the specification and concerned about security🤔

If you disclosure an old secret key which is not used and not managed, your EC2 instances may be illegally accessed.

How does everyone do an operation? Let me hear your what you think.

## Original

https://zenn.dev/yuta28/articles/ec2-keypair-replace
