# DirScan

<p align="center"><em>On-going</em></p>

DirScan scans your current directory and its subdirectories for a specific file name. It performs a case-insensitive search and displays not only the path to the desired file but also files with similar names.

## Installation

Clone the repository in your home directory, ~:

```sh
git clone https://github.com/Axelvazslima/dirscan
```

## Setup

Run the setup file:

```sh
go run setup/setup.go
```

This will create an alias for the program in your bash, allowing it to determine your current directory and start the search from there.

## Usage

Run the program with the following command:

```sh
godirscan <file_name>
```

It gets your input directly from your command line. Place it right after the `godirscan` command.
