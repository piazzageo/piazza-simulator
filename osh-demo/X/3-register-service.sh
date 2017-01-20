#!/bin/bash
set -e
. setup.sh

# tag::public[]
hello=`echo $PZSERVER | sed -e sXpiazzaXhttp://pzsvc-helloX`

service='{
    "url": "'"$hello"'",
    "contractUrl": "http://helloContract",
    "method": "GET",
    "isAsynchronous": "false",
    "resourceMetadata": {
        "name": "pzsvc-hello service",
        "description": "Hello World Example",
        "classType": {
           "classification": "UNCLASSIFIED"
        }
    }
}'

$curl -X POST -d "$service" $PZSERVER/service
# end::public[]
