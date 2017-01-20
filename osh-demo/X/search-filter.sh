#!/bin/bash
set -e
. setup.sh

check_arg $1 term

# tag::public[]
term=$1

$curl -X GET $PZSERVER/data?keyword=$term
# end::public[]
