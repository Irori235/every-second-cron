# Every Second Cron Sample
The trigger app sends an HTTP request to the target app every second. 

It runs as a Cron job, scheduled to run every minute. When a new job starts, the old job is stopped using Redis Pub/Sub.

## Setup Kind & Taskfile
```
brew install kind
brew install go-task
```

## Quick Start
1. Cluster Start
```
task up
```
2. Image Build
```
task build
```
3. k8s apply
```
task apply
```

## Check Log
```
kubectl logs [target-pod]
```

following logs:
```
/ GET 200 2024-02-20T09:00:40Z
/ GET 200 2024-02-20T09:00:41Z
/ GET 200 2024-02-20T09:00:42Z
/ GET 200 2024-02-20T09:00:43Z
/ GET 200 2024-02-20T09:00:44Z
/ GET 200 2024-02-20T09:00:45Z
/ GET 200 2024-02-20T09:00:46Z
/ GET 200 2024-02-20T09:00:47Z
/ GET 200 2024-02-20T09:00:48Z
/ GET 200 2024-02-20T09:00:49Z
/ GET 200 2024-02-20T09:00:50Z
...
```
