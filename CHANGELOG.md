# Release Notes v3

## v3.1.0 (2024-11-25)
- Updated the `LICENSE.md` file
- Refactored codebase to make it more readable and maintainable

## v3.0.0 (2024-11-22)
> BREAKING CHANGES!

[Upgrade Guide from v2 to v3](https://time-ago.github.io/v3/upgrade.html)
- **Improved error handling**. The `Parse` function now returns an error as the second returned value
- **Update package namespace**. Changed package namespace to `github.com/SerhiiCho/timeago/v3`
- **Rename a function**. Renamed `SetConfig` function to `Configure` to make it better fit into Go naming conventions
- **New language files structure**. Change the file structure for `JSON` language files. They have now format to match [CLDR Specifications](https://cldr.unicode.org/index/cldr-spec/plural-rules)
- **New language addition is improved**. Add ability to change the output of the `Parse` function when adding a support for a new language by adding a `format` field to a `JSON` file
- **New function `Reconfigure`**. Added a new function to reconfigure the package configurations. Unlike the `Configure` function, it will overwrite the previous configurations with the new ones.

> Follow the [Upgrade Guide](https://time-ago.github.io/upgrade.html) from v2 to v3


## [Release Notes v2](.github/CHANGELOGV2.md)
## [Release Notes v1](.github/CHANGELOGV1.md)