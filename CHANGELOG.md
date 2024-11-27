# Release Notes v3

## v3.2.0 (2024-11-27)
- Added support for Belarusian language 🇧🇾
- Added support for Spanish language 🇪🇸
- Added support for Japanese language 🇯🇵
- Code refactoring and improvements for better readability and maintainability

## v3.1.0 (2024-11-25)
- Updated the `LICENSE.md` file
- Refactored codebase to make it more readable and maintainable
- Added [OnlineThreshold](https://time-ago.github.io/v3/configurations.html#thresholds) parameter to the configurations to set the threshold for the "Online" status
- Added [JustNowThreshold](https://time-ago.github.io/v3/configurations.html#thresholds) parameter to the configurations to set the threshold for the "Just now" status
- Added support for Chinese Simplified language 🇨🇳
- **POTENTIAL BREAK**. Might break your code if you were using the `Option` type or option constants directly. Which wasn't documented for public use.
    - Rename `Option` type to `opt` making it unexported
    - Add prefix to option constants `Opt` to make them like `OptOnline`, `OptNoSuffix`, etc.

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