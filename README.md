# Power Play

![image](https://github.com/jak103/powerplay/assets/16627408/4ec3df62-d760-40c6-aa57-fa63eaaaf61b)


[![Go](https://github.com/jak103/powerplay/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/jak103/powerplay/actions/workflows/go.yml)

## Table of Contents
- [Adding and Updating API Endpoints](#adding-and-updating-api-endpoints)

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

## Adding and Updating API Endpoints

### Ensure Model is Created and Up to Date
1. Navigate to the models directory: \
   `backend/internal/models`

### Add API Endpoint
1. Navigate to the server directory: \
    `backend/internal/server/apis`

### Add DB methods used in API endpoint
1. Navigate to the db directory: \
    `backend/internal/db`

### Update Open API docs**
1. Navigate to the open api spec directory: \
   `static/oas/v1`