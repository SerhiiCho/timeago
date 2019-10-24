<h2 align="center">TimeAgo</h2>
<h3 align="center">Datetime converter</h3>

![clothing shop](https://github.com/SerhiiCho/timeago/blob/master/.github/example.png?raw=true)

Date/time converter into "n time ago" format. Supports Russian and English languages.

## Example

Default language is English. Optionally you can set the language in your application by calling `timeago.SetLang()` method and passing flag "ru" or "en";

```go
timeago.SetLang("ru")
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
go get SerhiiCho/timeago
```
