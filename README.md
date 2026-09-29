# tako-demo-hello 🐙

A minimal, high-performance Go web service designed for zero-downtime deployment demos on [TAKO](https://gettako.dev).

## Features

- **Dual Response Formats**:
  - `Accept: text/html` returns a dark-mode styled HTML status page.
  - `Accept: application/json` or `curl` returns structured JSON:
    ```json
    {
      "message": "Hello World from TAKO! 🐙",
      "status": "ok",
      "timestamp": "2026-09-29T05:12:00Z",
      "hostname": "864d4b1a43a1"
    }
    ```
- **Health Check Endpoint**: `/healthz` returning HTTP 200 `OK`.
- **Configurable Port**: Listens on `PORT` environment variable (defaults to `3000`).
- **Tiny Docker Footprint**: Multi-stage static Go binary on Alpine (~10 MB image size).

## Local Development

```bash
go run main.go
```

## Docker Build & Run

```bash
docker build -t tako-demo-hello .
docker run -p 3000:3000 tako-demo-hello
```

Test it:
```bash
curl http://localhost:3000/
curl http://localhost:3000/healthz
```

## Deploying on TAKO

1. Open your TAKO Control Plane dashboard.
2. Click **Create Project** -> **Create Service**.
3. Enter repository: `https://github.com/supianidz/tako-demo-hello.git` (or select from GitHub App integration).
4. Branch: `main`.
5. Internal Port: `3000`.
6. Click **Deploy**. TAKO will automatically clone, build the container, perform health checks, and configure Traefik routing with zero downtime!
