# TinyShort

Simplest, resource-efficient self-hosted URL shortening app.

## Docker

```bash
docker login -u {{login}} -p {{token}} docker.pkg.github.com
docker pull docker.pkg.github.com/td0m/tinyshort/tinyshort:latest
docker run -v "$PWD/links.txt:/links.txt" -p 127.0.0.1:80:80/tcp docker.pkg.github.com/td0m/tinyshort/tinyshort:latest
```

Replace `{{login}}` with your GitHub login and `{{token}}` with a personal access token with permission to `read:packages`.