#!/bin/sh
set -euo pipefail

for i in pz-dispatcher pz-registry svc-sleeper
do
    go build ./$i
done
