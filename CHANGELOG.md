# Release Notes v3

## v3.4.2 (2026-08-11)

- 🧑 Moved from GitHub to Codeberg
- ⚠️ **POTENTIAL BREAK**. Dropped Go versions support for 1.16, 1.17, 1.18 and 1.19

## v3.3.2 (2026-02-08)

- 🧑 Added linting and formatting check for GitHub actions.
- 🧑 Code refactoring and improvements for better readability and maintainability.

## v3.3.1 (2025-06-20)

- 🧑 Added Container Engines like Docker and Podman to the project.

## v3.3.0 (2025-06-16)

- ✨ Now you can parse different times of timestamps, such as `string`, `int64`, `uint`, `uint64`, `int32`, `uint32`.

## v3.2.2 (2025-06-04)

- ✨ Change so that all JSON files are included in the final binary by using `go:embed` functionality.
- ⚠️ Remove support for go version 1.13, 1.14, 1.15 because they do not have `go:embed`.

## v3.2.1 (2024-12-06)

- 🐛 Bug fix with incorrect Russian, Ukrainian, and Belarusian languages rules.
- 🧑 Added more tests for Russian, Ukrainian, and Belarusian languages for testing seconds.

## v3.2.0 (2024-11-27)

- ✨ Added support for Belarusian language 🇧🇾.
- ✨ Added support for Spanish language 🇪🇸.
- ✨ Added support for Japanese language 🇯🇵.
- ✨ Added support for French language 🇫🇷.
- 🧑 Code refactoring and improvements for better readability and maintainability.

## v3.1.0 (2024-11-25)

- 📝 Updated the `LICENSE.md` file
- 🧑 Refactored codebase to make it more readable and maintainable
- ✨ Added [OnlineThreshold](https://timeago.codeberg.page/v3/configurations.html#thresholds) parameter to the configurations to set the threshold for the "Online" status
- ✨ Added [JustNowThreshold](https://timeago.codeberg.page/v3/configurations.html#thresholds) parameter to the configurations to set the threshold for the "Just now" status
- ✨ Added support for Chinese Simplified language 🇨🇳
- ⚠️ **POTENTIAL BREAK**. Might break your code if you were using the `Option` type or option constants directly. Which wasn't documented for public use.
  - Rename `Option` type to `opt` making it unexported
  - Add prefix to option constants `Opt` to make them like `OptOnline`, `OptNoSuffix`, etc.

## v3.0.0 (2024-11-22)

> BREAKING CHANGES!

[Upgrade Guide from v2 to v3](https://timeago.codeberg.page/v3/upgrade.html)

- **Improved error handling**. The `Parse` function now returns an error as the second returned value
- **Update package namespace**. Changed package namespace to `codeberg.org/timeago/timeago/v3`
- **Rename a function**. Renamed `SetConfig` function to `Configure` to make it better fit into Go naming conventions
- **New language files structure**. Change the file structure for `JSON` language files. They have now format to match [CLDR Specifications](https://cldr.unicode.org/index/cldr-spec/plural-rules)
- **New language addition is improved**. Add ability to change the output of the `Parse` function when adding a support for a new language by adding a `format` field to a `JSON` file
- **New function `Reconfigure`**. Added a new function to reconfigure the package configurations. Unlike the `Configure` function, it will overwrite the previous configurations with the new ones.

> Follow the [Upgrade Guide](https://timeago.codeberg.page/upgrade.html) from v2 to v3

## [Release Notes v2](.forgejo/CHANGELOGV2.md)

## [Release Notes v1](.forgejo/CHANGELOGV1.md)
