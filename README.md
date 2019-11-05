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

## Install the service to Kubernetes with Helm

```bash
helm upgrade \
  --install pubsub2http deployments/chart/ \
  --set gcp.auth_file=<your-service-account-json-file-name> \
  --set env.gcp_topic_name=<your-topic-name> \
  --set env.gcp_subscription_name=<your-subscription-name> \
  --set env.gcp_project_id=<your-project-id> \
  --force
```

## Helm Chart Values / default values

Value Name                   |Default Value                           |Description
-----------------------------|----------------------------------------|------------------------------------------------------------
`env.post_url`               |`http://nginx.default.svc.cluster.local`|HTTP Target to post the GCP Message
`env.gcp_credentials_file`   |`/tmp/gcp/auth.json` *don't change!*    |The location of the Auth file in the pod
`env.gcp_topic_name`         |none                                    |Your PubSub Topic Name
`env.gcp_create_topic`       |`"TRUE"`                                |If set to "TRUE", Topic will be created if not exists
`env.gcp_subscription_name`  |none                                    |Your Subscription Name
`env.gcp_create_subscription`|`"TRUE"`                                |If set to "TRUE", Subscription will be created if not exists
`env.gcp_project_id`         |none                                    |Your GCP Project ID
`gcp.auth_file`              |none                                    |The local name of your Auth file when installing
