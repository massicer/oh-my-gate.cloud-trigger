# oh-my-gate.cloud-trigger
Cloud function to trigger in order to publish open msg via Msg Broker (Pub/Sub)


### How to publish

```
gcloud config set project `PROJECT ID`
```
```
gcloud functions deploy HelloGet \
--env-vars-file=.env \
--runtime go111 --trigger-http --allow-unauthenticated
```