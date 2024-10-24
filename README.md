# Auto release conflux scan

This project is to automatically build conflux cSpace and eSpace scan, and you can download the static files of the build from the release page

# How to use

## Use Pre-build assets

You can use the pre-build assets, and you can download the static files of the build from the [release](https://github.com/Conflux-Chain/sirius-auto-release/releases) page

## Configure

After downloading the compressed package and extracting it, you will find a `config.json file`. Simply modify the parameters within it according to your needs.

## Nginx example

If you are using nginx as your server, you can configure it as follows:

```
server {
    listen    80;
    root      /your/static/files/build;

    location / {
            try_files $uri $uri/ /index.html;
    }
    location /v1/ {
        proxy_pass <your conflux san backend>/v1/;
    }
    location /stat/ {
        proxy_pass <your conflux san backend>/stat/;
    }
    location /rpcv2/ {
        proxy_pass <your conflux san backend>/rpcv2/;
    }
}

```
