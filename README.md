<p align="center">
  <img src="./assets/Logo_With_BG.png" width="280" alt="TgCloud Logo">
</p>

<h1 align="center">TgCloud</h1>
<p align="center">
  <b>Open-source image hosting powered by Telegram storage. Self-hosted, no database required.</b><br>
  Self-hosted • Zero Cloud Costs • Upload via File or URL • Powered by Go
</p>

<p align="center">
  <img src="https://img.shields.io/badge/go-1.22+-blue?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/github/v/release/AdrienMttn/TgCloud?style=flat-square&logo=github" alt="Release">
  <img src="https://img.shields.io/github/license/AdrienMttn/TgCloud?style=flat-square" alt="LICENSE">
  <img src="https://img.shields.io/github/downloads/AdrienMttn/TgCloud/total?style=flat-square" alt="downloads">
  <img src="https://img.shields.io/github/stars/AdrienMttn/TgCloud?style=flat-square" alt="stars">
</p>

<p align="center">
  <a href="https://adrienmttn.github.io/TgCloud">View Demo</a> •
  <a href="#features">Features</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#api-reference">API</a> •
</p>

---

## 🚀 Demo

**Live Instance:** `https://tgcloud.alwaysdata.net`

Upload from file:

```bash
curl -X POST -F "file=@photo.jpg" https://tgcloud.alwaysdata.net/img
```

Upload from URL:

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/image.png"}' \
  https://tgcloud.alwaysdata.net/img
```

**Result:**

```json
{
  "id": "AgACAgQAAxkDAAMraZyUGni2TrQXYkhK5ursyqon8wIAAsMOaxvNseFQ30BlvGjM3CIBAAMCAAN3AAM6BA"
}
```

## ✨ Features

- 📤 **Dual Upload**: Upload from local file (`multipart/form-data`) or remote URL (`application/json`)
- 🖼️ **Instant Hotlinks**: Direct image URLs compatible with Discord, Slack, Markdown, HTML `<img>`
- 💰 **$0 Forever**: Uses Telegram Bot API as free storage backend (2GB per channel)
- 🔒 **Privacy First**: Self-hosted, no third-party CDNs track your images
- 🚀 **One-Click Deploy**: Ready for Railway, Render, or any VPS
- 📋 **Format Validation**: Auto-rejects unsupported formats (JPEG, PNG, GIF, WebP only)

## 📋 Table of Contents

- [Quick Start](#quick-start)
- [API Reference](#api-reference)
- [Deployment](#deployment)
- [Environment Variables](#environment-variables)
- [Limitations](#limitations)

## ⚡ Quick Start

### Prerequisites

- Go 1.22+
- Telegram Bot Token ([@BotFather](https://t.me/BotFather))
- Channel/Group Chat ID (where bot is admin)

### Local Installation

```bash
git clone https://github.com/AdrienMttn/TgCloud.git
cd TgCloud

# Edit .env with your BOT_TOKEN and CHAT_ID and port

go mod tidy
go run main.go
```

Server starts on `http://localhost:8100` (or your `PORT` env var).

**Find your Chat ID:**

1. Message your bot or add it to a channel
2. Visit: `https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates`
3. Look for `"chat":{"id":-1001234567890}` (use the number, including the minus if present)

## 📚 API Reference

### `GET /format`

Returns accepted MIME types.

**Response:**

```json
{
  "supports": ["image/jpeg", "image/png", "image/gif", "image/webp"]
}
```

### `POST /img`

Upload image from local file.

- **Content-Type:** `multipart/form-data`
- **Field name:** `file`

**cURL Example:**

```bash
curl -X POST \
  -F "file=@/path/to/image.jpg" \
  https://tgcloud.alwaysdata.net/img
```

**Success (200):**

```json
{
  "id": "AgACAgQAAxkD..."
}
```

**Error (400):**

```json
{
  "error": "Error while uploading the image on Telegram"
}
```

### `POST /img`

Upload image from remote URL.

- **Content-Type:** `application/json`
- **Body:** `{"url": "https://example.com/image.jpg"}`

**cURL Example:**

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/photo.png"}' \
  https://tgcloud.alwaysdata.net/img
```

**Response:** Same as `/img`

### `GET /img/{id}`

Retrieve and display image by Telegram File ID.

**Example:**

```
https://tgcloud.alwaysdata.net/img/AgACAgQAAxkDAAMraZyUGni2TrQXYkhK5ursyqon8wIAAsMOaxvNseFQ30BlvGjM3CIBAAMCAAN3AAM6BA
```

<p align="center">
  <img src="https://tgcloud.alwaysdata.net/img/AgACAgQAAxkDAAMraZyUGni2TrQXYkhK5ursyqon8wIAAsMOaxvNseFQ30BlvGjM3CIBAAMCAAN3AAM6BA" width="300" alt="Example Image">
</p>

## 🔧 Environment Variables

Create `.env` file:

```env
BOT_TOKEN=your_telegram_bot_token_here
CHAT_ID=-1001234567890
PORT=8100
```

- `BOT_TOKEN`: Get from [@BotFather](https://t.me/BotFather)
- `CHAT_ID`: Channel or group ID (must be integer, channels usually start with `-100`)
- `PORT`: Server port

## 🚀 Deployment

TgCloud is distributed as a pre-built binary, so you don't need Go installed to run it.

### 1. Download the latest release

Open the [latest release](https://github.com/AdrienMttn/TgCloud/releases/latest) and download the archive that matches your operating system and CPU architecture:

| OS      | Architecture            | File                          |
| ------- | ----------------------- | ----------------------------- |
| Linux   | x86_64 (Intel/AMD)      | `tgcloud_Linux_x86_64.tar.gz` |
| Linux   | arm64 (e.g. Raspberry Pi 4/5) | `tgcloud_Linux_arm64.tar.gz`  |
| macOS   | Intel                   | `tgcloud_Darwin_x86_64.tar.gz` |
| macOS   | Apple Silicon (M1/M2/…) | `tgcloud_Darwin_arm64.tar.gz` |
| Windows | x86_64                  | `tgcloud_Windows_x86_64.zip`  |

> Not sure which one to pick? On Linux/macOS, run `uname -m`: `x86_64` means x86_64, `aarch64` or `arm64` means arm64.

On Linux or macOS you can also download and extract it in one command (example for Linux x86_64):

```bash
curl -L https://github.com/AdrienMttn/TgCloud/releases/latest/download/tgcloud_Linux_x86_64.tar.gz | tar -xz
```

On Windows, right-click the `.zip` archive and choose **Extract All...**

### 2. Configure the environment variables

TgCloud is configured through environment variables. Create a `.env` file next to the binary:

```env
BOT_TOKEN=your_telegram_bot_token_here
CHAT_ID=-1001234567890
PORT=8100
```

| Variable    | Description                                                                 |
| ----------- | --------------------------------------------------------------------------- |
| `BOT_TOKEN` | Your Telegram bot token, obtained from [@BotFather](https://t.me/BotFather) |
| `CHAT_ID`   | ID of the channel or group where images are stored (the bot must be an admin). Channel IDs usually start with `-100` |
| `PORT`      | Port the server listens on (e.g. `8100`)                                    |

On a VPS or hosting platform, you can set these variables in the platform's settings or in your service configuration instead of using a `.env` file.

### 3. Run it

```bash
# Linux / macOS
chmod +x ./tgcloud
./tgcloud

# Windows
.\tgcloud.exe
```

The server is now available at `http://localhost:8100` (or the port you configured).

## ⚠️ Limitations

- **Rate Limiting**: ~30 requests/minute to Telegram API (images are proxied, not direct links)
- **File Size**: Max 20MB per image (Telegram Bot API limit)
- **Permanence**: File IDs persist as long as the message exists in your Telegram channel
- **Bandwidth**: Your server proxies image data (uses your bandwidth, not Telegram's)

## 🤔 Why TgCloud vs Traditional Hosting?

| Feature         | TgCloud        | AWS S3 Free    | Imgur API     |
| --------------- | -------------- | -------------- | ------------- |
| **Cost**        | **$0 Forever** | $0 (12mo only) | Rate limited  |
| **Storage Cap** | Unlimited\*    | 5GB            | Unlimited     |
| **Hotlinking**  | ✅ Yes         | ✅ Yes         | ❌ Blocked    |
| **Direct URLs** | ✅ Yes         | ✅ Yes         | Redirect only |
| **Self-hosted** | ✅ Yes         | ❌ No          | ❌ No         |

\*2GB per channel, create unlimited channels

## 🛣️ Roadmap

- [X] Web UI for drag-and-drop upload
- [ ] Image resizing/thumbnails
- [ ] Multiple channel sharding for scale

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=AdrienMttn/TgCloud&type=date&legend=top-left)](https://www.star-history.com/#AdrienMttn/TgCloud&type=date&legend=top-left)

<p align="center">
  Made with ❤️ by <a href="https://adrienmttn.github.io">Adrien Mttn</a>
  <br>
  <a href="https://linkedin.com/in/adrien-metton">LinkedIn</a> •
  <a href="https://github.com/AdrienMttn">GitHub</a>
</p>
