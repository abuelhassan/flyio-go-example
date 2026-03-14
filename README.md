# Fly.io Go Example

## First-Time Launch
1. Install flyctl and login:
   ```bash
   fly auth login
   ```
2. Bootstrap the app without deploying:
   ```bash
   fly launch --no-deploy
   ```
3. Create the required S3 bucket. This also adds the required secrets to the app:
   ```bash
   fly storage create --name flyio-go-example --public
   ```
4. Upload a file named `file.txt` to the bucket using the Tigris UI.
5. Deploy using the generated `fly.toml`:
   ```bash
   fly deploy
   ```
6. Open the deployed app:
   ```bash
   fly apps open
   ```

## Relaunch
1. Bootstrap the app without deploying:
   ```bash
   fly launch --no-deploy
   ```
2. Load the required app secrets from `.env`:
   ```bash
   grep -v '^#' .env | fly secrets import --stage
   ```
3. Deploy using the generated `fly.toml`:
   ```bash
   fly deploy
   ```
4. Open the deployed app:
   ```bash
   fly apps open
   ```

## Cleanup
1. When done, remove the app:
   ```bash
   fly apps destroy <app>
   ```

Note: The Docker image must install `ca-certificates` so the app can make HTTPS requests to Tigris.

## Run Locally
```bash
env $(grep -v '^#' .env | xargs) go run .
```
