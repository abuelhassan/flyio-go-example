# Fly.io Go Example

## Deploy to Fly.io
1. Install flyctl and login:
   ```bash
   fly auth login
   ```
2. Bootstrap the app (first time deploys and provisions resources):
   ```bash
   fly launch
   ```
3. For subsequent updates, deploy using the existing `fly.toml`:
   ```bash
   fly deploy
   ```
4. Visit the allocated https://<app>.fly.dev URL.
5. When done, remove the app:
   ```bash
   fly apps destroy <app>
   ```
