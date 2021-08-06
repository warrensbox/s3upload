[![Build Status](https://travis-ci.org/warrensbox/terraform-switcher.svg?branch=master)](https://travis-ci.org/warrensbox/terraform-switcher)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/terraform-switcher)](https://goreportcard.com/report/github.com/warrensbox/terraform-switcher)
[![CircleCI](https://circleci.com/gh/warrensbox/terraform-switcher/tree/master.svg?style=shield&circle-token=55ddceec95ff67eb38269152282f8a7d761c79a5)](https://circleci.com/gh/warrensbox/terraform-switcher)

# s3uploader - Need further development - Open for MR - No longer maintained
                                    
<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/logo.png" alt="drawing" width="120" height="130"/>

The `s3upload` command line tool lets you upload files to s3. 
You have the option to provide a configuration file with s3 information and directory or provide the information as parameters on the command line.

<hr>

## Installation

`s3upload` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. If you do not have homebrew installed, [click here](https://brew.sh/){:target="_blank"}. 


```ruby
brew install warrensbox/tap/s3upload
```

### Linux

Installation for Linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/s3upload/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from the source [here](https://github.com/warrensbox/s3upload/releases) 

<hr>

## How to use:
### Use command line to provide s3 config information
<img  src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload.gif" alt="drawing" style="width: 480px;"/>

1.  Provide s3 configuration and directory to upload information.
2.  For example, `s3upload -b bucketname.s3.com -d directorytoupload`.
3.  Use the `-i` parameter to include the base directory.
4.  Pass `-h` for more help. 
5.  Pass `-r` for aws region. 
6.  Pass `-b` for aws s3 bucket. 

The most recently selected versions are presented at the top of the dropdown.

### Use configuration file
<img   src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload-v4.gif" alt="drawing" style="width: 480px;"/>

1. You can also supply a configuration file with 
2. For example, `s3upload -c locationOfConfigFile`. If you do not provide a config file, s3upload will look at the current directory.

#### Sample config file
Config json file should be named `s3config.json` and be placed in the root directory if you want `s3upload` to read the default conifg
```json
{
    "source": "test",
    "bucket": "simple.s3.uploader",
    "exclude": "main,README.md"
}
```
<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/s3upload/issues){:target="_blank"}



See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)
