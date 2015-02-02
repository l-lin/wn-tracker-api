#!/bin/bash

export PORT=3000
export GOOGLE_CLIENT_ID=""
export GOOGLE_CLIENT_SECRET=""
export GOOGLE_REDIRECT_URL="http://localhost:3000/oauth2callback"

if [ -z ${GOOGLE_CLIENT_ID} ]; then
    echo "[x] You need to provide the variable GOOGLE_CLIENT_ID from Google developer console: https://console.developers.google.com/project"
    exit -1
fi

if [ -z ${GOOGLE_CLIENT_SECRET} ]; then
    echo "[x] You need to provide the variable GOOGLE_CLIENT_SECRET from Google developer console: https://console.developers.google.com/project"
    exit -1
fi

gin run

