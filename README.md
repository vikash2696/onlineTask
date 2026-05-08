# GoLang2 Project

This project contains both a Go HTTP server and a Cloudflare Pages static site.

## Local Development (Go Server)

Run the Go server locally:

```bash
go run menu.go
```

Visit: http://localhost:8080

## Deploy to Cloudflare Pages

**Option 1: Via Cloudflare Dashboard (Recommended):**
1. Go to your Cloudflare Pages project settings
2. Change **Build command** to: (empty - delete `npx wrangler deploy`)
3. Set **Build output directory** to: `public`
4. Save and redeploy

**Option 2: Keep using wrangler.toml:**
The wrangler.toml is now configured correctly. In Cloudflare Pages settings:
- Set **Build command** to: (empty)
- It will auto-deploy the `public` folder

**Option 3: Manual CLI deployment:**
```bash
npx wrangler pages deploy public --project-name=golang2-static-site
```

## Project Structure

```
GoLang2/
├── menu.go          # Go HTTP server (local development)
├── go.mod              # Go module file
├── wrangler.toml       # Cloudflare configuration
├── package.json        # NPM config (optional)
├── public/             # Static files for Cloudflare Pages
│   └── index.html      # Main HTML file
└── README.md           # This file
```
