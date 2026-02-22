# Welcome to the schedule API!

## Overview
This API is responsible for managing the schedule of the games per season.

## Directory Structure
- `schedule.go` - Contains the logic for the schedule API.
- `internal` - Contains the internal logic for the schedule API.
  - `algorithms` - Contains the algorithms for generating games. The default schedule is based on the round robin algorithm.
  - `analysis` - Contains the logic for analyzing the schedule. We use this to be able to tell if we should keep optimizing the schedule created from the algorithm.
  - `optimize` - Contains the logic for optimizing the schedule. After the algorithm creates the schedule, we can optimize it to make it better. There are two main goals of the optimizers. First is to make sure there is a good amount of time between games so a team doesn't play back-to-back days. Second is to try to balance early and late games so a team doesn't only have games that start at 22:00. There are two optimizers:
    - `pair_swap` - Pair swap works by iterating over all the games in the schedule, and swaps which teams play in adjacent games if it improves the early-late balance of those teams.
    - `set_swap` - Set swap works by looking at sets of 6 games at a time. It finds which team has the worst early-late balance and it swaps that game with the team that is imbalanced the other way.
  - `structures` - Contains the structures for the schedule API.
  - `test_input` - Contains the test input for the schedule API.

## Authors
- Summer 2024 Semester
  - Eli Peterson
  - Zane Hirning
  - Marcus Quincy
  - Nate Stott

## Reinforcement Learning Alternative
Work was done on a reinforcement learning model which is an alternative to the schedules generated here. That code is hosted at [this repository](https://github.com/marcus-quincy/schedule-rl).

We were unable to make the reinforcement learning approach perform better than the algorithms currently implemented, so it is currently not integrated into this application.
