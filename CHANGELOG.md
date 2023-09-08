[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

# 🗒 Release Notes

----

## v2.1.7 (2023-09-08)

- Renamed master branch to main
- Added more tests
- Simplified code

----

## v2.1.6 (2022-11-27)

- Changed typo in README.md file;
- Added ability to overwrite output strings for custom values;

----

## v2.1.5 (2022-10-09)

- Code base refactoring. Made source code look nicer;

----

## v2.1.4 (2022-10-04)

- Code base refactoring. Made source code look nicer;

----

## v2.1.3 (2022-06-11)

- Documentation improvements and small changes;
- Fixed bug in test file `tests/utils.go` related to not properly counting months and years when testing;
- Added link to `README.md` file;

----

## v2.1.2 (2022-01-28)

- Added so that `Parse` function can except not only past date but also future date and return correct result. Closes #23;
- Added section `12 Features` to the `README.md`;

----

## v2.1.1 (2022-01-28)

- Added option `justNow` that prints `Just now` if time is within 60 minutes;
- Added option `noSuffix` that removes `ago` from the printed result;
- Added more info to `OPTIONS.md` documentation file;

----

## v2.1.0 (2022-01-27)

- Changed:
    - Renamed `Lang` structure to `lang` to make it private;
    - Renamed `Rule` structure to `rule` to make it private;
    - Changed location configurations. Now package can work without location configuration;
- Removed:
    - Removed badge from `README.md` file;
    - Removed tests from language files and added 1 test to `online_test.go` file;

----

## v2.0.4 (2022-01-22)

- Added:
    - Added anchors to "Contribute translation" docs;
    - Added 2 new GitHub badges;
    - Added more tests and test coverage;
- Changed:
    - Improved tests for languages by changing the way to write them;
- Fixed:
    - Fixed mistakes in Ukrainian language translation;

----

## v2.0.3 (2022-01-22)

- Added support for Dutch language;

----

## v2.0.2 (2022-01-21)

- Fixed typo in docs;
- Added more info to `contribute a translation guide`;

----

## v2.0.1 (2022-01-20)

- Fixed:
    - Fixed not working anchor tag in `OPTIONS.md` file;
- Added:
    - Added emoji to Release notes title in `CHANGELOG.md`;
    - Added `.gitattributes` file;
    - Added Example usage on `repl.it` website;
    - Added nice image banner to `README.md` file;

----

## v2.0.0 (2022-01-19)

- Added:
    - Added ability for `Parse` method to except unix timestamp;
    - Added ability for `Parse` method to except `Time` from Go time package;
- Changed:
    - Renamed `Take` method to `Parse`;
    - Changed the way you pass options;
    - Renamed `Set` method to `SetConfig`;
    - Changed the way you set configurations for the package;
- Documentation:
    - Added more information to docs;
    - Added `docs` directory with all the docs;
    - Changed structure of the `CHANGELOG.md`;
- Other:
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