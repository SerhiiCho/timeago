# Release Notes v4

## v4.0.0 (2026-11-25)

- 🐛 Fixed rounding up issues and inconsistencies, so everything is floored instead of rounded. Closes [#54](https://github.com/SerhiiCho/timeago/issues/54).
- ⚠️ **BREAKING CHANGE**. Remove support for Go version `1.20`.
- 🐛 Fixed month and year calculations to use calendar boundaries instead of fixed-duration approximations, including month-end and leap-year cases.
- ⚠️ **BREAKING CHANGE**. `Parse` now returns an error for `uint` and `uint64` timestamps greater than `math.MaxInt64` instead of silently wrapping to negative dates. Closes [#52](https://github.com/SerhiiCho/timeago/issues/52).

## [Release Notes v3](.github/CHANGELOGV3.md)
## [Release Notes v2](.github/CHANGELOGV2.md)
## [Release Notes v1](.github/CHANGELOGV1.md)
