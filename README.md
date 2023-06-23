# Go Markdown Block Exporter
Cli tool to extract code blocks from Readme files, saves them onto separate files.

# Why
This tools makes it easy to check code blocks inside markdown files. 

# Instalation
First, install the [go programming language](https://go.dev/doc/install). 

Then run:

```sh
go install github.com/djrmarques/gmbe
```

Verify installation with
 
 ```sh
gmbe --help
 ```
 
 
# Usage

Use the following command:

```sh
gmbe -f Readme.md
```

This command will create a folder called `Readme`, and inside there will be files with the code blocks.

# Usage options
There are a few usage flags that can be used:

It's possible to specify the output folder with:
``` sh
gmbe -o folder -f Readme.md # This will save the files inside this folder
```

It's also possible to join all the codeblocks of the same type inside the same file
``` sh
gmbe -j -f Readme.md # This will join all code blocks of the same type in the same file
```



