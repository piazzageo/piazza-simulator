#!/bin/bash
set -e
. setup.sh

check_arg $1 dataId
check_arg $2 filename

# tag::public[]
dataId=$1
filename=$2

$curl -X GET $PZSERVER/file/$dataId?fileName=terrametrics.tif > $filename
# end::public[]
