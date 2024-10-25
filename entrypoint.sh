#!/bin/sh

default_scan_path="/app/scan/build"
default_scan_eth_path="/app/scan-eth/build"

scan_path="/data/scan"
scan_eth_path="/data/scan-eth"

if [ -z "$(ls -A $scan_path 2>/dev/null)" ]; then

    cp -r $default_scan_path/* $scan_path
fi


if [ -z "$(ls -A $scan_eth_path 2>/dev/null)" ]; then

    cp -r $default_scan_eth_path/* $scan_eth_path
fi

exec nginx -g 'daemon off;'