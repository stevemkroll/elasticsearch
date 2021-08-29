## Setup Environment
* Install Docker Desktop [link](https://docs.docker.com/get-docker/)


## Seeding
* In the root directory of the project run the following commands:
```
docker compose up -d
go run main.go -service=seed
go run main.go -service=api
```

## Run Queries
* In a new terminal run the following curl requests:
```
// Look up employees by phone

curl -v -H "GovernMint-token: pa\$\$word" http://localhost:8080/employee\?phone\=1111111111

// Seeded phone numbers

1111111110
1111111111
1111111112
1111111113
1111111114
1111111115
1111111116
1111111117
1111111118
1111111119
```

```
// Look up employees by task name

curl -v -H "GovernMint-token: pa\$\$word" http://localhost:8080/task\?name\=task_3

// Seeding task names

task_0
task_1
task_2
task_3
task_4
task_5
task_6
task_7
task_8
task_9
```