[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Usage](#usage)
- [Available options](#availableoptions)

# 🤲 Options

## Usage

As the seconds argument `Conv()` method excepts strings. Here is an example of passed option.

```go
timeago.Conv(time.Now(), "online") // output: Online
timeago.Conv(time.Now()) // output: 0 seconds ago
```

## Available options

| Option | Description |
| --- | --- |
| `online` | Displays **Online** if date interval withing 60 seconds. |


[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)