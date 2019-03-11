# Terragrunt Switcher 

The `s3upload` command line tool lets you upload files to s3. 
You have the option to provide a configuration file with s3 information and directory or provide the information as parameters on the command line.

<hr>

## Installation

`s3upload` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/){:target="_blank"}. 


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
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload.gif" alt="drawing" style="width: 480px;"/>

1.  Provide s3 configuration and directory to upload information.
2.  For example, `s3upload -b bucketname.s3.com -d directorytoupload`.
3.  Use the `-i` parameter to include the base directory.
4.  Pass `-h` for more help. 

The most recently selected versions are presented at the top of the dropdown.

### Supply version on command line
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload-v4.gif" alt="drawing" style="width: 480px;"/>

1. You can also supply a configuration file with 
2. For example, `s3upload -c locationOfConfigFile`. If you do not provide a config file, s3upload will look at the current directory.

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/s3upload/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)