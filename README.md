![clothing shop](https://github.com/SerhiiCho/timeago/blob/master/.github/example.png?raw=true)

[![Code Coverage](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)
[![Build Status](https://travis-ci.com/SerhiiCho/timeago.svg?branch=master)](https://travis-ci.com/SerhiiCho/timeago)
[![Build Status](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/build.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/build-status/master)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/SerhiiCho/timeago/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/SerhiiCho/timeago/?branch=master)

100% covered with tests, fast and lightweight datetime converter into "n time ago" format from [year-month-day hours:minutes:seconds]. Supports Russian and English languages.

## Configurations
##### Language
Default language is English. Optionally you can set the language in your application by calling `timeago.Set()` function and pass the flag "ru" or "en".

```go
timeago.Set("language", "ru")
```

##### Location
Default location is Europe/Kiev. Optionally you can set the location in your application by calling `timeago.Set()` function and pass the location you need.

```go
timeago.Set("location", "India/Delhi")
```

## Usage

For outputting post publishing date or something else you can just pass the date to method `timeago.Take()`. It will count the interval between now and given date and returns converted format.

```go
timeago.Take("2019-10-23 10:46:00") // after 10 seconds outputs: 10 seconds ago
```

If you want to show last user login like if user is online or not, you can optionally add `|online` to the datetime string. All it does is just displaying **Online** if date interval withing 60 seconds.

```go
timeago.Take("2019-10-23 10:46:00|online")
```

## Quick Start

```bash
go get github.com/SerhiiCho/timeago
```
