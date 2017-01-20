#!/bin/bash
set -e

. setup.sh

check_arg $1 eventTypeId
eventTypeId=$1

#----------------------------------------------------------------------------

echo "POST /service"

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

out=`$curl -X POST -d "$service" $PZSERVER/service`

serviceId=`echo $out | jq -r .data.serviceId`
echo "    serivceId: $serviceId"

#----------------------------------------------------------------------------

echo "POST /trigger"

#"match": {"data.Severity": 200}

trigger='{
    "name": "MyTrigger",
    "eventTypeId": "'"$eventTypeId"'",
    "condition": {
        "query": { "query": { "match_all": {} } }
    },
    "job": {
        "userName": "test",
        "jobType": {
            "type": "execute-service",
            "data": {
                "serviceId": "'"$serviceId"'",
                "dataInputs": {},
                "dataOutput": [
                    { 
                        "mimeType": "application/json", 
                        "type": "text"
                    }
                ]
            }
        }
    },
    "enabled": true
}'

out=`$curl -X POST -d "$trigger" $PZSERVER/trigger`

triggerId=`echo $out | jq -r .data.triggerId`
echo "    triggerId=$triggerId"
