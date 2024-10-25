# Auto release conflux scan

This project is to automatically build conflux cSpace and eSpace scan, and you can download the static files of the build from the release page

# How to use

## Use Pre-build assets

You can download the pre-build assets from [release](https://github.com/Conflux-Chain/scan-auto-release/releases) page, `scan.zip` is Conflux core space and `scan-eth.zip` is Conflux eSpace

## Configure

After downloading the compressed package and extracting it, you will find a `config.json` file. Simply modify the parameters within it according to your needs.

```bash
# unzip scan.zip  or unzip scan-eth.zip
vim ./build/config.json
```

1. scan (conflux core space)

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

2. scan-eth (conflux eSpace)

```js
{
    // conflux eSpace open api url
    "ENV_API_HOST": "https://evmapi.confluxscan.net",
    // conflux core space open api url
    "ENV_CORE_API_HOST": "https://api.confluxscan.net",
    // conflux scan backend url
    "ENV_CORE_SCAN_HOST": "https://www.confluxscan.net",
    // eSpace rpc url
    "ENV_RPC_SERVER": "https://evm-cfxbridge.confluxrpc.com",

    // network config
    "ENV_NETWORK_ID": 1030,
    // keep the default value
    "ENV_NETWORK_TYPE": "EVM",
    // MAINNET or TESTNET or DEVNET
    "ENV_CHAIN_TYPE": "MAINNET",

    // this config is control add network to wallet
    "ENV_WALLET_CONFIG": {
        "chainId": 1030,
        "chainName": "Conflux eSpace",
        "rpcUrls": [
            "https://evm.confluxrpc.com"
        ],
        "blockExplorerUrls": [
            "https://evm.confluxscan.io/"
        ],
        "nativeCurrency": {
            "name": "Conflux",
            "symbol": "CFX",
            "decimals": 18
        }
    },
    // logo url
    "ENV_LOGO": "https://evm.confluxscan.net/static/media/logo.8e57dceb.svg",

    // theme
    "ENV_THEME": {
        "primary": "#17B38A",
        "antdPrimaryButtonBg": "#7789D3",
        "buttonBg": "rgba(0, 84, 254, 0.8)",
        "outlineColor": "#7789D3",
        "shadowColor": "rgba(30, 61, 228, 0.2)",
        "searchButtonBg": "#AFE9D2",
        "searchButtonHoverBg": "#17B38A",
        "gasPriceLineBg": "#F0F4F3",
        "footerBg": "#05343F",
        "footerHighLightColor": "#AFE9D2",
        "linkColor": "#1e3de4",
        "linkHoverColor": "#0f23bd",
        "chartColors": [
            "#7cb5ec",
            "#434348",
            "#f7a35c",
            "#2b908f",
            "#91e8e1",
            "#90ed7d",
            "#8085e9",
            "#f15c80",
            "#e4d354",
            "#f45b5b"
        ],
        "mixedChartColors": [
            "#7cb5ec",
            "#90ed7d",
            "#434348"
        ],
        "pieChartColors": [
            "#7cb5ec",
            "#434348",
            "#f7a35c",
            "#2b908f",
            "#91e8e1",
            "#90ed7d",
            "#8085e9",
            "#f15c80",
            "#e4d354",
            "#f45b5b"
        ],
        "chartTitleColor": "#7789D3",
        "chartDetailLinkColor": "#1e3de4"
    },
    "ENV_ICONS": {
        "imgArrow": "/static/media/arrow.4e79d439.svg"
    },
    "ENV_LOCALES_EN": {},
    "ENV_LOCALES_CN": {}
}

```

Note: config.json is json format it not support comment, you need remove the comment.

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
        proxy_pass <your conflux scan backend URL>/v1/;
    }
    location /stat/ {
        proxy_pass <your conflux scan backend URL>/stat/;
    }
    location /rpcv2/ {
        proxy_pass <your conflux scan backend URL>/rpcv2/;
    }
}

```

Note: Conflux core space and Conflux eSpace have different backend interfaces, so when configuring, it is necessary to use front-end interfaces and backend interfaces that match each other.
