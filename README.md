[![Scc Count Badge](https://sloc.xyz/github/dertuxmalwieder/clipurlcleaner?category=code)](https://github.com/dertuxmalwieder/clipurlcleaner) [![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://paypal.me/GebtmireuerGeld)

![Screenshot](https://i.imgur.com/7VSRqHb.png)

Watches your clipboard for shortened URLs and unshortens them for a less stupid paste.

# Build

`go build`

## On Windows

`go build -ldflags="-H windowsgui"`

(Or use the [prebuilt binaries](https://cdn.tuxproject.de/projects/clipurlcleaner/).)

# Run

`./clipurlcleaner >/dev/null &`
