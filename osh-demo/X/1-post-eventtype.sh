#!/bin/bash
set -e
. setup.sh

eventtype='{
    "name": "test-'"$(date +%s)"'",
    "mapping": {
        "Unit": "string",
        "Value": "integer"
    }
}'

$curl -X POST -d "$eventtype" $PZSERVER/eventType
