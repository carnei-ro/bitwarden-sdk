# Fetch Secret Value

This example shows how to fetch a secret value from the Bitwarden SECRETS MANAGER (not BitWarden Password Manager).

## Prerequisites

Go to repo root and run `cargo build`, then `npm install` and `npm run schemas`.

## Build

Simply run `go build`.

If it fails about ld (libs) copy the target/debug folder by doing: `cp -vpr ../../../target/debug ../internal/cinterface/lib` then build again.

## Run

This example works with SaaS version, simply export:

- ACCESS_TOKEN: the service account access token
- ORGANIZATION_ID: the organization id (that's the UUID in the URL when you are in the Secrets Manager, such as `https://vault.bitwarden.com/#/sm/d2f88df6-b07b-4acb-9a0f-b09600cf89e5` the value would be `d2f88df6-b07b-4acb-9a0f-b09600cf89e5`)
- PROJECT_NAME: the project name the secret is in
- SECRET_KEY: the secret key name

Then run `./example` and you should see the secret value printed.

## Run as Docker

Build the image with `docker build -t example .` then run it with:

```bash
docker run -it --rm \
  -e ACCESS_TOKEN=... \
  -e ORGANIZATION_ID=... \
  -e PROJECT_NAME=... \
  -e SECRET_KEY=... \
  example
```
