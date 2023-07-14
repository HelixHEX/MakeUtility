# Proposal for Directory Size Calculator

## Goals

The goal of this project is to create a simple utility that will calculate the size of a directory and all of its subdirectories. The utility will output the results to a text file. This utility is useful for finding out how much space a directory is taking up on your computer.'

## Audience

This utility is for anyone who wants to find out how much space a directory is taking up on their computer.

## Packages

* [fmt](https://golang.org/pkg/fmt/) - To print to the console
* [os](https://golang.org/pkg/os/) - To get the directory path from the user
* [path/filepath](https://golang.org/pkg/path/filepath/) - To get the size of the directory and all of its subdirectories
* [strings](https://golang.org/pkg/strings/) - Used to ignore certain files and directories (ex: .git)
* [go-bytesize](https://github.com/inhies/go-bytesize) - Used to convert the size of the directory to a human readable format
* [progressbar](https://github.com/schollz/progressbar) - Used to display a progress bar while calculating the size of the directory

## Task Description

[x] Get the directory path from the user: 5 minutes

[x] Calculate the size of the directory and all of its subdirectories (excluding hidden files and directories): 1 hour

[x] Get output file name from the user: 5 minutes

[x] Output the results to a text file: 30 minutes

[x] Display loading bar while calculating the size of the directory: 1 hour

[x] Display the size of the directory in a human readable format: 30 minutes