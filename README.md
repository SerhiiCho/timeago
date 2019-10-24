![clothing shop](https://github.com/SerhiiCho/timeago/blob/master/.github/example.png?raw=true)

100% covered with tests, fast and lightweight datetime converter into "n time ago" format from [year-month-day hours:minutes:seconds]. Supports Russian and English languages.

## Configurations
##### Language
Default language is English. Optionally you can set the language in your application by calling `timeago.Set()` function and pass the flag "ru" or "en".

```go
timeago.Set("language", "ru")
```

##### Location
Default location is UTC. Optionally you can set the location in your application by calling `timeago.Set()` function and pass the location you need.

```go
timeago.Set("location", "America/New_York")
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
