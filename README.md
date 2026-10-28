![golang](https://codeberg.org/timeago/timeago/raw/branch/master/.forgejo/header.png)

Timeago is a fast and lightweight date time package that converts given date into human readable "n time ago" format in different languages, such as 🇬🇧 🇷🇺 🇺🇦 🇳🇱 🇩🇪 🇨🇳 🇧🇾 🇪🇸 🇯🇵 🇫🇷. For more information you can read the [documentation](https://time-ago.github.io/).

> [!IMPORTANT]
> The project moved from `github.com/SerhiiCho/timeago` to `codeberg.org/timeago/timeago`. See [Migrating from GitHub](#migrating-from-github) below.

### Follow the [Official documentation](https://timeago.serhiicho.com/) for all the details

- [🗒 Documentation](https://timeago.serhiicho.com/)
- [🗒 Release notes](https://codeberg.org/timeago/timeago/src/branch/master/CHANGELOG.md)

## Quick Start

```bash
go get codeberg.org/timeago/timeago/v3
```

## Migrating to Codeberg

1. Replace the import path in your code:

   ```bash
   grep -rl 'github.com/SerhiiCho/timeago' --include='*.go' . | xargs sed -i 's#github.com/SerhiiCho/timeago#codeberg.org/timeago/timeago#g'
   ```

2. Fetch the new module and drop the old one:

   ```bash
   go get codeberg.org/timeago/timeago/v3
   go mod tidy
   ```

3. Check that everything still builds and passes:

   ```bash
   go build ./... && go test ./...
   ```

## License

This project is open-sourced software licensed under the [MIT license](https://codeberg.org/timeago/timeago/src/branch/master/LICENSE.md).
