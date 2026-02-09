# sf
`sf` is a universal convertor between biodiversity exchange formats.

The `sf` app converts most popular biodiversity formats to each other and
is able to compare datasets. It is very useful for maintaining checklists, or
for adjusting a local checklist for a global consumption.

The program uses Species File Group Archive (SFGA) format as an intermediary.
SFGA is based on SQLite database, easily exchangeable (just one file) stable
(SQLite promises backward compatibility until 2050), makes it easy to query and
modify (via SQL language) and even can be used as a standard backend for an
ecosystem of applications.

## Installation

Compiled programs in Go are self-sufficient and small (`sf` is only a
few megabytes). As a result the binary file of `sf` is all you need to
make it work. You can install it by downloading the [latest version of the
binary][releases] for your operating system and CPU architecture, and
placing it in your `PATH`.

### Install with Homebrew (Mac OS X, Linux)

[Homebrew] is a packaging system originally made for Mac OS X. You can use it
now for Mac, Linux, or Windows X WSL (Windows subsystem for Linux).

1. Install Homebrew according to their [instructions][Homebrew].

2. Install `sf` with:

   ```bash
   brew update
   brew tap gnames/gn
   brew install sf
   ```

### Linux or Mac OS X

Move `sf` executable somewhere in your PATH
(for example `/usr/local/bin`)

```bash
tar xvf sf-xxx.tar.gz
sudo mv sf /usr/local/bin
```

If you're using Mac OS, you might encounter a security warning that prevents
`sf` from running. Here's how to fix it:

1. In the warning dialog click the `Done` button (not the `Move to Trash`
   button).

1. Locate the Security Settings: Go to `System Settings -> Privacy & Security`
   and scroll down to the `Security` section.

1. Allow `sf`: You should see a message saying
   `"sf" was blocked...`. Click the `Allow Anyway` button next to it.

1. Run `sf` again: Try running `sf` from your terminal. This time,
   a dialog box will pop up with an `Open Anyway` button.

1. Open and Unblock: Click `Open Anyway` and enter your administrator
   password when prompted. This will unblock the `sf` binary.

After these steps, you should be able to use `sf` without any issues.
You can also copy, move, or rename it freely.

### Windows

One possible way would be to create a default folder for executables and place
`sf` there.

Use `Windows+R` keys
combination and type "`cmd`". In the appeared terminal window type:

```cmd
mkdir C:\bin
copy path_to\sf.exe C:\bin
```

[Add `C:\bin` directory to your `PATH`][winpath] `user` and/or `system`
environment variables.

It is also possible to install [Windows Subsystem for Linux][wsl] on Windows
(v10 or v11), and use `sf` as a Linux executable.

### Install with Go

If you have Go installed on your computer use

```bash
go install github.com/sfborg/sf@latest
```

For development install [just] and use the following:

```bash
git clone https://github.com/sfborg/sf.git
cd sf
just tools
just install
```

You do need your `PATH` to include `$HOME/go/bin`

## Usage

### Conversions

The `sf` app supports conversions from DarwinCore Archive (DwCA),
Catalogue of Life Data Package (CoLDP), comma-separated, tab-separated,
pipe-separated files, or texts where scientific name strings are placed one
name per line. All these formats are converted to SQLite-based Species
File Group Archive (SFGA). The resulting SFGA file can then be converted to
DwCA, CoLDP, CSV, or plain text.

```bash
sf from coldp /path_to/coldp/dataset.zip /path_to/output/dataset
sf to dwca /path_to/sfga/dataset.sqlite.zip /path_to/output/dwca/dataset
```

Note that file extensions will be added automatically, so output like
`dataset` would create SFGA outputs as `dataset.sql` and `dataset.sqlite`,
where the `dataset.sql` file is a text dump of sql statements required
for generation of the dataset, and `dataset.sqlite` is a binary file that
can be used directly as a SQLite database.

It is also possible to use remote input files by providing their URLs.

### Datasets Comparison

For example we want to compare a DwCA dataset with a checklist of names in a
CSV file. First both files need to be converted to SFGA archives which we can
then compare. The resulting difference between the files is going to be saved
as a SQLite database which can be queried with SQL statements.

```bash
sf from dwca /path_to/dadaset1.dwca.zip /path_to/sfga/dataset1
sf from xsv /path_to/dataset2.csv /path_to/sfga/dataset2
sf diff /path_to/sfga/dataset1.sqlite /path_to/sfga/dataset2.sqlite diff.sqlite
```

### Updating older SFGA files to the latest schema.

To be able to use `sf` the SFGA files should correspond to the latest version
of [SFGA schema], so `sf` has a convenience command that seamlessly converts
older SFGA versions to the latest ones migrating all the data.

```bash
sf update /path_to/outdated/dataset.sqlite /path_to/updated/dataset
```

[winpath]: https://www.computerhope.com/issues/ch000549.htm
[wsl]: https://docs.microsoft.com/en-us/windows/wsl/
[Homebrew]: https://brew.sh/
[just]: https://github.com/casey/just
[LICENSE]: ./LICENSE
