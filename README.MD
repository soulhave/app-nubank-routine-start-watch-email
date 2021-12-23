# App Nubank Routine Start Watch Email

App responsible to enable the e-mail to be watched and notify a pubsub queue.

Deploy Cloud Functions
```shell
gcloud functions deploy StartWatchEmailHTTP --runtime go116 --trigger-http --allow-unauthenticated
```