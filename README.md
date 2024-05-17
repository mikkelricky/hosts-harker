# Hosts harker

## Quick start

Assuming [Go is installed](https://go.dev/doc/install), you can run a quick test with

``` shell
go run github.com/mikkelricky/hosts-harker@latest --help
```

to see what Hosts harker cat do for you.

## Installation

[Install Go](https://go.dev/doc/install) and install `hosts-harker` with

``` shell
go install github.com/mikkelricky/hosts-harker@latest
```

See [Compile and install packages and
dependencies](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) for details on where
`hosts-harker` is actually installed.

To set things straight and clean up, it may be worth running these commands:

``` shell
# Create the default installation location
mkdir -p ~/go/bin
# Clear GOBIN to use the default installation location
go env -w GOBIN=''
go install github.com/mikkelricky/hosts-harker@latest
```

Add `~/go/bin` to your `PATH`, e.g.

``` zsh
# ~/.zshrc
export PATH=$PATH:$HOME/go/bin
```

See [Completions](#completions) for details in how to set up completions for your terminal.

## Usage

``` shell
hosts-harker --help
```

### Completions

`hosts-harker` can automatically generate completions for your shell:

``` shell name=completion-help
hosts-harker help completion
```

#### Zsh

Load completions in [Zsh](https://en.wikipedia.org/wiki/Z_shell) by adding

``` zsh
# ~/.zshrc
eval "$(hosts-harker completion zsh)"; compdef _hosts-harker hosts-harker
```

to your `~/.zshrc`. If you're cool, you do it all from the command line:

``` shell name=zshrc-install-completion
cat >> ~/.zshrc <<'EOF'
eval "$(hosts-harker completion zsh)"; compdef _hosts-harker hosts-harker
EOF
```
