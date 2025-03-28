package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	version      = "1.0.0"
	configFile   = ".turtle_config.json"
	apiURL       = "https://api.telegram.org/bot%s/%s"
)

type Config struct {
	BotToken string `json:"bot_token"`
	ChatID   int64  `json:"chat_id"`
}

func main() {
	filePath := flag.String("f", "", "File path to share")
	botToken := flag.String("id", "", "Set Telegram Bot Token")
	clean := flag.Bool("clean", false, "Remove configuration")
	showVersion := flag.Bool("v", false, "Show version")
	showHelp := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	if *showVersion {
		fmt.Printf("🐢 turtle v%s\n", version)
		return
	}

	if *clean {
		removeConfig()
		return
	}

	if *botToken != "" {
		handleBotTokenSetup(*botToken)
		return
	}

	if *filePath != "" {
		shareFile(*filePath)
		return
	}

	printHelp()
}

func printHelp() {
	fmt.Println(`🐢 turtle - Telegram File Sharing Tool

Usage:
  turtle [command]

Commands:
  -f string    File path to share
  -id string   Set Telegram Bot Token (first-time setup)
  -clean       Remove configuration
  -v           Show version
  -h           Show help

Examples:
  turtle -id "123456:ABC-DEF1234ghIkl"
  turtle -f document.pdf
  turtle -clean`)
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, configFile)
}

func loadConfig() (*Config, error) {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("❌ config not found. Run 'turtle -id YOUR_BOT_TOKEN' first")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	return &config, err
}

func saveConfig(config *Config) error {
	file, err := os.Create(getConfigPath())
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(config)
}

func removeConfig() {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("ℹ️ No configuration found")
		return
	}

	os.Remove(configPath)
	fmt.Println("✅ Configuration removed")
}

func handleBotTokenSetup(token string) {
	fmt.Println("🔄 Please send a message to your bot in Telegram and press Enter...")
	fmt.Scanln()

	chatID, err := getChatID(token)
	if err != nil {
		fmt.Printf("❌ Error getting chat ID: %v\n", err)
		return
	}

	config := &Config{
		BotToken: token,
		ChatID:   chatID,
	}

	if err := saveConfig(config); err != nil {
		fmt.Printf("❌ Error saving config: %v\n", err)
		return
	}

	fmt.Printf("✅ Configuration saved!\nChat ID: %d\n", chatID)
}

func getChatID(token string) (int64, error) {
	url := fmt.Sprintf(apiURL, token, "getUpdates")
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		OK     bool `json:"ok"`
		Result []struct {
			Message struct {
				Chat struct {
					ID int64 `json:"id"`
				} `json:"chat"`
			} `json:"message"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}

	if !result.OK {
		return 0, fmt.Errorf("telegram API error")
	}

	if len(result.Result) == 0 {
		return 0, fmt.Errorf("no messages found. Send a message to your bot first")
	}

	return result.Result[0].Message.Chat.ID, nil
}

func shareFile(filePath string) {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("❌ Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	url := fmt.Sprintf(apiURL, config.BotToken, "sendDocument")
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Printf("❌ Error creating request: %v\n", err)
		return
	}

	query := req.URL.Query()
	query.Add("chat_id", fmt.Sprintf("%d", config.ChatID))
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Error sending file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("❌ Telegram API error (%d): %s\n", resp.StatusCode, body)
		return
	}

	fmt.Println("✅ File sent successfully!")
}