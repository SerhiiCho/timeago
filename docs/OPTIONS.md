[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Usage](#usage)
- [Available options](#available-options)

# 🤲 Options

## Usage

As the seconds argument `Parse()` method excepts strings. Here is an example of passed option.

```go
timeago.Parse(time.Now(), "online") // Online
timeago.Parse(time.Now()) // 0 seconds ago
timeago.Parse(time.Now(), "justNow") // Just now
```

## Available options

| Option | Description |
| --- | --- |
| `online` | Displays **Online** if date interval withing 60 seconds. For example instead if `13 seconds ago` prints `Online` |
| `justNow` | Displays **Just now** if date interval withing 60 seconds. For example instead of `32 seconds ago` prints `Just now`. |

> More options are coming in the next versions.


[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)