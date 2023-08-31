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
- Try [Alpine.js](http://alpinejs.dev/) or [HTMX](https://htmx.org/) for
  client-side interactivity.

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

When you are designing tools, you are designing things that will be used
productively. These things will be copied from, studied from, and transformed as
people learn how to wield them. It is a lot better to start out with something
very simple (and even a little bit ugly) that is easy to understand and modify,
than it is to start out with something complex and beautiful that is hard to
understand and modify. Twitter succeeded because it was as easy to use as
sending an SMS to 40404. It was not because it was a beautiful website. The main
difference that makes something a tool is it being adaptable to a given usecase.
[tclip](https://tailscale.dev/blog/tclip) is a tool because it's a place to put
data to share with your coworkers. It can be part of a greater flow or just a
one-off thing to share code snippets with friends. The point is that it's
adaptable enough to be used in usecases that you may have never dreamed
possible.

When you are designing such a tool, start with the problem you want to solve. In
the case of "what", here is such a problem statement:

> The standup flow is forward-looking and reality can and will diverge from what
> we hope we will get done. What if there was a place to explicitly share what
> we've already accomplished?

Everything else should flow out from there. The goal is to share with your
coworkers what you've done, what if there was a text box that you could only
update at the end of your workday with a summary of what you accomplished? How
would it look like to see what other people are doing? What are the most
important parts of this? How would you hook things in to drive people towards
this tool? What is the minimal expression that can help you test your
hypothesis?

The rest will come from trial, error, experimentation, failure, and eventual
success.

Start from the problem, and the solution will follow.
