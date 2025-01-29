---
description: How to get started with gocord
---

# Getting Started

Getting started with gocord is easy! Thi guide will walk you through the process, starting with creating a bot in Discord's dev portal, and ending with a running bot.

## Creating a bot

To create a bot, you will need to create an application in the [Discord Developer Portal](https://discord.com/developers/applications). Once signed in, click the "New Application" button. Give your application a name, and click "Create". Next, click the "Bot" tab on the left. On this page, if there is no content, click "Add a bot". Here there are two important things you will need for later:

1. The bot's token. This can be obtained by clicking "Reset Token" and copying the value.
2. The application's ID. This is found on the "General Information" page.

## Installing gocord

To install gocord, first ensure you have [Go](https://go.dev/doc/install) installed. Then, run the following command:

```bash
go install github.com/dickeyy/gocord@latest
```

This will download the gocord binary to your `$GOPATH/bin` directory.

## Building a bot

Now that gocord is installed, you are ready to build a bot. In your terminal, run the following command:

```bash
gocord -name "Some name" -module "{your username}/some-name"
```

This will generate all the required files for your bot and setup your Go module.

Next, open the `.env` file in your project directory. This file contains all the environment variables your bot will use. You will need to add the following variables:

```bash
ENV={dev or prod}
REGISTER_CMDS=true

TOKEN=Bot {token from step 1}
APP_ID={application ID from step 2}
DEV_TOKEN=Bot {token from step 1}
DEV_APP_ID={application ID from step 2}
```

More documentation on the `.env` file and configuration options can be found [here](../configuration.md).

## Running your bot

Now that your bot is set up, you can run it,

```bash
go run main.go
```

This will start your bot, and register all the commands if you have `REGISTER_CMDS` set to `true` in your `.env` file. _Note: this needs to be set to `true` on the first run._

Now you can add your bot to your server via the [Discord Developer Portal](https://discord.com/developers/applications) and start using it.
