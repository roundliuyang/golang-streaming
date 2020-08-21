#! /bin/bash

# Build web UI
cd ~/go/src/github.com/alanhou/golang-streaming/video_server/web/
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui/web
cp -R ~/go/src/github.com/alanhou/golang-streaming/video_server/templates ~/go/bin/video_server_web_ui/web