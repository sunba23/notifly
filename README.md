[![CI Status](https://github.com/sunba23/notifly/actions/workflows/release.yml/badge.svg)](https://github.com/sunba23/notifly/actions/workflows/build_and_push.yml)
[![Last Commit](https://img.shields.io/github/last-commit/sunba23/notifly)](https://github.com/sunba23/mpk-isochrone/commits/master)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/sunba23/sunba.dev/blob/master/LICENSE)

# Notifly

Notifly is a simple CLI app for flight price monitoring. It composes of a few services ran in goroutines and communicating through channels.

## Example usage

```sh
go run main.go monitor --date-from 2025-04-10 --date-to 2025-05-31 --from WRO --to STN --email franeksu@gmail.com --noti-price 500
```

## Tech Stack

![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

