#!/bin/bash
set -e
. setup.sh

check_arg $1 eventTypeId

eventTypeId=$1

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
