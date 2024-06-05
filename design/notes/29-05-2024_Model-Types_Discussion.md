# Model-Types discussion

- Date: 5/29/2024
- Present: Skyler Cain, Ethan Christensen

Skyler and Ethan met over Zoom to clarify the model design.

Ethan was working on TypeScript types, and Skyler had been working with Dr. Christensen on the Go/DB models recently.

- [Model-Types discussion](#model-types-discussion)
  - [Questions and Answers](#questions-and-answers)
  - [Refining the Models](#refining-the-models)
    - [ShotOnGoal Model](#shotongoal-model)
    - [AuditLog Model](#auditlog-model)
    - [Goal Model](#goal-model)
    - [Penalty Model](#penalty-model)
    - [Game Model](#game-model)
    - [Team](#team)
  - [Action Items](#action-items)

## Questions and Answers

Some key questions were asked:

- What is the AuditLogEntry model?
  - The audit log entry table will keep a record of every change made to the database, who made it, and what they changed.

- What type is going to be used for duration?
  - This is still somewhat undecided. `time.Duration` gets cast to a uint64 when marshalled. By default, this number represents a duration in *nanoseconds* (Citation needed). We would like to have more control over the granularity it represents.

- How does DbModel get passed in each json response?
  - The fields of DbModel *spread out* into the struct it's inside of. It's similar in concept to an "extends" or "inherits" relation. Here's an example:

    ```go
    type SimpleDbModel struct {
      RowId uint `json:"id" gorm:"primarykey"`
      CreatedAt time.Time `json:"created_at"`
    }
    
    type UserA struct {
      SimpleDbModel
      Name: string `json:"string"`
    }
    
    
    type UserB {
      RowId uint `json:"id" gorm:"primarykey"`
      CreatedAt time.Time `json:"created_at"`
      Name: string `json:"string"`
    }
    
    // UserA is structurally equivalent to UserB
    ```

  UserA and UserB would marshall to this TypeScript type:

    ```ts
    interface Game {
      id: number,
      created_at: string,
      name: string
    }
    ```

- What is the CorrelationID in the Team and League models?
  - It's supposed to be a way to track the same "team" or "league" across multiple iterations of the team. Teams can change over time, but we want to keep records of what they were in previous seasons. Each team/league gets a guid that will uniquely identify it across all seasons.

- Referring to the previous question, is there a better name than Correlation for this concept?
  - This was tricky to answer. More on that later.

## Refining the Models

Some changes were suggested:

### ShotOnGoal Model

- Typo in tag of Scorekeeper! Add a close quote after `json:"scorekeeper`.

### AuditLog Model

- Perhaps this should include a DbModel?
  - Remove ID key if we add DbModel?

### Goal Model

- Change the JSON key of PlayerDifferential from "playerdifferential" to "player_differential".
- Change the JSON key of IsPenaltyShot from "ispenaltyshot" to "is_penalty_shot"

### Penalty Model

- Duration is used in both PenaltyType and Penalty, but mean different things.
  - Solution: rename Penalty.Duration to Penalty.GameTime

### Game Model

- The JSON structure can be clarified.
  
  ```go
  type Game struct {
    // ...
    HomeTeam Team `json:"home_team"`
    HomeTeamId uint `json:"home_team_id"`
    HomeTeamRoster Roster `json:"home_team_roster"`
    // ...
    AwayTeam Team `json:"away_team"`
    AwayTeamId uint `json:"away_team_id"`
    AwayTeamRoster Roster `json:"away_team_roster"`
    // ...
  }
  ```
  
  Could become:
  
  ```go
  type Game struct {
    // ...
    HomeTeam Team `json:"home.team"`
    HomeTeamId uint `json:"home.team_id"`
    HomeTeamRoster Roster `json:"home.team_roster"`
    // ...
    AwayTeam Team `json:"away.team"`
    AwayTeamId uint `json:"away.team_id"`
    AwayTeamRoster Roster `json:"away.team_roster"`
    // ...
  }
  ```
  
  Which would look like this in TypeScript:
  
  ```ts
  interface Game {
    // ...
    home: {
      team: Team,
      team_id: number,
      team_roster: Roster
      // ...
    },
    
    away: {
      team: Team,
      team_id: number,
      team_roster: Roster
      // ...
    }
    // ...
  }

### Team

(suggestions here also apply to League)

- The strategy of tracking a team over time by an ID is great: we should keep that.
- Rename the Team table to TeamRecord
  - TeamRecord more accurately represents what the data in the table is. It's a snapshot of a team in history, kept as a record.
- Rename CorrelationId to TeamGUID
  - This lets us *correlate* a TeamRecord with a Team
  - TeamId was an alternative idea, but we figured it might be confusing. TeamRecord.TeamId could be confused with TeamRecord.ID and TeamRecordId, and vice versa.

## Action Items

- [] Ethan: Make notes of meeting to remember our ideas
- [] Skyler: Make branch (A) with changes to go models
- [] Ethan: Make branch (B) with changes to ts models
- [] Ethan: Pull request branch B into ~~branch A~~ main
- [] Skyler: Pull request branch A into main
- [] (Anyone): Research having shared ts type definitions in a monorepo.
