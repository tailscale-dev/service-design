# what

what is a service for recording and sharing what you've been working on. It
implements the flows of
[Google's Snippets](http://blog.idonethis.com/google-snippets-internal-tool/) so
that you can record what you worked on after the fact as opposed to recording
what you'd like to work on before the fact.

## Usage

1. Open http://what
2. Type words in the box
3. Hit post

## Configuration

Configuration is done primarily with enviroment variables. Here are the relevant
variables for what:

| Variable         | Description                                                                                                            | Default                                   |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------- | ----------------------------------------- |
| `TS_AUTHKEY`     | The authkey to connect your node to Tailscale. This should only be needed once when setting up the service.            |                                           |
| `TSNET_HOSTNAME` | The hostname to use on your tailnet.                                                                                   | `what`                                    |
| `TSNET_VERBOSE`  | Whether to print verbose logs from tsnet. This can be useful to help diagnose connection issues with your what server. | `false`                                   |
| `DATA_DIR`       | The directory to store data in.                                                                                        | `$HOME/.config/tailscale/$TSNET_HOSTNAME` |
| `SLOG_LEVEL`     | The level of logs to print.                                                                                            | `INFO`                                    |

## Development

Install the following tools:

- Go 1.21 or later
- Node.js 18 or later
- Yarn 1.22 or later

Run this command to set up your node_modules for tailwind:

```
yarn
```

Then run this command to start the server:

```
go run . --tsnet-verbose --tsnet-hostname what-dev
```

Then authenticate with your Tailscale account and visit your service.

## Making production builds

Build it with Docker or Podman:

```
docker build .
```

or

```
podman build .
```

## Deployment

You can use the following Docker compose file to deploy this to your
infrastructure:

```yaml
services:
  what:
    image: ghcr.io/tailscale-dev/what:latest
    restart: always
    volumes:
      - what-data:/data
    environment:
        - TS_AUTHKEY=tskey-auth-hunter2-hunter2hunter2hunter2
        - TSNET_HOSTNAME=what
        - TSNET_VERBOSE=true
        - DATA_DIR=/data
        - SLOG_LEVEL=INFO

volumes:
    what-data:
```
