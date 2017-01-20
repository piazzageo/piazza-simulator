#!/bin/bash
set -e

. setup.sh

#----------------------------------------------------------------------------

echo "POST /eventType"

eventtype='{
    "name": "test-'"$(date +%s)"'",
    "mapping": {
        "Unit": "string",
        "Value": "integer"
    }
}'

tmp=`$curl -X POST -d "$eventtype" $PZSERVER/eventType`


eventTypeId=`echo $tmp | jq -r .data.eventTypeId`
echo "  eventTypeId: $eventTypeId"

sleep 5

#----------------------------------------------------------------------------

echo "POST /event"

event='{
    "eventTypeId": "'"$eventTypeId"'",
    "data": {
        "Unit": "degrees",
        "Value": XXX
    }
}'

count=0
while [ true ]; do
    #echo $count

    curevent=`echo $event | sed -e s/XXX/$count/`
    $curl -X POST -d "$curevent" $PZSERVER/event
    echo
    
    sleep 2
    let count=count+1
done

echo "FAIL - notreached"
exit 1
