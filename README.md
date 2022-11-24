![golang](https://serhii.io/storage/other/timeago.png)

[![GoReportCard example](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/SerhiiCho/timeago)
[![Code Coverage](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FSerhiiCho%2Ftimeago%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/SerhiiCho/timeago/goto?ref=master)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![GitHub](https://img.shields.io/github/license/SerhiiCho/timeago)](https://github.com/SerhiiCho/timeago/blob/master/LICENSE.md)

Fast and lightweight date time package that converts given date into "n time ago" format. The package has many cool features and several supported languages.

- [😎 12 Features](#-12-features)
- [⚙️ Configurations](https://github.com/SerhiiCho/timeago/blob/master/docs/CONFIGURATIONS.md)
- [🚩 Supported languages](#-supported-languages)
- [👏 Usage](#-usage)
- [🤲 Options](https://github.com/SerhiiCho/timeago/blob/master/docs/OPTIONS.md)
- [🇸🇿 Contribute translation](https://github.com/SerhiiCho/timeago/blob/master/docs/CONTRIBUTE_TRANS.md)
- [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/docs/CHANGELOG.md)
- [🚀 Quick Start](#-quick-start)
- [📖 Example usage on repl.it](https://replit.com/@SerhiiCho/Usage-of-timeago-package)

## 😎 12 Features

- 🕐 Parses any given date, no matter it is the future date or the past;
- 🕑 Has several options that you can use depending on your use case;
- 🕒 Well tested;
- 🕓 Supports several languages;
- 🕔 Easy to contribute a new language support;
- 🕧 Small codebase;
- 🕖 Frequent small releases without breaking changes;
- 🕗 Can parse Unix timestamp;
- 🕘 Can parse date time string in `YYYY-MM-DD HH:MM:SS` format;
- 🕙 Can parse time from `time.Time` GO package;
- 🕚 All the changes and features are written in the [CHANGELOG.md](https://github.com/SerhiiCho/timeago/blob/master/docs/CHANGELOG.md);
- 🕛 Well documented package;

## 🚩 Supported languages

| Flag | Language | Code (ISO 639-1) |
| --- | --- | --- |
| 🇬🇧 | English | en |
| 🇷🇺 | Russian | ru |
| 🇺🇦 | Ukrainian | uk |
| 🇳🇱 | Dutch | nl |

## 👏 Usage

Pass the date to `timeago.Parse()` function. It counts the interval between current datetime and given datetime and returns parsed string in format `x time ago`. The package can work not only with dates in the past but future dates as well. The usage is pretty straight forward.

### Allowed types

Method `timeago.Parse()` excepts different types of datetime:

- `int` Unix timestamp
- `time.Time` Type from Go time package
- `string` Datetime string in format `YYYY-MM-DD HH:MM:SS`

> Any other type will trigger a panic.

```go
timeago.Parse("2019-10-23 10:46:00") // string date
timeago.Parse(time.Now()) // time.Time
timeago.Parse(1642607826) // Unix timestamp
```

### Usage with the date in the past

```go
pastDate := time.Now().Add(-time.Hour)

res := timeago.Parse(pastDate)

fmt.Println(res) // 1 hour ago
```

### Usage with the date in the future

```go
pastDate := time.Now().Add(time.Hour * 2)

res := timeago.Parse(pastDate)

fmt.Println(res) // 2 hours
```

## 🚀 Quick Start

```bash
go get -u github.com/SerhiiCho/timeago
```