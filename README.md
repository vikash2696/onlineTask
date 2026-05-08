# GoLang2 Project

This project contains both a Go HTTP server and a Cloudflare Pages static site.

## Local Development (Go Server)

Run the Go server locally:

```bash
go run menu.go
```

Visit: http://localhost:8080

## Deploy to Cloudflare Pages

**Via Cloudflare Dashboard (Automatic):**
1. Connect your Git repository to Cloudflare Pages
2. Configure build settings:
   - **Build command**: (leave empty)
   - **Build output directory**: `public`
3. Deploy automatically on git push

**Via Wrangler CLI (Manual):**
```bash
npx wrangler pages deploy public --project-name=golang2-static-site
```

## Project Structure

```
GoLang2/
├── menu.go          # Go HTTP server (local development)
├── go.mod              # Go module file
├── package.json        # NPM config (optional)
├── public/             # Static files for Cloudflare Pages
│   └── index.html      # Main HTML file
└── README.md           # This file
```
