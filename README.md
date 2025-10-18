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

## Getting Started 

You'll need the `go` compiler. This is for linux so use docker if you don't use linux.
You could also figure out how to compile for your platform in `build.sh`


- **Compile** - `./scripts/build.sh`
- **Test Binary** - `./scripts/run_all.sh`   
* **Benchmark** - `./scripts/benchmark.php 10`
