#!/bin/sh
rm -rfv pkg/linux_386/ngrok/ src/ngrok/client/assets/ src/ngrok/server/assets/ 
rm -fv bin/ngrok*
GOOS=linux GOARCH=386 make release-server release-client

