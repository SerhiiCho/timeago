![golang](.github/header.png)

[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)

Timeago is a fast and lightweight date time package that converts given date into human readable "n time ago" format in different languages, such as 🇬🇧 🇷🇺 🇺🇦 🇳🇱 🇩🇪 🇨🇳 🇧🇾 🇪🇸 🇯🇵 🇫🇷.

- [🗒 Documentation](https://timeago.serhiicho.com/)
- [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/CHANGELOG.md)

## Quick Start

```bash
go get github.com/SerhiiCho/timeago/v4
```

## Codebase naming

Here are some of the naming conventions used in the codebase and their meanings for better understanding:

- **time unit** - a single time unit like `second`, `seconds`, `minute`, `minutes`, etc.
- **time number** - a number of time units like `1`, `2`, `3`, etc. in a string like `1 minute ago`, `2 minutes ago`, `3 minutes ago`, etc.
- **suffix** - the suffix `ago` in the final output like `1 minute ago`, `2 minutes ago`, `3 minutes ago`, etc.
- **time since** - the final output result like `2 minutes`, `3 minutes ago`, `Just now`, `Online`, `3 years ago` etc.

### Available Commands

All commands run inside a container. The image is built automatically the first time you run any command:

```bash
make test
```

- `make build` - Build the checks container image
- `make test` - Run all tests
- `make fmt` - Format Go code
- `make lint` - Run linter
- `make check` - Combines commands like: test, fmt, lint
- `make` - Defaults to `make check`

Commands use `podman-compose` by default. To use Docker Compose instead, override `OCI`:

```bash
make check OCI='docker compose'
```

## License

This project is open-sourced software licensed under the [MIT license](https://github.com/SerhiiCho/timeago/blob/master/CHANGELOG.md).
