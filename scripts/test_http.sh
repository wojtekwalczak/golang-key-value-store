#!/bin/bash

echo Checking version of the service:
curl http://localhost:8080/
echo

echo Putting key-value pair into the store
curl -X PUT -d 'value-1' -v http://localhost:8080/key-1
echo

echo Fetching the value for key
curl http://localhost:8080/key-1
echo
