# Stripey Horse 🦓
### A knock-off Zebra renderer

A simple wrapper around a `Go` ZPL renderer to avoid constantly getting rate limited by [Labelary](https://labelary.com/service.html).

## Why?

The Labelary API is great, but it only allows 5 requests per second. This leads to constant errors and failures when generating labels at scale.

Sadly, there are no simple ways to pay for higher limits. You have to contact Labelary directly and negotiate a custom deal to self-host.

Luckily, there is a Go library that renders ZPL almost perfectly. This project provides:
- A CLI wrapper around the Go ZPL renderer
- A PHP wrapper around the CLI for easy integration

## Benefits

- **No rate limits** - render as many labels as you need
- **No external dependencies** - works completely offline
- **Cost effective** - hence "Stripey Horse" instead of "Zebra"
- **Near-perfect rendering** - comparable quality to Labelary
- **Faster** - about 4x faster (see `./scripts/benchmark.php`)

## Installation

### Go Install
```bash
go install github.com/benfaerber/stripey-horse@latest
```

### Manual Installation
Download the latest binary for your platform from [Releases](https://github.com/benfaerber/stripey-horse/releases) and add it to your PATH.

### From Source
```bash
git clone https://github.com/benfaerber/stripey-horse.git
cd stripey-horse
go build -o stripey_horse
```

## Usage

After installation, run:
```bash
stripey_horse
```

## Development

- **Compile** - `./scripts/build.sh`
- **Test Binary** - `./scripts/run_all.sh`
- **Benchmark** - `./scripts/benchmark.php 10`
