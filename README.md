# LIFF command-line utility

A command-line utility for Line Frontend Framework

## Prerequisite

You needs to issue channel access token from LINE api (see https://developers.line.me/en/reference/messaging-api/#issue-channel-access-token) first. After
issued, you can set access token via shell environment `LIFF_ACCESS_TOKEN` or
use `--access-token`.

## List apps

```bash
$ liff list
```

## Add a LIFF app

```bash
$ liff add --view-type <compact,tall,full> --view-url https://<your-url>
```

## Delete a LIFF app

```bash
$ liff delete --liff-id <id>
```
