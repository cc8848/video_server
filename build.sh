#! /bin/bash

# Build web UI
cd /home/zereker/Documents/Go/src/github.com/Zereker/video_server/web
go install
cp /home/zereker/Documents/Go/bin/web /home/zereker/Documents/Go/bin/video_server_web_ui/
cp -R /home/zereker/Documents/Go/src/github.com/Zereker/video_server/templates /home/zereker/Documents/Go/bin/video_server_web_ui/

