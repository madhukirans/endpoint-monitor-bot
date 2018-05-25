#!/bin/bash

cd monitor
#clean
mvn install
if docker build -t madhukirans/endpoint-monitor-bot:1.0 . ; then
  docker push madhukirans/endpoint-monitor-bot:1.0
fi
cd ..

