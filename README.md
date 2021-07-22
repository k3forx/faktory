# faktory

## Set up docker container for faktory

```bash
❯ docker-compose up -d

❯ docker container ls
CONTAINER ID   IMAGE                      COMMAND                  CREATED          STATUS          PORTS                                                           NAMES
9f443577f769   contribsys/faktory:1.5.3   "/faktory -w 0.0.0.0…"   43 seconds ago   Up 41 seconds   0.0.0.0:7419-7420->7419-7420/tcp, :::7419-7420->7419-7420/tcp   faktory
```

Access to [localhost:7420](localhost:7420)

![image](https://user-images.githubusercontent.com/45956169/126645569-07719fc5-738e-4508-acab-fb39a6366734.png)

## Push a new job to the server

```bash
❯ go run ./main.go
Pushing my first job
Job is pushed
```

## Listen the job
