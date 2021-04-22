# labchat

[![Go Report Card](https://goreportcard.com/badge/github.com/yoonsue/labchat)](https://goreportcard.com/report/github.com/yoonsue/labchat)
[![Coverage](https://codecov.io/gh/yoonsue/labchat/branch/master/graph/badge.svg)](https://codecov.io/gh/yoonsue/labchat)
[![Build Status Travis](https://img.shields.io/travis/yoonsue/labchat.svg?style=flat-square&&branch=master)](https://travis-ci.org/yoonsue/labchat)
[![Releases](https://img.shields.io/github/release/yoonsue/labchat/all.svg?style=flat-square)](https://github.com/yoonsue/labchat/releases)
[![LICENSE](https://img.shields.io/github/license/yoonsue/labchat.svg?style=flat-square)](https://github.com/yoonsue/labchat/blob/master/LICENSE)

### AS Kakao changed the policy commercially, we stopped the service.

LABchat is a chatbot for convenient conversation via texual methods.
It's just for 1:1 chat. (If Kakao opens the chat with multiple users for free, it will be implemented.)
Kakao provides notify function which sends message immediately or on set time, but it's not for free. [Kakao Bizmessage: Notification API](https://bizmessage.kakao.com)

LABchat contains useful functions below :
* auto-reply at some text
* get server status
* get today's school cafeteria menu

labchat is written in Go and uses the [Kakao plusfriend API](https://github.com/plusfriend/auto_reply).

## Build

```sh
go build labchat.go
```

## Run

```sh
./labchat
```
