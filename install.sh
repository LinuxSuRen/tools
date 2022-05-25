#!/bin/sh

which curl
if [ $0 = 0 ]; then
curl -L https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-linux-amd64.tar.gz | tar xzv hd
else
wget https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-linux-amd64.tar.gz -O - | tar xzv hd
fi

mv hd /usr/bin/hd
hd fetch
hd completion
