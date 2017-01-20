#!/bin/bash
set -e

. setup.sh

echo "post-eventtype.sh"
eventTypeId=`sh post-eventtype.sh | jq -r .data.eventTypeId`
echo "    eventTypeId: $eventTypeId"

echo "register-service.sh..."
serviceId=`sh register-service.sh  | jq -r .data.serviceId`
echo "    serivceId: $serviceId"

echo "post-trigger.sh"
triggerId=`sh post-trigger.sh $eventTypeId $serviceId | jq -r .data.triggerId`
echo "    triggerId: $triggerId"

echo "post-event.sh"
eventId=`sh post-event.sh $eventTypeId | jq -r .data.eventId`
echo "    eventId: $eventId"

sleep 3

echo "get-alerts.sh"
result=`sh get-alerts.sh $triggerId`

resultTriggerId=`echo $result | jq -r .data[0].triggerId`
resultEventId=`echo $result | jq -r .data[0].eventId`

echo "    resultTriggerId: $resultTriggerId"
if [ "$triggerId" != "$resultTriggerId" ]
then
    echo FAIL
    exit 1
fi

echo "    resultEventId: $resultEventId"
if [ "$eventId" != "$resultEventId" ]
then
    echo FAIL
    exit 1
fi

echo PASS Test7
