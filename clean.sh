#!/bin/sh
rm -fv assets/client/tls/ngrokroot.crt device.crt assets/server/tls/snakeoil.crt assets/server/tls/snakeoil.key
rm -fv device.* rootCA.*
rm -rfv pkg/linux_386/ src/ngrok/client/assets/ src/ngrok/server/assets/ 
rm -fv bin/ngrok*
rm -rfv pkg/windows_386/ bin/windows_386/

