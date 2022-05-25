#! usr/bin/bash
openapi-generator generate -i https://api.twitter.com/labs/2/openapi.json -g go -o ./ --additional-properties=packageName=twitterapi --git-user-id cute-gobbo \
  --git-repo-id twitter-go-api
go mod tidy