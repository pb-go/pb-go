# pb-cli specification

## Overview

`pb-cli` is a command line tool which communicate with `pb-go` backend. Although user can just upload data to pastebin directly with `curl`, we still provide a tools to help user easily use our advanced feature such as burn-after-reading and private share.

We guarantee everything that worked with `pb-cli`, also worked with `curl`.

Every command which communicates to server will print log(include http status code) to `stderr`, and print response body to `stdout`.

## Configuration

You can use `-c file` to provide a config file. If you didn't provide one, it will find config file called `.pbcli.yaml` from current directory then home directory.

Config file should be written in `yaml` like:

```yaml
host: 'http://your.pastebin.bakcend:port'
expire: 12
masterKey: 'Same as your config on server'
```

About `host`:  
You can use `-h <url>` with any command in `pb-cli`, it is a global flag.  
You should specify schema(http/https) in url, and no need to add port when using default port(80 for http, 443 for https).  
Notice we do NOT encrypt data while transporting, we have NO guarantee about transport layer safety, so just use HTTPS.

About `masterKey`:  
You can use `-k <masterKey>` with `delete` command.  
Only admin has masterKey. It is only for admin usage like delete some pastes.  
It you are a user, just ignore this.

About `expire`:  
You can use `-e <expire>` with `upload` command.  
Default pastes expire time in hours. If you do not set this value or `-e` flag, it will use 24 as expire time.

**Using flag in commands will override the value from config file.**

## Upload

Upload data to pastebin.

Usage:

```text
pb-cli upload [options] <file>
-p                Optional. Private share. Will using a random password for private share.
-P <PASSWORD>     Optional. Private share with specificated password.
-e <EXPIRE>       Optional. Set to 0 means burn-after-read. Default 24. (unit: hrs)
```

When file is not given, `pb-cli` will read data from `stdin`, like as a pipe.

## Get

Fetching data from patesbin with id.

Usage:

```text
pb-cli get [options] <id>
-p <PASSWORD>     Optional. Provide password for private share.
```

## Delete

Delete a paste from pastebin with id.(Need admin permission)
Usage:

```text
pb-cli delete [options] <id>
-k <MASTER_KEY>   Required. Master key in `pb-go` server's config.
```
