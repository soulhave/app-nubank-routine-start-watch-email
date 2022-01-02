# App Nubank Routine Start Watch Email

App responsible to enable the e-mail to be watched and notify a pubsub queue.

Run Local
```shell
cd cmd && go run main.go
```

Deploy Cloud Functions
```shell
gcloud functions deploy StartWatchEmailHTTP --runtime go116 --trigger-http --env-vars-file .env.yaml --allow-unauthenticated
```

> Disclaimer: I encourage you to take a look on the google functions documention to understand how to deploy the application and the var environments.

You need to set some env vars to be used on your machine or deploy the function.

1. `APP_NUBANK_TOPIC:`
Define which for which topic the Gmail API must notify when receive a new Nubank e-mail.
2. `APP_NUBANK_MAIL_LABEL:`
Define your label on the gmail that will be flaged the nunbak e-mails.
3. `APP_NUBANK_BUCKET:`
The bucket on GCP where you store all secret information about your account.
4. `APP_NUBANK_SECRET_FILE:`
The secret file with the data about your email account.