# DiscordGo Starter Files

[![Go Version](https://img.shields.io/badge/Go-1.16%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Discord](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white)](https://discord.gg/your-invite-link)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

A simple and efficient Discord bot starter template built with Go using the [discordgo](https://github.com/bwmarrin/discordgo) library. This project provides a solid foundation for building feature-rich Discord bots.

## ✨ Features

- 🏓 Simple ping command to check bot responsiveness
- 🛠️ Easy-to-extend command structure
- ⚡ Built with Go for high performance
- 🔄 Slash commands integration

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- A Discord Bot Token from the [Discord Developer Portal](https://discord.com/developers/applications)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/1AhmedS/discordgo-starter-files.git
   cd discordgo-starter-files
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build and run the bot:
   ```bash
   go run main.go
   ```

## 🔧 Configuration

1. Create a new application in the [Discord Developer Portal](https://discord.com/developers/applications)
2. Go to the "Bot" tab and create a new bot
3. Copy the bot token
4. Set the token in the `main.go` file:
   ```go
   var TOKEN = "your-bot-token-here"
   ```
5. Invite the bot to your server with the following permissions:
   - `applications.commands`
   - `bot` (with necessary permissions)

## 🛠️ Available Commands

- `/ping` - Check if the bot is alive

## 📝 Adding New Commands

1. Create a new file in the `commands` directory (e.g., `commands/yourcommand.go`)
2. Implement your command following the same pattern as `ping.go`
3. Register your command in `main.go`

## 💖 Support

If you find this project useful, please consider supporting my work:

[![PayPal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/1AhmedS)

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📧 Contact

For any questions or suggestions, please open an issue on GitHub.

---

<p align="center">
  <a href="https://github.com/1AhmedS">
    <img src="https://img.shields.io/badge/Follow-@1AhmedS-1DA1F2?style=for-the-badge&logo=github&logoColor=white" alt="GitHub">
  </a>
</p>
