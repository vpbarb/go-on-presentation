#!/bin/bash

i=1
while [ $i -le $1 ]; do
    payload="{\"key\":\"value_$i\"}"
    curl -i -d $payload "http://localhost:9090/"
    let i+=1
done