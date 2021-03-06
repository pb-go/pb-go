# pb-go

![Logo](./readme-logo.png)

Yet Another PasteBin implemented in Golang. The `pb-go` team have no affiliate with [ptpb/pb](https://github.com/ptpb/pb).

![GitHub stars](https://img.shields.io/github/stars/pb-go/pb-go?style=social)
[![Go Report Card](https://goreportcard.com/badge/github.com/pb-go/pb-go)](https://goreportcard.com/report/github.com/pb-go/pb-go)
![Go CI Build Status](https://github.com/pb-go/pb-go/workflows/Go/badge.svg)
![GitHub](https://img.shields.io/github/license/pb-go/pb-go)
![GitHub last commit](https://img.shields.io/github/last-commit/pb-go/pb-go)
![GitHub All Releases](https://img.shields.io/github/downloads/pb-go/pb-go/total)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/269b77a2b64c41bbaa4aa109ecf4d55a)](https://www.codacy.com/manual/pb-go/pb-go)
![Codacy coverage](https://img.shields.io/codacy/coverage/269b77a2b64c41bbaa4aa109ecf4d55a?logo=codacy)

We use [Sentry.io](https://sentry.io) for bug tracking and log collection which was GDPR-complaint, 
their privacy policy can be found at: [here](https://sentry.io/legal/privacy/2.1.0/)

Our Documentation is [here](/docs).

## Discussion

We need developer and help, for feature request and discussion, 
please go to our [Telegram Group](https://t.me/pb_go_discuss).

Bug Report please attach log and finish the whole GitHub Issue template. Thanks.

## Prerequisites

-   MongoDB
-   Reverse Proxy with HTTPS and Rate-Limit Support (Recommend: Traefik, Caddy)
-   A Linux Server (If you need Windows version, compile by yourself.)

Note: Since we are offering public services, we don't want to implement any rate-limit
on application side. You must apply a reverse proxy or something else do that.
You must ensure your proxy is properly configured to send `X-Real-IP` header.
Your data is encrypted and finally stored on our server using Chacha20 algorithm.

## To-Do list (features)

-   [X] | Content detection, only allow pure texts.
-   [X] | Expiring feature done in MongoDB. Support Read-After-Burn.
-   [X] | Private Share optionally, Sharing password using BLAKE2b stored. 
-   [X] | <del> Rate-limit to avoid abusing. (SHOULD BE DONE IN REVERSE PROXY SIDE) </del>
-   [X] | ReCaptcha v2 support to prevent from a large scale abusing.
-   [X] | Code Syntax Highlighting.
-   [X] | Shortlink using nano-id.
-   [X] | Pure CLI. (You could also use `curl` instead)
-   [X] | Web page upload.

## Usage

Current version: v1.0.4

### Client

Download `pb-cli` From release assets, use `--help` to get usage.

We have an example config named `pbcli.example.yaml` in our root folder to simplify your typing, 
Copy it to `~/.pbcli.yaml`, and change it according to [the document](/docs/client.md). **This configuration file is OPTIONAL, not REQUIRED.**

### Server

Download `pb-go` From release assets, use `--help` to get usage.

We have an example config named `config.example.yaml` in our root folder, rename to `config.yaml`, 
and change it according to [the document](/docs/config.md), finally put it in the same folder of the binary.

## Compile

`make build`

## License

 pb-go
 Copyright (C) 2020  kmahyyg

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
