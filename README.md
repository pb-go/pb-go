# pb-go

Yet Another PasteBin implemented in Golang.

## Prerequisites

- MongoDB
- Reverse Proxy with HTTPS and Rate-Limit Support (Recommend: Traefik, Caddy)
- A Linux Server (If you need Windows version, compile by yourself.)

Note: Since we are offering public services, we don't want to implement any rate-limit
on application side and we will ONLY allow you to listen `127.0.0.1` instead of 

## To-Do list (features)

- [ ] | Content detection, only allow pure texts.
- [ ] | Expiring feature done in MongoDB.
- [ ] | Encrypting optionally.
- [X] | <del> Rate-limit to avoid abusing. (SHOULD BE DONE IN REVERSE PROXY SIDE) </del>
- [ ] | Code Syntax Highlighting.
- [ ] | Shortlink using hashids.
- [ ] | Pure CLI.
- [ ] | Web page upload.

## Usage

TODO

## Compile

TODO

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

