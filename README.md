# X-AIR Downloader

This application is a super simple little app that speeds up downloading the X-AIR Edit application rather that having to find it on the music-tribe website.

# Usage

```
A Downloader for the Behringer X-AIR program

Usage:
  xair-downloader [flags]

Flags:
  -y, --accept-eula           automatically accept the Music Tribe End-User License Agreement
  -h, --help                  help for xair-downloader
      --version               version for xair-downloader
  -v, --xair-version string   the version of X-AIR to download (default "1.5")
```

* The -y|--accept-eula flag will allow you to skip the input necessary to accept the EULA

  * Example: `xair-downloader_darwin_amd64 -y true`

* The -h|--help flag will print this usage message

* The --version flag will print the current app version

* The -v|--xair-version sets the version of X-AIR Edit to download.

  * Example - This will download version 1.13 of X-AIR Edit: `xair-downloader_darwin_amd64 -v 1.13`

### Disclaimer

I am not affiliated with Music Tribe or Behringer, I just got tired of hunting for the app every time I nuked my raspberry pi.