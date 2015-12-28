#!/bin/sh
rm -fv assets/client/tls/ngrokroot.crt device.crt assets/server/tls/snakeoil.crt assets/server/tls/snakeoil.key
export NGROK_DOMAIN="tunnel.ptwop.xyz"
echo $NGROK_DOMAIN
openssl genrsa -out rootCA.key 2048
openssl req -x509 -new -nodes -key rootCA.key -subj "/CN=$NGROK_DOMAIN" -days 5000 -out rootCA.pem
openssl genrsa -out device.key 2048
openssl req -new -key device.key -subj "/CN=$NGROK_DOMAIN" -out device.csr
openssl x509 -req -in device.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out device.crt -days 5000
cp -fv rootCA.pem assets/client/tls/ngrokroot.crt
cp -fv device.crt assets/server/tls/snakeoil.crt
cp -fv device.key assets/server/tls/snakeoil.key
rm -fv device.* rootCA.*
