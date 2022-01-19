[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Language](#language)
- [Location](#location)

# ⚙️ Configurations

## Language

Default language is English. Optionally you can set the language in your application by calling `timeago.SetConfig()` function and pass the flag "ru" or "en" in your init function.

```go
func init() {
    timeago.SetConfig(timeago.Config{
        Language: "ru",
    })
}
```

## Location

Default location is Europe/Kiev. Optionally you can set the location in your application by calling `timeago.SetConfig()` function and pass the location you need in your init function.

```go
func init() {
    timeago.SetConfig("location", "America/New_York")
}
```

> Please make sure that timezone configuration is correct for your location. It is very important for displaying the correct datetime.

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)