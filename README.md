# action

an example/experiment with a GitHub:Action written in [go].

[![Build](https://github.com/polis-dev/example-go-action/actions/workflows/build.yml/badge.svg)](https://github.com/polis-dev/example-go-action/actions/workflows/build.yml)

## features

- **minimalist**  
  - fewer "_moving parts_" to fail. 
  - easier to use as a template to start from. 

- **cross-platform**
  - written to be `*NIX` friendly, macOS/Linux hosts should work fine!
  - written to be CPU agnostic (`aarch64` vs `x86_64`) _[on my Apple Silicon MacBook Pro]_.
  
- **minimal dependencies**
  - [`lo`] adds a _many_ useful generic utilities (similar to [lodash]).
  - [`go-githubactions`] is a pleasant interface for writing GitHub:Actions.
  
- **fast startup time**
  - uses a pre-built image to avoid compiling the image at execution.
  - minimal dependencies mean the image stays _smaller_ so its faster to download at runtime.
  
- **very low maintenance**
  - using a `FROM scratch` container image means no packages to update; _perhaps aside from updating dependencies and/or [go] version(s)._
  - dependabot should open pull-requests automatically as necessary.

- **easy to write**
  - concrete types for discovering context provided by github.
  - easy to experiment locally via a few environment variables.
  - supports devcontainers/codespaces for easy remote development.

[`go-githubactions`]: https://github.com/sethvargo/go-githubactions#readme
[`lo`]: https://github.com/samber/lo#readme
[lodash]: https://lodash.com/
[go]: https://golang.org
