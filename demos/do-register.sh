#!/bin/sh

read -d '' regzoo << EOF
{
  "name": "zookeeper",
  "data": {
    "type": "infrastructure",
    "address": "http://localhost:2181",
    "infrastructure": {},
    "local": {}
  }
}
EOF

read -d '' regkafka << EOF
{
  "name": "kafka",
  "data": {
    "type": "infrastructure",
    "address": "http://localhost:9092",
    "infrastructure": {},
    "local": {}
  }
}
EOF

read -d '' reges << EOF
{
  "name": "elastic-search",
  "data": {
    "type": "infrastructure",
    "address": "http://localhost:9200",
    "infrastructure": {},
    "local": {}
  }
}
EOF

read -d '' regdisc << EOF
{
  "name": "pz-discover",
  "data": {
    "type": "core-service",
    "address": "http://localhost:3000",
    "core-service": {},
    "local": {}
  }
}
EOF

curl -H "Content-Type: application/json" -X PUT -d "$regzoo" http://localhost:3000/api/v1/resources
curl -H "Content-Type: application/json" -X PUT -d "$regkafka" http://localhost:3000/api/v1/resources
#curl -H "Content-Type: application/json" -X PUT -d "$reges" http://localhost:3000/api/v1/resources
curl -H "Content-Type: application/json" -X PUT -d "$regdisc" http://localhost:3000/api/v1/resources

echo ========================

curl -X GET http://localhost:3000/api/v1/resources/kafka
echo
