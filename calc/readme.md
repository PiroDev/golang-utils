# Calc
Simple calculator utility. Supports simple operations (**+ - * /**).

## Requirements

>To use this utility your need to [install Go](https://golang.org/doc/install)

## Usage

`go run main.go "expression"`

Expression can contains:

1. Float numbers (doesn't support unary operations like unary minus or plus (-num, +num))
2. Operations: **+ - * /**
3. Brackets: **( )**

Example expression:

```math
2 * ((3 + 5) * 10 - 35 / 23) * 33 + 12
```