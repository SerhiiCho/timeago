[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)

- [Description](#description)
- [Translation files](#translation-files)
- [Rules](#rules)

# 🇸🇿 Contribute translation

## Description

If you want to contribute support for a language that is fully supported, all you need to do is to copy/paste 2 files and change them to match the language that you want to add.

After than, add 1 line to `README.md` file and 1 rule to a `rules.go`. Here is my [commit](https://github.com/SerhiiCho/timeago/commit/d2f9e7f41d17ea3fc8ee10df2e1ac2e47f8e7e69) for supporting Ukrainian language that shows changes that I did to add the support. It's pretty straightforward. Waiting for you PR 😉.

## Translation files

Translation files live in `langs` directory. Each translation file is pretty simple json. Here's the example of `en.json`.

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

You can see that it has `SecondsSpecial`, `MinutesSpecial`, `HoursSpecial`, `DaysSpecial`, `WeeksSpecial` and `YearsSpecial` keys.

## Rules

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

[<< Go back to home](https://github.com/SerhiiCho/timeago/blob/master/README.md)