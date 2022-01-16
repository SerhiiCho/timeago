[![Code Coverage](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FSerhiiCho%2Ftimeago%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/SerhiiCho/timeago/goto?ref=master)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/SerhiiCho/timeago)](https://goreportcard.com/report/github.com/SerhiiCho/timeago)
![GitHub](https://img.shields.io/github/license/SerhiiCho/timeago)

Fast and lightweight datetime converter that converts given datetime into "n time ago" format from [YEAR-MONTH-DAY HOURS:MINUTES:SECONDS]. Supports Russian and English languages.

- [⚙️ Configurations](https://github.com/SerhiiCho/timeago/blob/master/docs/CONFIGURATIONS.md)
- [🚩 Supported languages](#-supported-languages)
- [👏 Usage](#-usage)
- [🇸🇿 Contribute translation](https://github.com/SerhiiCho/timeago/blob/master/docs/CONTRIBUTE_TRANS.md)
- [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/docs/CHANGELOG.md)
- [🚀 Quick Start](#-quick-start)
- [📖 Example usage on repl.it](https://repl.it/@SerhiiCho/Usage-of-timeago-package)

## 🚩 Supported languages

<table>
  <thead>
    <tr>
      <th>FLag</th>
      <th>Language</th>
      <th>Short representation</th>
    </tr>
  </thead>
  <tbody>
     <tr>
      <td>🇬🇧</td>
      <td>English</td>
      <td>en</td>
    </tr>
    <tr>
      <td>🇷🇺</td>
      <td>Russian</td>
      <td>ru</td>
    </tr>
    <tr>
      <td>🇺🇦</td>
      <td>Ukrainian</td>
      <td>uk</td>
    </tr>
  </tbody>
</table>

## 👏 Usage

For outputting post publishing date or something else you can just pass the date to method `timeago.Take()`. It will count the interval between now and given date and returns converted format.

```go
timeago.Take("2019-10-23 10:46:00") // after 10 seconds outputs: 10 seconds ago
```

If you want to show last user login like if user is online or not, you can optionally add `|online` to the datetime string. All it does is just displaying **Online** if date interval withing 60 seconds.

```go
timeago.Take("2019-10-23 10:46:00|online")
```

## 🚀 Quick Start

```bash
go get -u github.com/SerhiiCho/timeago
```