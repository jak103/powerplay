# Power Play Backend

![image](https://github.com/jak103/powerplay/assets/16627408/4ec3df62-d760-40c6-aa57-fa63eaaaf61b)

[![Go](https://github.com/jak103/powerplay/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/jak103/powerplay/actions/workflows/go.yml)

## Table of Contents
- [Adding and Updating API Endpoints](#adding-and-updating-api-endpoints)
- [Adding and Updating Authentication](#adding-and-updating-authentication)
- [Adding and Updating Unit Testing for Database Model](#adding-and-updating-unit-testing-for-database-model)
- [Adding and Updating Unit Testing for API Endpoint](#adding-and-updating-unit-testing-for-api-endpoint)
## Adding and Updating API Endpoints
There are 5 major steps to creating or editing an endpoint.
- [Database Model](#ensure-model-is-created-and-up-to-date)
- [API Handlers](#add-or-edit-api-endpoint-handlers)
- [Database Methods](#add-or-edit-any-db-methods-used-in-the-api-endpoint)
- [Documentation](#update-the-open-api-docs)
- [Testing](#adding-and-updating-unit-testing-for-database-model)

### Ensure Model is Created and Up to Date
1. Navigate to the models directory:  
   `backend/internal/models`
2. Find the model for the endpoint you are working on.
   - If the model is not there, create it
   - If the model is not accurate, make changes to ensure the model will reflect the needs.

### Add or Edit API Endpoint Handlers
1. Navigate to the server directory:  
    `backend/internal/server/apis`
2. Create or find the .go file for the Endpoint you are working on.
3. Implement or edit the API Handlers. The most common handlers are GET and POST. 
   
#### **GET and POST Handlers**
The following links to the current penalty API handlers. Use this as an example for creating new handlers so everything in the backend follows a similar pattern.

- [Example Penalties API Handler](/backend/internal/server/apis/sports/stats/penalty.go)

4. Add in any handlers to the init function
``` go
func init() {
	apis.RegisterHandler(fiber.MethodGet, "/generics", auth.Public, getGenericsHandler)
	apis.RegisterHandler(fiber.MethodPost, "/generics", auth.Public, postGenericHandler)
}
```

### Add or Edit any DB methods used in the API endpoint
1. Navigate to the db directory:  
    `backend/internal/db`
2. Create or find the .go file for the endpoint you are working on.
3. Implement or edit the methods within the .go file. These can include Get, Create

#### **Get and Create Method**
The following link is to an example .go file containing methods to be used as an example in creating new methods to keep methods consistent throughout the backend. The preload can be used to load any needed database relation.

- [Example Get and Create Method](/backend/internal/db/penalty.go)

### Update the Open API docs
1. Navigate to the open api spec directory:  
   `static/oas/v1`
2. Create or update the corresponding .yml file to correctly reflect any changes you have made.

3. Use the following linked .yml file to keep the documentation consistent throughout the backend.
- [Example .yml file](/static/oas/v1/sports/stats/penalties.yml)

## Adding and Updating Authentication

### Generating Key

1. Install openssl
2. Generate secret key
   - `openssl rand -base64 32`
3. Add key to [local.env](../config/local.env) config

### Adding Auth to API Endpoints
There are several different roles that a user can be associated to:

Roles:
- None 
- Player
- Captain
- Referee
- ScoreKeeper
- Manager

There are roles that contain a set or group of roles:
| Role Group | Role |
| --- | --- | 
| Public | [None] |
| Authenticated | [Manager, Referee, ScoreKeeper, Captain, Player] | 
| Staff | [Manager, Referee, ScoreKeeper] | 
| ManagerOnly | [Manager] |

Each API endpoint will have a set of roles that will be allowed to hit that endpoint. Here are the endpoints with the allowed roles:

| Endpoint | HTTP Method | Roles |
| --- | --- | --- |
| /auth | POST | Public |
| /games | GET | Public |
| /games | POST | ManagerOnly | 
| /goals | GET | Public | 
| /goals | POST | Staff | 
| /leagues | GET | Authenticated | 
| /leagues | POST | ManagerOnly | 
| /logo | GET | Public | 
| /logo | POST | ManagerOnly | 
| /penalties | GET | Public | 
| /penalties | POST | Staff | 
| /penaltyTypes | GET | Public | 
| /rosters | GET | Authenticated | 
| /rosters | POST | ManagerOnly | 
| /seasons | GET | Authenticated | 
| /seasons | POST | ManagerOnly | 
| /shotsongoal | POST | Staff | 
| /teams | GET | Authenticated | 
| /teams | POST | ManagerOnly | 
| /user | GET  | Authenticated | 
| /user | POST | ManagerOnly |
| /venues | GET | Public |
| /venues | POST | ManagerOnly |




## Adding and Updating Unit Testing for Database Model
Adding a unit test for a database model. These tests will be using a docker spin up of the actual database for testing the database interfacing functions.

Uses the dockertest go package

To be expounded upon once a unit testing pattern is achieved.

## Adding and Updating Unit Testing for API Endpoint
Adding a unit test for an API endpoint. These tests will be using a mock up of the database models involved.

To be expounded upon once a unit testing pattern is achieved.
