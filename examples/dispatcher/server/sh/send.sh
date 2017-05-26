#!/bin/bash

for (( i=0; i<$1; i++ ))
do
    curl -i -d '{}' "http://localhost:9090/"
done