# Squawker App backend

This is a simple Squawker backend to help understand how to make send downstream messages to squawker clients (Android, iOS, web)

This is just a simple webserver. There is more to learn



# third part libraries
There are third party libraries in go that implemented all these features. These packages might be good to use if you dont want to dive to much and wirte all code for this. I probably will do if i am in a start up situation. After all, the udacity guys used a third party library for their server :)

- [https://blog.intelligentbee.com/2018/02/27/firebase-integration-golang/](https://blog.intelligentbee.com/2018/02/27/firebase-integration-golang/)
- [https://pkg.go.dev/github.com/maddevsio/fcm](https://pkg.go.dev/github.com/maddevsio/fcm)
- [https://pkg.go.dev/firebase.google.com/go/v4/messaging](https://pkg.go.dev/firebase.google.com/go/v4/messaging) You can use the actual firebase go, but you need to figure out how to get the credentials [here](https://stackoverflow.com/questions/63836187/firebase-go-sdk-with-custom-http-client)
- [https://firebase.google.com/docs/cloud-messaging/server](https://firebase.google.com/docs/cloud-messaging/server)

