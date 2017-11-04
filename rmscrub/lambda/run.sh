#!/bin/sh

make

aws lambda create-function \
  --function-name HandlerA \
  --runtime python2.7 \
  --role arn:aws:iam::943964223958:role/mpg-serverless-stack-TestFunctionRole-I0YNTX7VFMN3 \
  --handler handler.HandleA \
  --zip-file fileb://handler.zip

aws lambda create-function \
  --function-name HandlerB \
  --runtime python2.7 \
  --role arn:aws:iam::943964223958:role/mpg-serverless-stack-TestFunctionRole-I0YNTX7VFMN3 \
  --handler handler.HandleB \
  --zip-file fileb://handler.zip

aws lambda update-function-code \
  --function-name HandlerA \
  --zip-file fileb://handler.zip

aws lambda update-function-code \
  --function-name HandlerB \
  --zip-file fileb://handler.zip

aws lambda invoke \
  --function-name HandlerA \
  --invocation-type RequestResponse \
  --payload file://input.txt \
  x

aws lambda invoke \
  --function-name HandlerB \
  --invocation-type RequestResponse \
  --payload file://input.txt \
  y
