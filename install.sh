#!/usr/bin/env sh

ARCH=$(uname -m)
if [ ${ARCH} == 'x86_64' ]; then
  ARCH=amd64
elif [ ${ARCH} == 'armv6l' ]; then
  ARCH=arm6
elif [ ${ARCH} == 'armv7l' ]; then
  ARCH=arm7
elif [ ${ARCH} == 'arm64' ]; then
  ARCH=arm64
else
  echo "do not support ${ARCH}"
  echo "please feel free to help improve this at https://github.com/LinuxSuRen/tools/blob/master/install.sh"
  exit 1
fi

OS=$(uname -s)
if [ ${OS} == 'Darwin' ]; then
  OS=darwin
elif [ ${OS} == 'Linux' ]; then
  OS=linux
else
  echo "do not support ${OS}"
  echo "please feel free to help improve this at https://github.com/LinuxSuRen/tools/blob/master/install.sh"
  exit 2
fi

which curl
if [ $? = 0 ]; then
  curl -L https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-${OS}-${ARCH}.tar.gz | tar xzv hd
else
  wget https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-${OS}-${ARCH}.tar.gz -O - | tar xzv hd
fi

sudo mv hd /usr/bin/hd
hd fetch
hd completion
