# Go Discord Bot

This is a Discord bot written in Go that sends notifications to a channel.

## Project Structure

```
go-discord-bot
├── .env
├── .vscode
│   └── launch.json
├── cmd
│   └── command-bot
│       └── main.go
│   └── message-bot
│       └── main.go
├── command-bot.dockerfile
├── go.mod
├── go.sum
├── internal
│   └── api
│       └── hub
│           └── hub.go
│       └── market
│           └── market.go
│   └── config
│       └── config.go
│   └── embed
│       └── embed.go
│   └── model
│       └── model.go
├── message-bot.dockerfile
└── README.md
```

## Usage

For the configuration of the bot, you need to create a `.env` file in the root directory of the project and use it in your current terminal session.
The `.env` file should contain the following environment variables:

```env
BOT_TOKEN=<your token>
CHANNEL_ID=<your channel id>
APPLICATION_ID=<your application id>
GUILD_ID=<your guild id>
```

You can also export the environment variables in your terminal session using the following command:

```bash
export BOT_TOKEN=<your token>
export CHANNEL_ID=<your channel id>
export APPLICATION_ID=<your application id>
export GUILD_ID=<your guild id>
```

Then, you can run the bot using the following command:
    
```bash
go run cmd/ordis/main.go
```

Feel free to modify and extend the functionality of this bot according to your needs.

```
go run cmd/<bot you want to use>/main.go
```

## Docker

To run a bot in a Docker container, you need to build the Docker image first. You can do this by running the following command:

```bash
docker build -t command-bot . -f command-bot.Dockerfile
```

Then, run Ordis the Docker container with the following command:

```bash
docker run --rm --env-file .env command-bot
```

## License

This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). See the [LICENSE](LICENSE) file for more details.
