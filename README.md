# service-design

A template repository for creating tsnet services that "feel like Tailscale".
This will be for an imaginary service named "What", for recording and sharing
what you've been working on.

## Essential parts

- [ ] A `README.md` file that describes the service and how to use it, including
      installation instructions.
- [ ] A Dockerfile that builds the service.
- [ ] Tailwind configuration in `tailwind.config.js`.
- [ ] The base HTML skeleton in `tmp/base.html`.
- [ ] A static asset folder `static/`.
- [ ] A `main.go` file that serves the static assets and the HTML skeleton over
      [tsnet](https://tailscale.com/kb/1244/tsnet/).
- [ ] The font [Inter](https://rsms.me/inter/).

Some other things to consider:

- Users will run this in potentially many environments, so it should be as
  self-contained as possible given the constraints of your service.
- Any external dependencies (aside from Tailscale) MUST be documented clearly in
  plain English in the README.
- The service should be as easy to run as possible. If it requires a database,
  it should be created automatically.
- All configuration should be done in environment variables with flag overrides,
  and documented in the README.
- The main page of the service should contain everything users need to know in
  order to interact with it.
- Tailscale is your method of authentication. You should not need to implement
  any authentication logic.
- Commit the generated CSS to the repository. This makes it easier to run the
  service without having to install Node.js and Tailwind.
- Make your service one word. This makes it easier to use in URLs and
  documentation, not to mention pronouncing it.

## Optional parts

- [ ] A [Nix flake](https://nixos.wiki/wiki/Flakes) for users on NixOS. This
      should export a NixOS module that can be imported into a users' deployment
      flake.
- [ ] Deployment instructions for [fly.io](https://fly.io).

## Development

Install the following tools:

- [ ] [Go](https://go.dev/dl)
- [ ] [Node.js](https://nodejs.org/en/download/)
- [ ] [Yarn](https://classic.yarnpkg.com/en/docs/install)

Clone the repo and run:

```sh
yarn
go run . --tsnet-verbose
```

Then authenticate with your Tailscale account and visit your service.

## On designing tools
