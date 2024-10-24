# Auto release conflux scan

This project is to automatically build conflux cSpace and eSpace scan, and you can download the static files of the build from the release page

# How to use

## Use Pre-build assets

You can use the pre-build assets you can download the static files of the build from the [release](https://github.com/Conflux-Chain/sirius-auto-release/releases) page, `sirius.zip` is Conflux core space and `sirius-eth.zip` is Conflux eSpace

## Configure

After downloading the compressed package and extracting it, you will find a `config.json file`. Simply modify the parameters within it according to your needs.

1. sirius (conflux core space)

```js
{
  // open api url required
  "ENV_OPEN_API_HOST": "https://api.confluxscan.net",
  // conflux rpc url required
  "ENV_RPC_SERVER": "https://main.confluxrpc.com",

  // network config
  // network id
  "ENV_NETWORK_ID": 1029,

  // keep the default value
  "ENV_NETWORK_TYPE": "CORE",
  // MAINNET or TESTNET or DEVNET
  "ENV_CHAIN_TYPE": "MAINNET",

  // conflux fc feature
  // fc address
  "ENV_FC_ADDRESS": "cfx:achc8nxj7r451c223m18w2dwjnmhkd6rxawrvkvsy2",
  // fc exchange address
  "ENV_FC_EXCHANGE_ADDRESS": "cfx:acdrd6ahf4fmdj6rgw4n9k4wdxrzfe6ex6jc7pw50m",
  // fc exchange interest address
  "ENV_FC_EXCHANGE_INTEREST_ADDRESS": "cfx:acag8dru4527jb1hkmx187w0c7ymtrzkt2schxg140",

  // conflux core space ENS feature
  // ENS registry address
  "ENV_ENS_REGISTRY_ADDRESS": "cfx:acemru7fu1u8brtyn3hrtae17kbcd4pd9uwbspvnnm",
  // ENS public resolver address
  "ENV_ENS_PUBLIC_RESOLVER_ADDRESS": "cfx:acasaruvgf44ss67pxzfs1exvj7k2vyt863f72n6up",
  // logo url
  "ENV_LOGO": "https://confluxscan.io/static/media/logo.1c92ddc8.svg"
}

```

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

Note: Conflux core space and Conflux eSpace have different backend interfaces, so when configuring, it is necessary to use front-end interfaces and backend interfaces that match each other.
