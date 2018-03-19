#!/bin/bash

cd monitor
mvn clean install -Dhttp.proxyHost=www-proxy.idc.oracle.com -Dhttp.proxyPort=80 -Dhttps.proxyHost=www-proxy.idc.oracle.com -Dhttps.proxyPort=80
cd ..

