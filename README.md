[![Code Coverage](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FSerhiiCho%2Ftimeago%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/SerhiiCho/timeago/goto?ref=master)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/SerhiiCho/timeago)](https://goreportcard.com/report/github.com/SerhiiCho/timeago)
![GitHub](https://img.shields.io/github/license/SerhiiCho/timeago)

Fast and lightweight datetime converter that converts given datetime into "n time ago" format from [YEAR-MONTH-DAY HOURS:MINUTES:SECONDS]. Supports Russian and English languages.

- Chapters
    - [⚙️ Configurations](#%EF%B8%8F-configurations)
        - [Language](#language)
        - [Location](#location)
    - [🚩 Supported languages](#-supported-languages)
    - [👏 Usage](#-usage)
    - [🤲 Options](#-options)
    - [🇸🇿 Contribute translation](#-contribute-translation)
        - [Translation files](#translation-files)
        - [Rules](#rules)
    - [🚀 Quick Start](#-quick-start)
- Useful links
    - [📖 Example usage on repl.it](https://repl.it/@SerhiiCho/Usage-of-timeago-package)
    - [🗒 Release notes](https://github.com/SerhiiCho/timeago/blob/master/CHANGELOG.md)

## ⚙️ Configurations

#### Language

Default language is English. Optionally you can set the language in your application by calling `timeago.Set()` function and pass the flag "ru" or "en" in your init function.

```go
func init() {
    timeago.Set("language", "ru")
}
```

##### Location

Default location is Europe/Kiev. Optionally you can set the location in your application by calling `timeago.Set()` function and pass the location you need in your init function.

```go
func init() {
    timeago.Set("location", "America/New_York")
}
```

> Please make sure that timezone configuration is correct for your location. It is very important for displaying the correct datetime.

## 🚩 Supported languages

| Flag | Language | Short representation |
| --- | --- | --- |
| 🇬🇧 | English | en |
| 🇷🇺 | Russian | ru |
| 🇺🇦 | Ukrainian | uk |

## 👏 Usage

For outputting post publishing date or something else you can just pass the date to method `timeago.Conv()`. It will count the interval between now and given date and returns converted format. `timeago.Conv` method excepts 3 types: unix timestamp, string date and Time type from Go `time` package.

```go
timeago.Conv("2019-10-23 10:46:00") // string date
timeago.Conv(time.Now()) // time.Time type
timeago.Conv(1642607826) // Unix timestamp
```

## 🤲 Options

As the seconds argument `Conv()` method excepts strings. Here is an example of passed option.

```go
timeago.Conv(time.Now(), "online") // output: Online
timeago.Conv(time.Now()) // output: 0 seconds ago
```

#### Available options

| Option | Description |
| --- | --- |
| `online` | Displays **Online** if date interval withing 60 seconds. |

## 🇸🇿 Contribute translation

If you want to contribute support for a language that is fully supported, all you need to do is to copy/paste 2 files and change them to match the language that you want to add.

After than, add 1 line to `README.md` file and 1 rule to a `rules.go`. Here is my [commit](https://github.com/SerhiiCho/timeago/commit/d2f9e7f41d17ea3fc8ee10df2e1ac2e47f8e7e69) for supporting Ukrainian language that shows changes that I did to add the support. It's pretty straightforward. Waiting for you PR 😉.

### Translation files

Translation files live in `langs` directory. Each translation file is pretty simple json. Here's the example of `en.json`.

```json
{
    "Ago": "ago",
    "Online": "Online",
    "Second": "second",
    "Seconds": "seconds",
    "Minute": "minute",
    "Minutes": "minutes",
    "Hour": "hour",
    "Hours": "hours",
    "Day": "day",
    "Days": "days",
    "Week": "week",
    "Weeks": "weeks",
    "Month": "month",
    "Months": "months",
    "Year": "year",
    "Years": "years"
}
```

Some languages (like Russian) have multiple plural forms of the word. For example English has only `second` and `seconds`, but Russian language has 3 types `секунда`, `секунд` and `секунды`. For these cases we can add additional translation for seconds, minutes, hours, days, weeks, months and years. Here is the example of `ru.json`.

```json
{
    "Ago": "назад",
    "Online": "В сети",
    "Second": "секунда",
    "Seconds": "секунды",
    "SecondsSpecial": "секунд",
    "Minute": "минута",
    "Minutes": "минуты",
    "MinutesSpecial": "минут",
    "Hour": "час",
    "Hours": "часа",
    "HoursSpecial": "часов",
    "Day": "день",
    "Days": "дня",
    "DaysSpecial": "дней",
    "Week": "неделя",
    "Weeks": "недели",
    "WeeksSpecial": "недель",
    "Month": "месяц",
    "Months": "месяца",
    "MonthsSpecial": "месяцев",
    "Year": "год",
    "Years": "года",
    "YearsSpecial": "лет"
}
```

You can see that it has `SecondsSpecial`, `MinutesSpecial`, `HoursSpecial`, `DaysSpecial`, `WeeksSpecial` and `YearsSpecial` keys.

### Rules

All rules for each language is defined in `getRules` function in `rules.go` file. Rule is just a set of conditions that define when to apply singular form and when to apply plural form.

Here is the example for English rules:

```go
func getRules(number, lastDigit int) map[string]models.Rule {
	return map[string]models.Rule{
		"en": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
	}
}
```

We'll use singular form when number is equal to 1, and plural if number is more than 1 or number is 0. You can easily write your own rules for your language.

## 🚀 Quick Start

```bash
go get -u github.com/SerhiiCho/timeago
```
