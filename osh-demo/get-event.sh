#!/bin/bash
set -e
. setup.sh

check_arg $1 eventId

eventId=$1

$curl -X GET $PZSERVER/event/$eventId
