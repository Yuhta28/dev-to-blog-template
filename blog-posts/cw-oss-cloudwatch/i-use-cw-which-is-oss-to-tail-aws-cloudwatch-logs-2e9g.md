---
title: 'I use cw, which is OSS to tail AWS CloudWatch Logs'
published: true
description:
tags: aws, cloudwatch, opensource, tutorial
canonical_url: https://dev.to/yuta28/i-use-cw-which-is-oss-to-tail-aws-cloudwatch-logs-2e9g
---

## Introduction

I will share cw, OSS to help to tail CloudWatch Logs.

## Prerequisite knowledge

AWS CLI has a command which tails the logs collected CloudWatch Logs.

```bash
$ aws logs tail access_log
2021-07-22T07:45:55.422000+00:00 wordpress1 18.207.253.146 - - [22/Jul/2021:07:45:54 +0000] "GET /.env HTTP/1.1" 404 196 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
2021-07-22T07:45:59.986000+00:00 wordpress1 18.207.253.146 - - [22/Jul/2021:07:45:55 +0000] "POST / HTTP/1.1" 404 196 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
2021-07-22T07:49:31.689000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:49:31.689000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET /icons/apache_pb2.gif HTTP/1.1" 200 4234 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
```

If you would like to watch the logs in real time like `tail -f`, add the following option. `aws logs tail --follow <Log Group>`

```bash
$ aws logs tail --follow access_log
2021-07-22T07:45:55.422000+00:00 wordpress1 18.207.253.146 - - [22/Jul/2021:07:45:54 +0000] "GET /.env HTTP/1.1" 404 196 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
2021-07-22T07:45:59.986000+00:00 wordpress1 18.207.253.146 - - [22/Jul/2021:07:45:55 +0000] "POST / HTTP/1.1" 404 196 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
2021-07-22T07:49:31.689000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:49:31.689000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET /icons/apache_pb2.gif HTTP/1.1" 200 4234 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:49:31.940000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET /icons/poweredby.png HTTP/1.1" 200 3412 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:49:33.195000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:31 +0000] "GET /favicon.ico HTTP/1.1" 404 196 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:49:34.198000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:33 +0000] "GET / HTTP/1.1" 403 4890 "https://www.google.com/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"
2021-07-22T07:49:34.198000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:34 +0000] "GET /icons/poweredby.png HTTP/1.1" 200 3412 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"
2021-07-22T07:49:38.949000+00:00 wordpress2 27.121.46.196 - - [22/Jul/2021:07:49:34 +0000] "GET /icons/apache_pb2.gif HTTP/1.1" 200 4234 "http://13.231.136.114/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"
2021-07-22T07:54:52.911000+00:00 wordpress1 27.121.46.196 - - [22/Jul/2021:07:54:52 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
2021-07-22T07:54:52.911000+00:00 wordpress1 27.121.46.196 - - [22/Jul/2021:07:54:52 +0000] "GET /icons/apache_pb2.gif HTTP/1.1" 200 4234 "http://54.168.233.33/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
```

The output describes this.

`Timestamp + time zone streams messages`

It is useful that I can watch logs in real time to use `tail -f`, but I cannot specify each streams, and the access logs of multiple servers are mixed together.

I will share cw which solves this problem.

## What's cw?

`cw` is the best way to tail AWS CloudWatch Logs from your terminal. https://github.com/lucagrulla/cw

`cw` is a native executable targeting your OS, and not needed external dependencies such as `pip` and `npm`. Compared to `awslogs` which is famous helpful tool for CloudWatch Logs[^1], cw is written in golang and faster.

[^1]: https://github.com/jorgebastida/awslogs

Moreover, it is very helpful that is possible to specify multiple groups. To use cw, I will try to execute `tail -f` for each streams, and watch logs.

## Installation

You can install `cw` with brew command easily.

```bash
brew tap lucagrulla/tap
brew install cw
```

## Usage

If you have installed the AWS CLI locally, you have probably used registered your credentials to your local machine with `aws configure`, but if you have used aws cli within an EC2 instance with an IAM role, you will need to set the default region information in the file `~/.aws/config`.

```bash
[default]
region = ap-northeast-1
```

First, command `cw -h`.

```bash
$ cw -h
Usage: cw <command>

The best way to tail AWS Cloudwatch Logs from your terminal.

Flags:
  -h, --help               Show context-sensitive help.
      --endpoint=URL       The target AWS endpoint url. By default cw will use the default aws endpoints. NOTE: v4.0.0
                           dropped the flag short version.
      --profile=PROFILE    The target AWS profile. By default cw will use the default profile defined in the
                           .aws/credentials file. NOTE: v4.0.0 dropped the flag short version.
      --region=REGION      The target AWS region. By default cw will use the default region defined in the
                           .aws/credentials file. NOTE: v4.0.0 dropped the flag short version.
      --no-color           Disable coloured output.NOTE: v4.0.0 dropped the flag short version.
      --version            Print version information and quit

Commands:
  ls groups
    Show all groups.

  ls streams <group>
    Show all streams in a given log group.

  tail <groupName[:logStreamPrefix]> ...
    Tail log groups/streams.

Run "cw <command> --help" for more information on a command.
```

There is the introduction of basically how to use and sub command. Next, command `cw ls groups` and you see list of Log groups collected CloudWatch Logs.

```bash
$ cw ls groups
/aws/lambda/Yuta-rds-auto-stop
/ecs/first-run-task-definition
RDSOSMetrics
access_log
```

Command `cw tail -f` to watch CloudWatch Logs in real time.

```bash
$ cw tail access_log --follow --stream-name
wordpress2 - 27.121.46.196 - - [22/Jul/2021:09:19:33 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress2 - 27.121.46.196 - - [22/Jul/2021:09:20:24 +0000] "-" 408 - "-" "-"
wordpress2 - 27.121.46.196 - - [22/Jul/2021:09:20:29 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress2 - 27.121.46.196 - - [22/Jul/2021:09:20:29 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:20:48 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:20:49 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:20:49 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
```

I can do that like `tail -f`. The main theme is that watching logs for each streams. You add option `:stream` after group.

```bash
$ cw tail access_log:wordpress1 --follow --stream-name
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:24:53 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:24:53 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:24:54 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:24:54 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:24:55 +0000] "GET / HTTP/1.1" 403 4890 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
```

The stream, `wordpress1` is only displayed.

## Unresolved

I continue to output logs for a while, an error suddenly occurs and the program terminates. An error occurs a few minutes after command execution.

```bash
....

wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:29:45 +0000] "GET /wordpress/wp-admin/plugin-install.php?tab=popular HTTP/1.1" 200 126481 "http://54.168.233.33/wordpress/wp-admin/plugin-install.php" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
wordpress1 - 27.121.46.196 - - [22/Jul/2021:09:29:47 +0000] "GET /wordpress/wp-admin/plugin-install.php?tab=recommended HTTP/1.1" 200 127412 "http://54.168.233.33/wordpress/wp-admin/plugin-install.php?tab=popular" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"
operation error CloudWatch Logs: FilterLogEvents, exceeded maximum number of attempts, 3, https response error StatusCode: 400, RequestID: 69b981bb-0bcf-4263-951e-73aabf9ab379, api error ThrottlingException: Rate exceeded
```

According to message, it is the API error that said The maximum number of attempts has been exceeded.

I don't know details, so I have a issue for developer. https://github.com/lucagrulla/cw/issues/214

## Conclusion

The tool, `cw` is fast and useful. The error you saw made me bothered. I have a issue about that, but if someone has an idea, please submit a Pull Request.

## Original

https://zenn.dev/yuta28/articles/cloudwatch-fast-tail
