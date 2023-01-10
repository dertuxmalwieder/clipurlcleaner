[![Scc Count Badge](https://sloc.xyz/github/dertuxmalwieder/clipurlcleaner?category=code)](https://github.com/dertuxmalwieder/clipurlcleaner) [![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/GebtmireuerGeld)

![Screenshot](https://i.imgur.com/7VSRqHb.png)

Watches your clipboard for shortened URLs and unshortens them for a less stupid paste. While doing this, it also removes known tracking parameters from the copied URLs.

## Example

Start the tool and copy this link:

    https://bit.ly/3QqdCUd

When you paste it anywhere, it will automatically be converted to

    https://code.rosaelefanten.org/clipurlcleaner

# Build

`go build`

## On Windows

`go build -ldflags="-H windowsgui"`

### Prebuilt binaries (Windows-only)

You can find the latest [prebuilt binaries](https://cdn.tuxproject.de/projects/clipurlcleaner/) on *cdn.tuxproject.de*. Those are compressed with [WinRAR](https://www.rarlab.com). The current version is [clipurlcleaner-win-20221021.rar](https://cdn.tuxproject.de/projects/clipurlcleaner/clipurlcleaner-win-20221021.rar) (MD5: `0C2D638CE38DD49E9941F3A612B1DD11`).

# Run

`./clipurlcleaner >/dev/null &`
