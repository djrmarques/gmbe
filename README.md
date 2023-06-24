# Go Markdown Block Exporter
Cli tool to extract code blocks from Readme files, saves them onto separate files.

Go Markdown Block Extractor will extract all the codeblocks from markdown files and save them as their own separate files.
This allows to run checks like linting, formating or any other custom checks to make sure that the code blocks in the markdown files,
which most likely are some sort of documentation, are actually valid. 

# Instalation
First, install the [go programming language](https://go.dev/doc/install). 

Then run:

```sh
go install github.com/djrmarques/gmbe@latest
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



