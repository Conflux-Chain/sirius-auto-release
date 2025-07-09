# Auto release conflux scan

This project is to automatically build conflux core space and eSpace scan, and you can download the static files of the build from the release page

# How to use

## Use Pre-build assets

You can download the pre-build assets from [release](https://github.com/Conflux-Chain/sirius-auto-release/releases) page, `scan.zip` is Conflux core space and `scan-eth.zip` is Conflux eSpace

## Sirius CLI

Sirius CLI is a command-line tool designed to simplify the setup of the Conflux Scan frontend. It helps you download necessary components and generate the required configuration files.

You can download from [release v0.0.3](https://github.com/Conflux-Chain/sirius-auto-release/releases/tag/v0.0.3).

### Usage

Sirius CLI offers two main modes of operation:

### 1. Interactive Mode (Generate a Configuration File)

This mode is ideal if you are using Sirius for the first time or wish to create a new, customized configuration file. The `init` command will guide you through a series of prompts to gather the necessary information.

```bash
sirius init
```

To specify a custom name and path for the generated configuration file:

```bash
sirius init -o path/to/your/myConfig.toml
```

this will create myConfig.toml at the specified location with the settings gathered during the interactive session.

### 2. Using a Pre-built or Existing Configuration File

This mode is designed for situations where you already have a configuration file, or you want to use one of the pre-built configurations provided with Sirius.

When you download and unzip Sirius, you will find a collection of pre-built .toml configuration files (e.g., prebuilt-testnet.toml, prebuilt-mainnet.toml). These files are typically located in a configs/ subfolder within the unzipped archive, or directly alongside the sirius executable. They are tailored for common scenarios, such as connecting to the Testnet or Mainnet.

```bash
sirius --config ./configs/prebuilt-testnet.toml
```
