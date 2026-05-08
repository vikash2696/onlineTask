# GoLang2 Project

This project contains both a Go HTTP server and a Cloudflare Pages static site.

## Local Development (Go Server)

Run the Go server locally:

```bash
go run menu.go
```

Visit: http://localhost:8080

## Deploy to Cloudflare Pages

Deploy the static site to Cloudflare:

```bash
# Install Wrangler (if not already installed)
npm install -g wrangler

# Login to Cloudflare
wrangler login

# Deploy (use the correct Pages command)
wrangler pages deploy public --project-name=golang2-static-site

# Or using npx without global install:
npx wrangler pages deploy public --project-name=golang2-static-site
```

## Project Structure

```
GoLang2/
├── menu.go          # Go HTTP server (local development)
├── go.mod              # Go module file
├── wrangler.toml       # Cloudflare configuration
├── public/             # Static files for Cloudflare Pages
│   └── index.html      # Main HTML file
└── README.md           # This file
```
