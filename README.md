![golang](https://serhii.io/storage/other/timeago.png)

[![Go](https://github.com/SerhiiCho/timeago/actions/workflows/go.yml/badge.svg)](https://github.com/SerhiiCho/timeago/actions/workflows/go.yml)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FSerhiiCho%2Ftimeago%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/SerhiiCho/timeago/goto?ref=master)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![GitHub](https://img.shields.io/github/license/SerhiiCho/timeago)](https://github.com/SerhiiCho/timeago/blob/master/LICENSE.md)

Timeago is a fast and lightweight date time package that converts given date into human readable "n time ago" format in different languages, such as 🇬🇧 🇷🇺 🇺🇦 🇳🇱 🇩🇪 🇨🇳 🇧🇾 🇪🇸 🇯🇵 🇫🇷. For more information you can read the [documentation](https://time-ago.github.io/).

### Follow the [Official documentation](https://time-ago.github.io/) for all the details
- [🗒 Documentation](https://time-ago.github.io/)
- [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/CHANGELOG.md)

## Quick Start
```bash
go get github.com/SerhiiCho/timeago/v3
```

## Codebase naming
Here are some of the naming conventions used in the codebase and their meanings for better understanding:

- **time unit** - a single time unit like `second`, `seconds`, `minute`, `minutes`, etc.
- **time number** - a number of time units like `1`, `2`, `3`, etc. in a string like `1 minute ago`, `2 minutes ago`, `3 minutes ago`, etc.
- **suffix** - the suffix `ago` in the final output like `1 minute ago`, `2 minutes ago`, `3 minutes ago`, etc.
- **time since** - the final output result like `2 minutes`, `3 minutes ago`, `Just now`, `Online`, `3 years ago` etc.

## Star History
[![Star History Chart](https://api.star-history.com/svg?repos=SerhiiCho/timeago&type=Date)](https://www.star-history.com/#SerhiiCho/timeago&Date)

## License
This project is open-sourced software licensed under the [MIT license](https://github.com/SerhiiCho/timeago/blob/master/LICENSE.md).

## Contribute
### With Container Engine
> [!NOTE]
> If you use [🐳 Docker](https://app.docker.com/) instead of [🦦 Podman](https://podman.io/), just replace `podman-compose` with `docker compose`, and `podman` with `docker` in code examples below.

#### Build an Image
To build an image, navigate to the root of the project and run this command:
```bash
podman-compose build
```

#### Run the Container
To run a container, navigate to the root of the project and run this command to enter the Linux container with Go installed:
```bash
podman-compose run --rm app
```

Inside of this container you'll be able to run `make` commands like `make test` to run tests.
