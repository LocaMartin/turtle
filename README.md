<p align="center"><img src="https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExNGVmMXJmNWtzM3VyZ2draWd0NGtrenlhYndjdGVidGRudHF5N25kZiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/rlChzWTthWgyA/giphy.gif" width="200" height="200"/></p>

<h3 align="center">Turtle</h3>

Turtle is a commandline file sharing tool used to share files from terminal to telegram bot , Automates config management and simplifies file sharing with minimal flags.

![Go Version](https://img.shields.io/badge/Go-1.24.1-blue?logo=go&logoColor=white)

**Features:**

- **One-time setup** for bot token and chat ID
- **Simple file sharing** with single flag
- **Auto-saves config** to `~/.turtle_config.json`
- **Clean config removal**
- Supports any file type Telegram allows
- No external dependencies beyond Go

**Installation**

```bash
go install github.com/LocaMartin/turtle@latest
```
**Usage:**

- Run turtle setup command:

```bash
turtle -id "YOUR_BOT_TOKEN"
```
- send file

```bash
turtle -f file.txt
```
- Remove config
```bash
turtle -clean
```
e
First-time Setup

- Create a Telegram bot using @BotFather
- open telegram search **`BotFather`**
  - use these command/message in telegram's **`BotFather`** bot
    ```bash
    /start
    /newbot
    <your bot name> # bot name
    <yourbotnamebot> # your username must end with 'bot'
    ``` 
   - You will get message from @BotFather like this that will have link to your bot and bot id
```bash
Done! Congratulations on your new bot. 
You will find it at 't.me/<your bot name>bot'. # your bot link
You can now add a description, about section and profile picture for your bot, see /help for a list of commands.
By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. 
Just make sure the bot is fully operational before you do this.

Use this token to access the HTTP API:
'7204102675:ghdegbjbfgkjdfjwedwefweewefqwe_Y-YHA0' # your bot id
Keep your token secure and store it safely, it can be used by anyone to control your bot.

For a description of the Bot API, see this page: https://core.telegram.org/bots/api

````
- Open your bot using the link you got from `BotFather`
   - Start your bot with this command in your bot 
   ```bash
   /start
   ```
**Flags:**
| Flag | Description | Example |
|------|-------------|---------|
| `-f` | File path to share | `-f ~/documents/report.pdf` |
| `-id` | Set/update bot token (first-time setup) | `-id "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"` |
| `-clean` | Remove stored configuration | `-clean` |
| `-h` | Help message | `-h` |
| `-v` | Version | `-v` |

> Click [here](https://github.com/LocaMartin/turtle.wiki.git) to know how turtle tool works
## Contributing
1. Fork the repository
2. Create feature branch (`git checkout -b feature`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push to branch (`git push origin feature`)
5. Open Pull Request
