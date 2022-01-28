[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Usage](#usage)
- [Available options](#available-options)

# 🤲 Options

## Usage

As the seconds argument `Parse()` method excepts strings. Here is an example of passed option.

```go
currentTime := time.Now()
hourAgo := currentTime.Add(-time.Hour)

timeago.Parse(currentTime) // 0 seconds ago
timeago.Parse(currentTime, "online") // Online
timeago.Parse(currentTime, "justNow") // Just now

timeago.Parse(hourAgo) // 1 hour ago
timeago.Parse(hourAgo, "noSuffix") // 1 hour
```

## Available options

| Option | Description |
| --- | --- |
| `online` | Displays **Online** if date interval withing 60 seconds. For example instead if `13 seconds ago` prints `Online` |
| `justNow` | Displays **Just now** if date interval withing 60 seconds. For example instead of `32 seconds ago` prints `Just now`. |
| `noSuffix` | Removes suffix from datetime result and get for example "5 minutes" instead of "5 minutes ago". |

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)