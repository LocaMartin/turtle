## How It Works
### Configuration Management
On first run with `-id`, tool:
1. Stores bot token in `~/.turtle_config.json`
2. Fetches your Chat ID via Telegram API
3. Saves both credentials for future sessions

### File Sharing
When using `-f` flag:
1. Reads config file for credentials
2. Sends file via Telegram Bot API
3. Returns success/failure status

> Before this tool i had to mannually enter bot id and chat id every time inorder to send a file like this:
```bash
curl -F document=@"file.txt" "https://api.telegram.org/<bot id>/sendDocument?chat_id=534950675"
```
