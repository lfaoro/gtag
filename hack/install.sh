#!/usr/bin/env bash
#
# gTag Installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/lfaoro/gtag/master/hack/install.sh | bash

APP="gtag"
VERSION="0.6.0"

set -e

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  curl -fsSL \
    "https://github.com/lfaoro/${APP}/releases/download/v${VERSION}/${APP}_${VERSION}_linux_amd64.tar.gz" |
    tar -xzv ${APP} >>/dev/null
  mv ${APP} /usr/local/bin/${APP}

elif [[ "$OSTYPE" == "darwin"* ]]; then
  if [[ "$(command -v brew)" != "" ]]; then
    brew install lfaoro/tap/${APP}
  else
    curl -fsSL \
      "https://github.com/lfaoro/${APP}/releases/download/v${VERSION}/${APP}_${VERSION}_linux_amd64.tar.gz" |
      tar -xzv ${APP}
    sudo mv ${APP} /usr/local/bin/${APP}
  fi

else
  echo "The ${APP} installer does not work for your platform: $OS"
  echo "Please file an issue at https://github.com/lfaoro/${APP}/issues/new"
  exit 1
fi

# TODO(leo): Add verification that it installed successfully.

echo "${APP} installed! Run \`gtag -h\` to start."
