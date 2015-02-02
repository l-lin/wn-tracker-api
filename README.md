Web novels tracker API
======================

Getting started
---------------

Fetch the projet and install the dependencies:

```go
go get github.com/l-lin/wn-tracker-api
go get github.com/tools/godep
cd $GOHOME/github/l-lin/wn-tracker-api
godep go install
```

You need to set the following environment variables in the `startup.sh`:

|Variable name          |Description                            |Example                                |
|-----------------------|---------------------------------------|---------------------------------------|
|PORT                   |Server port                            |4747                                   |
|GOOGLE_CLIENT_ID       |The client ID of your google API       |xxxx.apps.googleusercontent.com        |
|GOOGLE_CLIENT_SECRET   |The client secret of your Google API   |ABCDEFGHIJKLMOPQRSTUVWXYZ              |
|GOOGLE_REDIRECT_URL    |The redirect URL after being connected |http://localhost:3000/oauth2callback   |

Thoses variables are available in your [Google developer console](https://console.developers.google.com/project).

After the configuration, you just need to execute `startup.sh` and access to `http://localhost:3000/`

