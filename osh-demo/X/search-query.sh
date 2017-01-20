#!/bin/bash
set -e
. setup.sh

check_arg $1 term

# tag::public[]
term=$1

query='{
    "query": {
        "match": { "_all": "'"$term"'" }
    }
}'

$curl -X POST -d "$query" $PZSERVER/data/query?perPage=100&page=0
# end::public[]
