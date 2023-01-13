# Advent of Code 2022

This is my implementation of [Advent of Code 2022](https://adventofcode.com/2022), in [Go](https://go.dev/).

You can run it with

```
go run . [ day [ star ] ]
```

where `day` is between 1 and 25 and `star` can be 1 or 2 for first or second star of the day. Without parameters, all stars are executed.

Unit tests are available, based on examples from the descriptions. You can run them with

```
go test ./days/...
```

to run all tests, or

```
go test ./days/dayXX
```

where `XX` is between `01` and `25`, to run tests for a single day. You can add verbosity with `-v` option.

All solutions are generic, that is you can replace inputs of my session with yours and you should get the correct answers.
