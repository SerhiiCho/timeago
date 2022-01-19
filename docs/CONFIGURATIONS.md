[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Description](#description)
- [Language](#language)
- [Location](#location)

# ⚙️ Configurations

## Description

We can set package configurations with `SetConfig` function that accepts `Config` structure.

## Language

Optionally you can set the language in your application.

> Default value is `"en"`

```go
import . "github.com/SerhiiCho/timeago"

func main() {
    SetConfig(Config{
        Language: "ru",
    })
}
```

## Location

Optionally you can set the location in your application.

> Default value is `"America/New_York"`

```go
import . "github.com/SerhiiCho/timeago"

func main() {
    SetConfig(Config{
        Location: "Europe/Kiev",
    })
}
```

> Please make sure that timezone configuration is correct for your location. It is very important for displaying the correct datetime.

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)