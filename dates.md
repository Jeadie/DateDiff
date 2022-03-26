# A delta of days

![https://xkcd.com/1883/](https://imgs.xkcd.com/comics/supervillain_plan.png)

## Description

We'd like to compute the number of days between two dates _from scratch_.
I.e. without importing or otherwise building on existing packages for date processing.

Despite the ominous XKCD comic above we don't need to worry about times or time zones for this problem.
Just the regular [calendar](https://en.wikipedia.org/wiki/Gregorian_calendar) and dates represented by strings of the form:
`YYYY-MM-DD`

For example, between `2012-01-01` and `2012-01-5` there are `3` days.

Note that start and end day should not be counted and we only care about the _absolute_ difference in dates.
If the order of two dates being compared is flipped the difference remains the same.

### Valid examples

`"2012-01-10"` <-> `"2012-01-11"` = `0` days

`"2012-01-01"` <-> `"2012-01-10"` = `8` days

`"1801-06-13"` <-> `"1801-11-11"` = `150` days

`"2021-12-01"` <-> `"2017-12-14"` = `1447` days

### Invalid examples

You should consider the possibility that malformed comparisons will be requested.
To simplify the problem we've constrained the definition of a calendar date to those which may be represented as: `YYYY-MM-DD`.

We leave it up-to you how these are handled or not by your application.

Examples of invalid dates:

- `"01-01-2020"`: not in the expected format..
- `"20000-01-01"`: invalid under our definition of allowed years (`YYYY`)..
- `"2022-02-29"`: format is not the only consideration for validity! 🤦‍♂️

## Solution shape

Construct an application that lets you answer a few arbitrary date comparison requests during the peer-review session.

For example, a command line application which takes two dates via stdin and returns the difference via stdout before exiting.

You're welcome to show off your skills by getting a little more fancy. Especially if you can demonstrate applicable technology and design choices relevant to your role.

A distributed machine learning solution running on the blockchain with virtual reality interface would be impressive; but overkill!
