#!/bin/bash
set -e

. setup.sh

check_arg $1 triggerId
triggerId=$1

echo "GET /alert"

out=`$curl -X GET $PZSERVER/alert?triggerId=$triggerId&perPage=1 ` #| jq -r .`
echo $out
exit


count=`echo $out | jq -r .pagination.count`
resultTriggerId=`echo $out | jq -r .data[0].triggerId`
#val=`echo $result | jq -r .data[0]`
resultEventId=`echo $out | jq -r .data[0].eventId`

echo "    count: $count"
