#!/bin/sh

read -d '' regzoo << EOF
{
  "name": "zookeeper",
  "data": {
    "type": "infrastructure",
    "host": "hostname",
    "port": "2181",
    "chroot": "/pz.services"
  }
}
EOF

read -d '' regkafka << EOF
{
  "name": "kafka",
  "data": {
    "type": "infrastructure",
    "brokers": "localhost:9092"
  }
}
EOF

read -d '' reges << EOF
{
  "name": "elastic-search",
  "data": {
    "type": "infrastructure",
    "host": "localhost:9200"
  }
}
EOF

read -d '' regdisc << EOF
{
  "name": "pz-discover",
  "data": {
    "type": "core-service",
    "address": "localhost:3000"
  }
}
EOF

curl -H "Content-Type: application/json" -X PUT -d "$regzoo" http://localhost:3000/api/v1/resources
curl -H "Content-Type: application/json" -X PUT -d "$regkafka" http://localhost:3000/api/v1/resources
curl -H "Content-Type: application/json" -X PUT -d "$reges" http://localhost:3000/api/v1/resources
curl -H "Content-Type: application/json" -X PUT -d "$regdisc" http://localhost:3000/api/v1/resources

echo ========================

curl -X GET http://localhost:3000/api/v1/resources
echo
