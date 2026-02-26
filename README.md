<p align="center">
  <img src="./assets/Logo_With_BG.png" width="280" alt="TgCloud Logo">
</p>

<h1 align="center">TgCloud</h1>
<p align="center">
  <b>Free Image Hosting CDN using Telegram as Unlimited Storage</b><br>
  Self-hosted • Zero Cloud Costs • Upload via File or URL • Powered by Go
</p>

<p align="center">
  <img src="https://img.shields.io/badge/go-1.22+-blue?style=flat-square&logo=go" alt="Go">
  <a href="https://railway.app"><img src="https://img.shields.io/badge/Deploy-Railway-9cf?style=flat-square&logo=railway" alt="Railway"></a>
  <a href="https://render.com"><img src="https://img.shields.io/badge/Deploy-Render-46e3b7?style=flat-square&logo=render" alt="Render"></a>
</p>

<p align="center">
  <a href="#demo">View Demo</a> •
  <a href="#features">Features</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#api-reference">API</a> •
  <a href="#deployment">Deploy</a>
</p>

---

## 🚀 Demo

**Live Instance:** `https://tgcloud.alwaysdata.net`

Upload from file:

```bash
curl -X POST -F "file=@photo.jpg" https://tgcloud.alwaysdata.net/post-from-file
```

Upload from URL:

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/image.png"}' \
  https://tgcloud.alwaysdata.net/post-from-url
```

**Result:**

```json
{
  "status": 200,
  "message": "Image postée avec succès",
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

### `GET /supports-formats`

Returns accepted MIME types.

**Response:**

```json
{
  "formats": ["image/jpeg", "image/png", "image/gif", "image/webp"]
}
```

### `POST /post-from-file`

Upload image from local file.

- **Content-Type:** `multipart/form-data`
- **Field name:** `file`

**cURL Example:**

```bash
curl -X POST \
  -F "file=@/path/to/image.jpg" \
  https://tgcloud.alwaysdata.net/post-from-file
```

**Success (200):**

```json
{
  "status": 200,
  "message": "Image postée avec succès",
  "id": "AgACAgQAAxkD..."
}
```

**Error (400):**

```json
{
  "status": 400,
  "message": "Erreur lors de l'upload."
}
```

### `POST /post-from-url`

Upload image from remote URL.

- **Content-Type:** `application/json`
- **Body:** `{"url": "https://example.com/image.jpg"}`

**cURL Example:**

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/photo.png"}' \
  https://tgcloud.alwaysdata.net/post-from-url
```

**Response:** Same as `/post-from-file`

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

### Railway (Recommended)

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/button)

### Render

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

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

- [ ] Local file caching (Redis/filesystem)
- [ ] Web UI for drag-and-drop upload
- [ ] Image resizing/thumbnails
- [ ] API authentication keys
- [ ] Multiple channel sharding for scale

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=AdrienMttn/TgCloud&type=date&legend=top-left)](https://www.star-history.com/#AdrienMttn/TgCloud&type=date&legend=top-left)

<p align="center">
  Made with ❤️ by <a href="https://github.com/AdrienMttn">Adrien Mttn</a>
  <br>
  <a href="https://linkedin.com/in/adrien-metton">LinkedIn</a> •
  <a href="https://github.com/AdrienMttn">GitHub</a>
</p>
