#!/bin/bash
set -e
. setup.sh

check_arg $1 eventTypeId
check_arg $2 serviceId

eventTypeId=$1
erviceId=$2

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

$curl -X POST -d "$trigger" $PZSERVER/trigger

