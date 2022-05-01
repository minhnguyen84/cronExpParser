Cron expression parser
=============================

This programme parse a string cron expression to "Expression" will make this easy to exploit

Run
--------------
The expression have to be surrounded by simple quote
``` 
./cronExpParser '*/15 0 1,15 * 1-5 /usr/bin/find' 
```


String expression
--------------
`*/15 0 1,15 * 1-5 /usr/bin/find` 

It has six time fields (in order) : 

    Field name     Allowed values    Allowed special characters
    ----------     --------------    --------------------------
    Minutes        0-59              * / , -
    Hours          0-23              * / , -
    Day of month   1-31              * / , - 
    Month          1-12              * / , -
    Day of week    1-7               * / , - 

#### Asterisk ( * )
The asterisk indicates that the cron expression matches for all values of the field. E.g., using an asterisk in the 2nd field (hours) -> every hour from 0 to 24 

#### Slash ( / )
Slashes describe increments of ranges. For example `5/15` in the minute field indicate the 5th minute of the hour and every 15 minutes thereafter. `*` is mean the very first value.

#### Comma ( , )
Commas are used to separate items of a list. For example, using `1,15` in the 3rd field (day of month) means 1 and 15 of the month.

#### Hyphen ( - )
Hyphens define ranges. For example, `1-5` in the 5th field (day of week) indicates days number 1, 2, 3, 4 and 5 => Monday, Tuesday, Wednesday, Thursday and Friday.

---
We can combine `Hyphen( - )` and `Slash ( / )` by using Comma `( , )`. Exp : `5/15,20-28`

---
Every other pattern can not be parsed

Parser
--------------
Take a `string` and return an `Expression` if the cron string is valide

Expression
--------------

minute, hour, daysOfMonth, months and daysOfWeek is an array list of int `[]int` make the result easy to compute the next time

command is a string

It has a `ToString` who allow a pretty print of the object
