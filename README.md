# Stripey Horse 🦓
### A knock-off Zebra renderer

A simple wrapper around a `Go` ZPL renderer ([ingridhq/zebrash](https://github.com/ingridhq/zebrash)) to avoid constantly getting rate limited by [Labelary](https://labelary.com/service.html).

## Getting Started
This CLI uses a JSON object for config and then the binary ZPL should be piped in.
This program was designed for process communication (ie PHP to Go) so it deals with binary blobs instead of filenames.
```sh
CONFIG='{"labelWidthMm": 101.6, "labelHeightMm": 152.4, "dpmm": 8, "rotation": 0}'
stripey-horse --config "$CONFIG" --output "./test_data/test_output.png" < "$zpl_file"
```
Or using the [PHP client](https://github.com/benfaerber/zpl-to-png):
```php
$client = new StripeyHorseClient("/usr/bin/stripey_horse");
$config = StripeyHorseConfig::builder()
    ->labelPreset("6x4")
    ->rotation(90)
    ->build();

$imageData = $client->convertZplToRawImage($zplContent, $config);

file_put_contents("my_converted_image.png", $imageData);
```

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

## Development

- **Compile** - `./scripts/build.sh`
- **Test Binary** - `./scripts/run_all.sh`
- **Benchmark** - `./scripts/benchmark.php 10`

## This uses process communication, why no FFI?
This was a tough design but I decided process communication was better
for my use case (a PHP app needing to render ZPL files to PNGs and serve them to customers).

### Reasons
- FFI would require going Go -> C -> PHP
- The C wrapper code would require manual memory management
- The `.so` would have to be installed in the PHP env instead of just installing a simple binary
- FFI would connect PHP to this program. A crash would be fatal for PHP.  
- FFI would have saved around 50ms invocation time, but for an app that takes 300ms to even render, not worth the trouble!
