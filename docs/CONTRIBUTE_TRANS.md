[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Description](#description)
- [Step 1. Add translation file](#step-1-add-translation-file)
- [Step 2. Add language rules](#step-2-add-language-rules)
- [Step 3. Add tests](#step-3-add-tests)
- [Step 4. Add info to README](#step-4-add-info-to-readme)

# 🇸🇿 Contribute translation

## Description

You can contribute a language support in 4 simple steps. All you need to do is to copy/paste 2 files and change them to match the language that you want to add.

Finally, add 1 line to `README.md` file and 1 rule to a `rules.go`. Here is my [commit](https://github.com/SerhiiCho/timeago/commit/c1ee0429b540f1cce5eb61b6a3441022d9cb43e7) for supporting Dutch language that shows changes that I did to add the support. It's pretty straightforward. Waiting for you PR 😉.

## Step 1. Add translation file

Translation files live in `langs` directory. Each translation file is pretty simple JSON object. Here's the example of `en.json`.

```json
{
    "Ago": "ago",
    "Online": "Online",
    "Second": "second",
    "Seconds": "seconds",
    "Minute": "minute",
    "Minutes": "minutes",
    "Hour": "hour",
    "Hours": "hours",
    "Day": "day",
    "Days": "days",
    "Week": "week",
    "Weeks": "weeks",
    "Month": "month",
    "Months": "months",
    "Year": "year",
    "Years": "years"
}
```

Some languages (like Russian) have multiple plural forms of the word. For example English has only `second` and `seconds`, but Russian language has 3 types `секунда`, `секунд` and `секунды`. For these cases we can add additional translation for seconds, minutes, hours, days, weeks, months and years. Here is the example of `ru.json`.

```json
{
    "Ago": "назад",
    "Online": "В сети",
    "Second": "секунда",
    "Seconds": "секунды",
    "SecondsSpecial": "секунд",
    "Minute": "минута",
    "Minutes": "минуты",
    "MinutesSpecial": "минут",
    "Hour": "час",
    "Hours": "часа",
    "HoursSpecial": "часов",
    "Day": "день",
    "Days": "дня",
    "DaysSpecial": "дней",
    "Week": "неделя",
    "Weeks": "недели",
    "WeeksSpecial": "недель",
    "Month": "месяц",
    "Months": "месяца",
    "MonthsSpecial": "месяцев",
    "Year": "год",
    "Years": "года",
    "YearsSpecial": "лет"
}
```

You can see that it has `SecondsSpecial`, `MinutesSpecial`, `HoursSpecial`, `DaysSpecial`, `WeeksSpecial` and `YearsSpecial` keys. Those are responsible for special age cases.

## Step 2. Add language rules

All rules for each language is defined in `getRules` function in `rules.go` file. Rule is just a set of conditions that define when to apply singular form and when to apply plural form.

Here is the example for English rules:

```go
func getRules(number, lastDigit int) map[string]Rule {
	return map[string]Rule{
		"en": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
	}
}
```

We'll use singular form when number is equal to 1, and plural if number is more than 1 or number is 0. You can easily write your own rules for your language.

## Step 3. Add tests

Tests for languages live in `tests` directory. Each language has it's own file. The easies way to add tests for your language is to copy paste one of the tests and change it to match your language.

## Step 4. Add info to README

Add a line to `README.md` file. `README.md` file has section ["Supported languages"](https://github.com/SerhiiCho/timeago#-supported-languages) where you need to add your language. For example for Dutch language I've added this line:

```
| 🇳🇱 | Dutch | nl |
```

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)