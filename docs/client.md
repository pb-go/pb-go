# pb-cli specification

## Overview

`pb-cli` is a command line tool which communicate with `pb-go` backend. Although user can just upload data to pastebin directly with `curl`, we still provide a tools to help user easily use our advanced feature such as burn-after-reading and private share.

We guarantee everything that worked with `pb-cli`, also worked with `curl`.

## Configuration

You can use `-c file` to provide a config file. If you didn't provide one, it will find config file called `.pb-cli.toml` from current directory then home directory.

Config file should be written in `toml` like:

```toml
host = "http://your.pastebin.bakcend:port"
master_key = "Same as your config on server"
```

You can use `-h <url>` with any command in `pb-cli`, it is a global flag.
You can use `-k <master_key>` with `delete` command.

Using flag in commands will override the value from config file.

## Upload

Usage:

```text
pb-cli upload [options] <file>
-p                Optional. Private share. Will using a random password for private share.
-P <PASSWORD>     Optional. Private share with specificated password.
-e <EXPIRE>       Optional. Set to 0 means burn-after-read. Default 24. (unit: hrs)
```

When file is not given, `pb-cli` will read data from `stdin`, like as a pipe.

## Get

Usage:

```text
pb-cli get <id>
-p <PASSWORD> Optional. Provide password for private share.
```

## Delete

Usage:

```text
pb-cli delete <id>
```
