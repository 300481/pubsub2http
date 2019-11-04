# pubsub2http

A bridge from Google PubSub to HTTP

## Environment Variables

Environment Variable     |Description                                 |Valid values
-------------------------|--------------------------------------------|-----------------------
`POST_URL`               |The URL to post the message                 |*String*
`GCP_CREDENTIALS_FILE`   |The location of the credentials file (JSON) |*String*
`GCP_TOPIC_NAME`         |The PubSub topic name                       |*String*
`GCP_CREATE_TOPIC`       |Allow creation of topic if not exists       |*String* `TRUE`/`FALSE`
`GCP_SUBSCRIPTION_NAME`  |The PubSub subscription name                |*String*
`GCP_CREATE_SUBSCRIPTION`|Allow creation of subscription if not exists|*String* `TRUE`/`FALSE`
`GCP_PROJECT_ID`         |The Google Project ID                       |*String*

## Prerequisites

* account for [Google Cloud Platform](https://cloud.google.com/)
* Docker

## Start the service

```bash
export POST_URL=<yourposturl>
export GCP_CREDENTIALS_FILE=<filelocation>
export GCP_TOPIC_NAME=<yourtopicname>
export GCP_CREATE_TOPIC=TRUE
export GCP_SUBSCRIPTION_NAME=<yoursubscriptionname>
export GCP_CREATE_SUBSCRIPTION=TRUE
export GCP_PROJECT_ID=yourprojectid
export VERSION=yourpreferredversion
docker run -d --rm --name=pubsub2http \
    -v credentialsfile-path:containerpath \
    -e POST_URL=${POST_URL} \
    -e GCP_CREDENTIALS_FILE=${GCP_CREDENTIALS_FILE} \
    -e GCP_TOPIC_NAME=${GCP_TOPIC_NAME} \
    -e GCP_CREATE_TOPIC=${GCP_CREATE_TOPIC} \
    -e GCP_SUBSCRIPTION_NAME=${GCP_SUBSCRIPTION_NAME} \
    -e GCP_CREATE_SUBSCRIPTION=${GCP_CREATE_SUBSCRIPTION} \
    -e GCP_PROJECT_ID=${GCP_PROJECT_ID} \
    300481/pubsub2http:${VERSION}
```
