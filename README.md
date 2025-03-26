[![CI Status](https://github.com/sunba23/notifly/actions/workflows/release.yml/badge.svg)](https://github.com/sunba23/notifly/actions/workflows/release.yml)
[![Last Commit](https://img.shields.io/github/last-commit/sunba23/notifly)](https://github.com/sunba23/notifly/commits/master)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sunba23/notifly/blob/master/LICENSE)

# Notifly

Notifly is a simple CLI app for flight price monitoring. It composes of a few services ran in goroutines and communicating through channels.

## Usage

Get the latest version from [releases](https://github.com/sunba23/notifly/releases/) or use the [docker image](https://github.com/sunba23/notifly/pkgs/container/notifly).

### Example

```sh
./notifly monitor --date-from 2025-04-10 --date-to 2025-05-31 --from WRO --to STN --noti-price 500
```

## Tech Stack

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

