# Go Readme Block Exporter
Cli tool to extract code blocks from Readme files, saves them onto separate files.

# Why
This tools makes it easy to check code blocks inside markdown files. 

# Instalation
Install the [go programming language](https://go.dev/doc/install). 
Afterwards run:

```sh
go get asdasd
```

 Verify installation with
 
 ```sh
grbe --help
 ```
 
# Usage

Use the following command:

```sh
grbe Readme.md
```

This command will create a folder called `Readme`, and inside there will be files with the code blocks.

# Usage options
There are a few usage flags that can be used:

It's possible to specify the output folder with:
``` sh
grbe -o folder Readme.md # This will save the files inside this folder
```

It's also possible to join all the codeblocks of the same type inside the same file
``` sh
grbe -j Readme.md # This will join all code blocks of the same
```

Finally, this tool can also run recursively

``` sh
grbe -R .
```


