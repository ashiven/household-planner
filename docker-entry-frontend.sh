#! /usr/bin/env sh

set -eu

envsubst '${DOMAIN_NAME}' </etc/nginx/conf.d/nginx.conf.template >/etc/nginx/nginx.conf

exec "$@"
