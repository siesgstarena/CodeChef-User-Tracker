# CodeChef-Track-User
A web app written in Go to keep track of the arena members for their participation in official codechef contest which will help the guy to keep track easily.

Link https://arena-codechef-tracker.herokuapp.com/

#### How it works
see the .example.data.xlsx file make sure to have the same column with Name and CodeChef in the CodeChef column the CodeChef Id of members should be there.

open the application enter the dates and upload the file a new file will be downloaded with the 2 additional column one contains all the contest code and the another contain the total of contest.
#### Get Started
copy the .env.exmaple to .env

```
  cp .\.env.example .\.env
```

To Install dependencies

```
    go mod download
```

To run 

```
    air
```

Rus using docker
```
  docker compose up
```