# Uniq
Utility to find unique strings. Similar to [uniq](https://en.wikipedia.org/wiki/Uniq) Unix command line utility.

## Requirements

>To use this utility your need to [install Go](https://golang.org/doc/install)

## Usage

`go run main.go [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`

```
-c
    prefix lines by the number of occurrences

-d
    only print duplicate lines, one for each group

-f num
    avoid comparing the first num fields

-i
    ignore differences in case when comparing

-s chars
    avoid comparing the first chars characters

-u
    only print unique lines
```

If input_file is not passed, input stream is **stdin** (command line input)

If output_file is not passed, output stream is **stdout** (command line output)
