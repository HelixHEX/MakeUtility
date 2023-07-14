# Directory Size Calculator

## Description
This is a simple go script that calculates the size of a directory, all of its subdirectories, and outputs to a text file. It is useful for finding out how much space a directory is taking up on your computer.

## Usage

1. Clone this repo
```bash
git clone https://github.com/HelixHEX/makeutility.git
```

2. Run the script

**Note:** When asked for the ouput file name, don't enter a file extension. The script will automatically add a .txt extension to the file name.

```bash
go run main.go
```

### Example

```bash
$ go run main.go 
Enter directory path: /Users/username/Desktop
Enter output file name: output
Total size of '/Users/username/Desktop' is 1.2 MB
```