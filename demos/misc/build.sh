#!/bin/sh

set -e

dir="github.com/venicegeo"

common="$dir/pz-gocommon"
logger="$dir/pz-logger"
uuidgen="$dir/pz-uuidgen"
alerter="$dir/pz-alerter"

for i in $common $logger $uuidgen $alerter
do
  go fmt $i
  go vet $i
  golint $i
  go doc $i > /dev/null
  go test $i
  go install $i
done

