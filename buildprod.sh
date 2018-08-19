#! /bin/bash

# Build web and other services

cd /home/zereker/Documents/Go/src/github.com/Zereker/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd /home/zereker/Documents/Go/src/github.com/Zereker/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd /home/zereker/Documents/Go/src/github.com/Zereker/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd /home/zereker/Documents/Go/src/github.com/Zereker/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web