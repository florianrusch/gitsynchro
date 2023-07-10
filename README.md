# gitsynchro

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/florianrusch/gitsynchro) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/florianrusch/gitsynchro) [![Go Report Card](https://goreportcard.com/badge/github.com/florianrusch/gitsynchro)](https://goreportcard.com/report/github.com/florianrusch/gitsynchro)

1. _As a developer, I want to automatically push the default branch to somewhere else._
2. _As a developer, I want to automatically push all tags to somewhere else._

## Run it locally

```shell
$ gitsynchro is a tool to synchronize git repos.

Usage:
  gitsynchro [flags]

Flags:
      --config string   config file (default: $HOME/gitsynchro.yaml or ./gitsynchro.yaml)
  -h, --help            help for gitsynchro
```

## Config Locations

The application will look for the application in under the following paths:

- `$HOME/gitsynchro.yaml`
- `$HOME/gitsynchro/gitsynchro.yaml`
- `./gitsynchro.yaml`

In addition, it's possible to specify a custom path. You just need to specify the path as a command line flag like this: `--config my-path/whatever.yaml`

## Configuration

```yaml
repos:
  - name: very-good-repo
    path: /tmp/test
    defaultBranch: main
    destinations:
      - remoteName: test
```

## ToDos

- [ ] Push Tags: <https://github.com/go-git/go-git/blob/35f7e6770361a2c16c9b6c44acdc38ae04c75bd3/_examples/tag-create-push/main.go#L128C6-L132>
- [ ] Real synchronization: Fetch changes from "origin"
- [ ] Generate JSON Schema for configuration
- [ ] Double check disabled linters
- [ ] Implement further subcommands to show and edit the config
- [ ] Improve use-case description

## License

gitsynchro is free and unencumbered public domain software. For more information, see <https://unlicense.org/> or the accompanying [LICENSE](/LICENSE) file.
