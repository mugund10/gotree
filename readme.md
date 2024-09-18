# gotree
*   its recursive directory listing program that produces a depth-indented listing of files.

## installation
```bash
git clone https://github.com/mugund10/gotree
cd gotree
go build -o gotree main.go
```
## usage 

```bash
./gotree /path/to/directory
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