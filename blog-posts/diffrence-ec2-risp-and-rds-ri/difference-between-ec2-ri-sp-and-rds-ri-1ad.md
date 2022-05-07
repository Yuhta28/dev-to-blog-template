---
title: Difference between EC2 RI/SP and RDS RI
published: true
description:
tags: aws, ec2, rds
---

## Introduction

Reserved Instances(RI) or Savings Plans(SP) of EC2 and Reserved Instances(RI) of RDS are for Cost Optimization, which is one of 5 pillar of the AWS Well-Architected Framework.

However, there are subtle differences in their specifications that made me confused. I summarize the difference of them and share post.

## Feature

At first, I introduce these basics features.

### About RI

Reserved Instances(RI) provide a significant discount compared to On-Demand pricing and provide a capacity reservation when used in a specific Availability Zone. If you would like to use some instances of specific instance type, you can reduce the cost to select RI.

And also, you pay for the entire Reserved Instance term with one upfront payment and get the best effective hourly price when compared to running the same instance on an On-Demand basis.

### About SP

Savings Plans also provides a discount compared to On-Demand pricing. RI is applied to discount regarding specific instance type or availability zone. On the other hand, SP is a flexible pricing model in exchange for a specific usage commitment. The discount rate of SP is smaller than one of RI, but you need not to specify instance type and can make more flexible cost optimization plans.

## Cautions for Purchase RI or SP

Both RI and SP are for cost reduction, but you pay attention to purchase. That is an order date. You can purchase EC2 RI to specify order date. ![image](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/2qxkxljyu2cmbjvrafq1.png) Because of this future, you don't have to work to launch RI on holiday and you can purchase it in advance.

Moreover, in case of EC2 SP, you can specify order time. ![image](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/5yb0xlqgafjdydmt2h89.png) Because of this future, you can apply SP at midnight local time.

However, you cannot purchase RDS RI in advance. ![Alt Text](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/f0neutysqhevknanb9lb.png)

Unfortunately RDS RI doesn't hove confirmation to order steps. You click submit button, and purchase is confirmed. I would like to purchase RDS RI in advance and made a mistake.

Let me show you below graph.

|                    | EC2 RI | EC2 SP | RDS RI |
| ------------------ | ------ | ------ | ------ |
| specify order date | ✓      | ✓      |        |
| specify order time |        | ✓      |        |

## Digression

To tell you the truth, in case of EC2 RI, you can specify order time with AWS CLI.

```bash
purchase-reserved-instances-offering
--instance-count <value> \
--reserved-instances-offering-id <value> \
--purchase-time <value>
```

## Conclusion

I share post about difference between EC2 RI/SP and RDS RI on order timing. There are subtle differences in their specifications that made me confused. Please be careful with these differences.

## Original

https://zenn.dev/yuta28/articles/sp-attention https://zenn.dev/yuta28/articles/rds-attention

## References

https://aws.amazon.com/ec2/pricing/reserved-instances/?nc1=h_ls https://aws.amazon.com/rds/reserved-instances/?nc1=h_ls https://aws.amazon.com/savingsplans/ https://awscli.amazonaws.com/v2/documentation/api/latest/reference/ec2/purchase-reserved-instances-offering.html
