# Yet Another Streaming Tool 

[![forthebadge](http://forthebadge.com/images/badges/made-with-go.svg)](http://forthebadge.com)
[![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com)

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

*YAST* is a TUI utility that will let you stream your favorite movies/tv-series in one command. This utility is built in Go using [Cobra](https://github.com/spf13/cobra) for Seamless CLI experience, [Bubbletea](https://github.com/charmbracelet/bubbletea) for beautiful TUI, [Go-Colly](https://github.com/gocolly/colly) for Web-Scraping and [WebTorrent](https://github.com/webtorrent/webtorrent) API for Streaming. 

## ⚡️ Quick start

1. [Download](https://golang.org/dl/) and install **Go**. Version `1.18` or higher is required.

2. [Download](https://github.com/webtorrent/webtorrent-cli) and install **WebTorrent-CLI**. 
```bash
npm install webtorrent-cli -g
```
3. [Download](https://www.videolan.org/) and install VLC Media Player. (Support for more players will be added in the next release)

4. Clone the repo and go inside the repo folder.
```bash
git clone https://github.com/qascade/yast && cd yast
```
5. Build the binary
```bash
go build ./...
```
6. Run the binary
```bash
./yast --help
YAST is a TUI utility that will let you stream your favorite movies/tv-series in one command.

Usage:
  yast [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Used to change already set user preferences or reset the user preferences to default
  help        Help about any command
  search      A brief description of your command
  setup       setup yast for first-time users

Flags:
  -h, --help   help for yast

Use "yast [command] --help" for more information about a command.
```
7. Run the following to choose default player.
```bash
./yast setup 
```
8. Search for a movie. 
```bash
./yast search -m spiderman
```
9. Select desired movie from the list.
<img height="550" width="800" alt="Screenshot 2022-08-25 at 5 40 13 AM" src="https://user-images.githubusercontent.com/54154054/186545637-6cc844c4-0102-4491-8dac-a6648002a219.png">

10. The movie starts with the default Player. Wohoo!!

## NOTE
1. The current supported target requires VPN to get results. Please make sure that you are connected to a VPN Server before searching or you may not get the results. We personally recommend using CloudFlare Warp. You can [download](https://1.1.1.1/) it here.
2. The Project is still in development stage. So, you might see undesired behaviour while using yast. Please file an issue if you experience any bugs or undesired behaviour. 


## Contribution
1. You can look at the design [docs](https://github.com/qascade/yast/blob/main/docs) on how we are planning to build this project.
2. Contributions and suggestions are always welcome. 
3. Look at [contibutions.md](https://github.com/qascade/yast/blob/main/contributions.md) for more details.

