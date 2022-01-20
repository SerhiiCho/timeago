[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Usage](#usage)
- [Available options](#available-options)

# 🤲 Options

## Usage

As the seconds argument `Parse()` method excepts strings. Here is an example of passed option.

```go
timeago.Parse(time.Now(), "online") // output: Online
timeago.Parse(time.Now()) // output: 0 seconds ago
```

## Available options

| Option | Description |
| --- | --- |
| `online` | Displays **Online** if date interval withing 60 seconds. |

> More options are coming in the next versions.


[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)