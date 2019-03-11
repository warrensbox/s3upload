# Terraform Switcher 

The `s3upload` command line tool lets you switch between different versions of [terraform](https://www.terraform.io/){:target="_blank"}. 
If you do not have a particular version of terraform installed, `s3upload` will download the version you desire.
The installation is minimal and easy. 
Once installed, simply select the version you require from the dropdown and start using terraform. 

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
curl -L https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from the source [here](https://github.com/warrensbox/terraform-switcher/releases) 

<hr>

## How to use:
### Use dropdown menu to select version
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload.gif" alt="drawing" style="width: 480px;"/>

1.  You can switch between different versions of terraform by typing the command `s3upload` on your terminal. 
2.  Select the version of terraform you require by using the up and down arrow.
3.  Hit **Enter** to select the desired version.

The most recently selected versions are presented at the top of the dropdown.

### Supply version on command line
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/s3upload/s3upload-v4.gif" alt="drawing" style="width: 480px;"/>

1. You can also supply the desired version as an argument on the command line.
2. For example, `s3upload 0.10.5` for version 0.10.5 of terraform.
3. Hit **Enter** to switch.

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/terraform-switcher/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)