# tcolors

Commandline color picker and palette builder

## Installing
Download the [latest release](https://github.com/bcicen/tcolors/releases) for your platform:

#### Linux / OSX

```bash
curl -Lo tcolors https://github.com/bcicen/tcolors/releases/download/v0.1/tcolors-0.1-$(uname -s)-amd64
chmod +x tcolors
sudo mv tcolors /usr/local/bin/
```
#### AUR

`tcolors` is also available for Arch in the [AUR](https://aur.archlinux.org/packages/tcolors)

#### Docker

```bash
docker run --rm -ti --name=tcolors \
  quay.io/vektorlab/tcolors:latest
```

## Usage

Simply run `tcolors` to view and modify the default palette. Changes are automatically saved and will persist across sessions.

### Keybindings

Key | Description
--- | ---
`↑, k` | navigate up
`↓, j` | navigate dow
`←, h` | decrease selected value
`→, l` | increase selected value
`a, <ins>` | add a new palette color
`x, <del>` | remove the selected palette color
`q, <esc>` | exit tcolors
`?` | show help menu

### Palette files

To create a new palette or use a specific palette, use the `-f` option:

```bash
tcolors -f logo-palette.toml
```

Palette colors are stored in a human-readable TOML format and all changes are saved on exit.

### Options

Option | Description
--- | ---
-f | specify palette file to load/save changes to
-p | output current palette contents
-o | color format to output (hex, rgb, hsv, all) (default "all")
-v | print version info
