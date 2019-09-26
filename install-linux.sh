#!/bin/sh
rm -rfv /usr/local/ngrok/
mkdir -p /usr/local/ngrok/bin
cp -r assets/ /usr/local/ngrok
cp -r bin/ /usr/local/ngrok/
