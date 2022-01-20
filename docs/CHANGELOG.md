[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

# Release Notes

----

## v2.0.1 (2022-01-21)

- Fixed
    - Fixed not working anchor tag in `OPTIONS.md` file;

----

## v2.0.0 (2022-01-19)

- Added
    - Added ability for `Parse` method to except unix timestamp;
    - Added ability for `Parse` method to except `Time` from Go time package;
- Changed
    - Renamed `Take` method to `Parse`;
    - Changed the way you pass options;
    - Renamed `Set` method to `SetConfig`;
    - Changed the way you set configurations for the package;
- Documentation
    - Added more information to docs;
    - Added `docs` directory with all the docs;
    - Changed structure of the `CHANGELOG.md`;
- Other
    - Refactored and rewritten code;

----

## v1.1.8 (2022-01-16)

- Improved documentation for the package.

----

## v1.1.7 (2022-01-16)

- Added caching the parsed results into memory to speed up the program. After this change, it will only parse json files once.

----

## v1.1.6 (2021-12-30)

- Changed type for the special rule in `rules.go` file.
- Improved docs by adding more information about how to contribute a language support.

----

## v1.1.5 (2021-12-22)

- Added support for Ukrainian language
- Improved documentation
- Made easy to contribute another language

----

## v1.1.4 (2021-12-05)

- Fixed mistake with failing test

----

## v1.1.3 (2021-12-04)

- Added ability to add rules to a language. We need it in order to make ability easily add a new translation for the package. Each language will have it's own set of rules.

----

## v1.1.2 (2021-11-28)

- Fixed bug with wrong path for plugin root

----

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)