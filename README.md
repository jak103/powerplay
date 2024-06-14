# Power Play

![image](https://github.com/jak103/powerplay/assets/16627408/4ec3df62-d760-40c6-aa57-fa63eaaaf61b)


[![Go](https://github.com/jak103/powerplay/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/jak103/powerplay/actions/workflows/go.yml)
[![Site Node.js CI](https://github.com/jak103/powerplay/actions/workflows/node.js.yml/badge.svg)](https://github.com/jak103/powerplay/actions/workflows/node.js.yml)

## Table of Contents
- [Team Specific Pages](#team-specific-pages)
- [Helpful Commands](#helpful-commands)
    - [Clear and Re-run Migrations](#clear-and-re-run-migrations)
    - [Run All Go Tests](#run-all-go-tests)

## Team Specific Pages
- [Frontend Site](/site/README.md)
- [Frontend App](/app/README.md)
- [Chat Functions]()
- [Scoreboard]()
- [Scheduling]()
- [Backend](/backend/README.md)

## Helpful Commands
### Clear and Re-run Migrations
Sometimes you need to clear out migrations due to a model change. 
While we are in early development we've decided to drop and recreate
migrations vs. source control and continuously run them. You can easily
drop the migrations table and have go auto migrate any changes.

** Note this will not drop columns / tables  
** This needs to be ran while docker is running

```shell
make nuke-migrations
```

### Run all Go tests 
Ability to quickly run all go tests to ensure your changes
pass all tests before committing code.

** This needs to be ran while docker is running

```shell
make test 
```
