#!/bin/bash

cd monitor
#clean
mvn install
if docker build -t madhukirans/monitor:1.0 . ; then
  docker push madhukirans/monitor:1.0
fi
cd ..

