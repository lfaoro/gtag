#!/usr/bin/env bash
#
# gTag Installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/lfaoro/gtag/master/hack/install.sh | bash

VERSION="0.6.0"
BREW=$(which brew)
APP="gtag"

set -e

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  set -x
  curl -fsSL "https://github.com/lfaoro/${APP}/releases/download/v${VERSION}/${APP}_${VERSION}_linux_amd64.tar.gz" \
  | tar -xzv ${APP}
  sudo mv ${APP} /usr/local/bin/${APP}

elif [[ "$OSTYPE" == "darwin"* ]]; then
  if [[ "$BREW" != "" ]]; then
    set -x
    brew install lfaoro/tap/${APP}
  else
    set -x
    curl -fsSL \
      "https://github.com/lfaoro/${APP}/releases/download/v${VERSION}/${APP}_${VERSION}_linux_amd64.tar.gz" |
      tar -xzv ${APP}
    sudo mv ${APP} /usr/local/bin/${APP}
  fi

else
  set +x
  echo "The ${APP} installer does not work for your platform: $OS"
  echo "Please file an issue at https://github.com/lfaoro/${APP}/issues/new"
  exit 1
fi

# TODO(leo): Add verification that it installed successfully.

set +x
echo "${APP} installed! Run \`gtag -h\` to start."
Type ! to start highlighting.
