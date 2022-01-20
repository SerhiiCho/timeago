[![Code Coverage](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FSerhiiCho%2Ftimeago%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/SerhiiCho/timeago/goto?ref=master)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/SerhiiCho/timeago)](https://goreportcard.com/report/github.com/SerhiiCho/timeago)
![GitHub](https://img.shields.io/github/license/SerhiiCho/timeago)

Fast and lightweight datetime package that converts given datetime into "n time ago" format. The list of supported languages you can find [here](#-supported-languages).

- [⚙️ Configurations](https://github.com/SerhiiCho/timeago/blob/master/docs/CONFIGURATIONS.md)
- [🚩 Supported languages](#-supported-languages)
- [👏 Usage](#-usage)
- [🤲 Options](https://github.com/SerhiiCho/timeago/blob/master/docs/OPTIONS.md)
- [🇸🇿 Contribute translation](https://github.com/SerhiiCho/timeago/blob/master/docs/CONTRIBUTE_TRANS.md)
- [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/docs/CHANGELOG.md)
- [🚀 Quick Start](#-quick-start)
<!-- - [📖 Example usage on repl.it](https://repl.it/@SerhiiCho/Usage-of-timeago-package) -->

## 🚩 Supported languages

| Flag | Language | Short representation |
| --- | --- | --- |
| 🇬🇧 | English | en |
| 🇷🇺 | Russian | ru |
| 🇺🇦 | Ukrainian | uk |

## 👏 Usage

Pass the date to `timeago.Parse()` function. It counts the interval between current datetime and given datetime and returns parsed string in format `x time ago`.

Method `timeago.Parse` excepts different types of datetime:

- `int` Unix timestamp
- `time.Time` Type from Go time package
- `string` Datetime string in format `YYYY-MM-DD HH:MM:SS`

> Any other type will trigger a panic.

```go
timeago.Parse("2019-10-23 10:46:00") // string date
timeago.Parse(time.Now()) // time.Time
timeago.Parse(1642607826) // Unix timestamp
```

## 🚀 Quick Start

```bash
go get -u github.com/SerhiiCho/timeago
```