# rename

Cross compatible rename cli, similar to sed.

## Usage

```
rename [-hn] s/pattern/replace/ path
  -h    Print help
  -n    Dry run, will not actually rename any files
```

## Example

```shell
tree
.
├── 1.yaml
├── 2.yaml
├── 3.yaml
├── 4.yaml
└── 5.yaml

rename -n 's/(\d)\.yaml/$1.yml/' .                                                                                                                                                                                                16:00:09
1.yaml   -->   1.yml
2.yaml   -->   2.yml
3.yaml   -->   3.yml
4.yaml   -->   4.yml
5.yaml   -->   5.yml
```

## Installation

### macOS
`rename` is available on Homebrew.

#### Homebrew
install
```shell
brew tap b4nst/homebrew-tap
brew install b4nst/homebrew-tap/rename
```

upgrade
```shell
brew upgrade rename
```

### Linux
#### Debian/Ubuntu
Install and upgrade:

Download the .deb file from the [releases page](https://github.com/b4nst/rename/releases/latest)
```shell
sudo apt install ./rename_*_linux_amd64.deb
```
#### Fedora
Install and upgrade:

Download the .rpm file from the [releases page](https://github.com/b4nst/rename/releases/latest)
```shell
sudo dnf install ./rename_*_linux_amd64.rpm
```
#### Centos
Install and upgrade:

Download the .rpm file from the [releases page](https://github.com/b4nst/rename/releases/latest)
```shell
sudo yum localinstall ./rename_*_linux_amd64.rpm
```
#### openSUSE/SUSE
Install and upgrade:

Download the .rpm file from the [releases page](https://github.com/b4nst/rename/releases/latest)
```shell
sudo zypper in ./rename_*_linux_amd64.rpm
```
### Windows
`rename` is available via [scoop](https://scoop.sh).
```shell
scoop bucket add scoop-bucket https://github.com/b4nst/scoop-bucket.git
scoop install rename
```
### Other platforms
Prebuilt binary available from the [release page](https://github.com/b4nst/rename/releases/latest).
