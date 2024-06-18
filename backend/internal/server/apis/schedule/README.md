# Welcome to the schedule API!

## Overview
This API is responsible for managing the schedule of the games per season.

## Directory Structure
- `schedule.go` - Contains the logic for the schedule API.
- `internal` - Contains the internal logic for the schedule API.
  - `algorithms` - Contains the algorithms for generating games.
  - `analysis` - Contains the logic for analyzing the schedule. We use this to be able to tell if we should keep optimizing the schedule created from the algorithm.
  - `optimize` - Contains the logic for optimizing the schedule. After the algorithm creates the schedule, we can optimize it to make it better.
  - `structures` - Contains the structures for the schedule API.
  - `test_input` - Contains the test input for the schedule API.

## Authors
- Summer 2024 Semester
  - Eli Peterson
  - Zane Hirning
  - Marcus Quincy
  - Nate Stott
