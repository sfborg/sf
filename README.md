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

### Importing from other formats to SFGA (`sf from`)

The `sf` app supports importing from the following formats into the
SQLite-based Species File Group Archive (SFGA):

- **DwCA** — Darwin Core Archive
- **CoLDP** — Catalogue of Life Data Package
- **CSV/TSV/PSV** — Comma-, tab-, or pipe-separated value files
  (with Darwin Core or CoLDP headers)
- **Text** — Plain text with one scientific name per line

```bash
sf from coldp /path_to/coldp/dataset.zip /path_to/output/dataset
sf from dwca /path_to/dataset.dwca.zip /path_to/output/dataset
sf from xsv /path_to/dataset.csv /path_to/output/dataset
sf from text /path_to/names.txt /path_to/output/dataset
```

Note that file extensions will be added automatically, so output like
`dataset` would create SFGA outputs as `dataset.sql` and `dataset.sqlite`,
where the `dataset.sql` file is a text dump of SQL statements required
for generation of the dataset, and `dataset.sqlite` is a binary file that
can be used directly as a SQLite database.

Input can be either a local file or a URL to a remote file. If the output
path is not provided, output files will be generated in the current directory.

### Exporting from SFGA to other formats (`sf to`)

SFGA files can be exported to the following formats:

- **CoLDP** — Catalogue of Life Data Package
- **DwCA** — Darwin Core Archive
- **CSV** — Comma-separated value file
- **Text** — Plain text with one scientific name per line

```bash
sf to coldp /path_to/sfga/dataset.sqlite.zip /path_to/output/dataset.zip
sf to dwca /path_to/sfga/dataset.sqlite.zip /path_to/output/dataset.zip
sf to xsv /path_to/sfga/dataset.sqlite /path_to/output/dataset.csv
sf to text /path_to/sfga/dataset.sqlite /path_to/output/names.txt
```

### Datasets Comparison (`sf diff`)

Two SFGA files can be compared to find taxonomic differences. For example,
to compare a DwCA dataset with a checklist of names in a CSV file, first
convert both to SFGA and then run the diff:

```bash
sf from dwca /path_to/dataset1.dwca.zip /path_to/sfga/dataset1
sf from xsv /path_to/dataset2.csv /path_to/sfga/dataset2
sf diff /path_to/sfga/dataset1.sqlite /path_to/sfga/dataset2.sqlite diff.sqlite
```

The comparison can be scoped to a specific taxon in each file:

```bash
sf diff dataset1.sqlite dataset2.sqlite diff.sqlite \
    --source-taxon Plantae --target-taxon Plantae
```

### Updating SFGA Files (`sf update`)

To be able to use `sf`, SFGA files should correspond to the latest version
of the [SFGA schema]. The `update` command seamlessly converts older SFGA
versions to the latest schema, migrating all the data:

```bash
sf update /path_to/outdated/dataset.sqlite /path_to/updated/dataset
```

It can also convert a flat classification into a nested parent/child
hierarchy:

```bash
sf update flat.sqlite /path_to/tree.sqlite --add-parents
```

[Homebrew]: https://brew.sh/
[LICENSE]: ./LICENSE
[SFGA schema]: https://github.com/sfborg/sfga
[just]: https://github.com/casey/just
[releases]: https://github.com/sfborg/sf/releases
[winpath]: https://www.computerhope.com/issues/ch000549.htm
[wsl]: https://docs.microsoft.com/en-us/windows/wsl/
