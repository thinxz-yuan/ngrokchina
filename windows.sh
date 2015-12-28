#!/bin/sh
rm -rfv pkg/windows_386/ngrok/ src/ngrok/client/assets/ src/ngrok/server/assets/ bin/windows_386/
GOOS=windows GOARCH=386 make release-server release-client
