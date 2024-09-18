# gotree
*   its recursive directory listing program that produces a depth-indented listing of files.

## installation

*   ### Option 1: Clone and Build
```bash
git clone https://github.com/mugund10/gotree
cd gotree
go build -o gotree main.go
```
*   ### Option 2: Install Directly using Go
```bash
go install github.com/mugund10/gotree@latest
```

## usage 

```bash
./gotree /path/to/directory
```
or
```bash 
//only if you installed with go
$(go env GOPATH)/bin/gotree  /path/to/directory
```
###  example

```bash
$ ./gotree /SampleDirectory


SampleDirectory
├──── Project
│     ├──── file1.txt
│     └──── file2.txt
├──── Images
│     ├──── image1.png
│     └──── image2.jpg
└──── Notes
      └──── notes.txt

 directories : 3 , files : 5 

```