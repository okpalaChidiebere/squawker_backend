# Squawker App backend

This is a simple Squawker backend to help understand how to make send downstream messages to squawker clients (Android, iOS, web)

This is just a simple webserver. There is more to learn

We used this backend to test this React native [app](https://github.com/okpalaChidiebere/reactnd-Squawker)

If you want to lean how to use firebase for react web see [Notifications API](https://developer.mozilla.org/en-US/docs/Web/API/Notifications_API/Using_the_Notifications_API), [JS client](https://firebase.google.com/docs/cloud-messaging/js/client) and [here](https://firebase.google.com/docs/cloud-messaging/js/receive?_gl=1*156yvsy*_up*MQ..*_ga*NjkxNjcyNTA4LjE3MTI0Njg0NDE.*_ga_CW55HF8NVT*MTcxMjQ2ODQ0MS4xLjAuMTcxMjQ2ODQ0MS4wLjAuMA). You MUST have Service workers enabled

I am getting the Private key for a service account gotten from fcm console and i uploaded it to S3 bucket and read the file in our server. Another option is to convert the credential file into base64 like `base64 service-account-private-key.json`. Then save into a AWS Secret manager. Then read from the secret manager decode the base64 string the pass that string as bytes to firebase sdk. See [here](https://github.com/expo/eas-cli/issues/228#issuecomment-861407074)

# third party libraries

There are third party libraries in go that implemented all these features. These packages might be good to use if you dont want to dive to much and write all code for this. I probably will do if i am in a start up situation. After all, the udacity guys used a third party library for their server :)

- [Shape notification icon](https://romannurik.github.io/AndroidAssetStudio/icons-notification.html#source.type=image&source.space.trim=1&source.space.pad=0&name=ic_duck)

- [https://firebase.google.com/docs/firestore/quickstart#go_1](https://firebase.google.com/docs/firestore/quickstart#go_1)
- [https://firebase.google.com/docs/cloud-messaging/server](https://firebase.google.com/docs/cloud-messaging/server)
