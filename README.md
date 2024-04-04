# ann2html
## Description
A small program that uses Kindle's vocabulary builder and turns it into a usable HTML page.
![example of program](usage.png)


## Features
* Single executable, portable install.
* Pressing b on the page will bookmark your current position.
* The template file's css and javascript can be easily modified.
* Automatically bolds selected word from Vocabulary builder(It will skip it words that vocabulary builder deconjugates and it can't match)

## Installation
Download your operation system specific release,extract it and move the folder wherever you would like it, ann2html does not need to be installed in any specific location.

## Usage
Copy your vocab.db file from your kindle which should be located at

| Operation System | vocab.db location |
| ---------------- | ----------------- |
| Windows | KINDLEDRIVELETTER:\system\vocabulary\vocab.db |
| MacOS | MOUNTPOINT/system/vocabulary/vocab.db |
| Linux | MOUNTPOINT/system/vocabulary/vocab.db |

Where KINDLEDRIVELETTER is the drive letter that Windows assigns to your Kindle and MOUNTPOINT being the mount point that MacOS or Linux assigns to your Kindle.

Copy this vocab.db into the folder with the ann2html executable and run ann2html. It should generate an edit.html file which has all your annotations. Everytime you run ann2html it will check to see if your vocab.db has changed from the last time you ran ann2html and if it has it will regenerate, showing only your new annotations.


## Configuration
The program can be configured in two ways, by directly editing the configuration file or setting the environment variable, with environment variables taking precedence.

### Default configuration values

| Option | default value | explanation |
| --------------- | --------------- | --------------- |
| ANN2HTML_CONFIG |NOTSET | env variable to change the location of your config file. |
| ANN2HTML_VOCABDB | vocab.db | location of your vocab.db.  |
| ANN2HTML_NUM | 0 | unix timestamp of the last annotation, this is changed by the program.|
| ANN2HTML_TEMPLATE | template.html | File to use as a template for your output file.|
| ANN2HTML_OUTPUT | edit.html | where you want your annotation to be written, this file will be replaced on each run. |
| ANN2HTML_LNG| ja | list of languages you want to get annotations from, split by , example:ja,en|

### Windows
Windows hides the vocab.db file in a odd way, so if you are on windows go to your kindle's drive and then search for vocab.db, then copy vocab.db into the folder where the ann2html program is and it should work. 

### Macos
Macos does not allow running unsigned executables easily. To run this open up your terminal and `cd` into the folder with the ann2html executable. Then it can be ran running `./ann2html`


### Linux script
For linux users there is an optional script called ann which allows easy mounting of your kindle. To use the script add your kindle's uuid to it,make it executableand add it to your path.This script does depend on udisksctl so make sure to install it.

To find your uuid run ```ls -l /dev/disk/by-uuid``` and then connect your kindle and run it again.
```
ls -l /dev/disk/by-uuid
total 0
lrwxrwxrwx 1 root root 10 Nov 10 00:00 3472-5482 -> ../../nvme
lrwxrwxrwx 1 root root 10 Nov 10 00:00 3472-54821212 -> ../../nvme2
PLUG IN KINDLE AND RUN AGAIN
ls -l /dev/disk/by-uuid
total 0
lrwxrwxrwx 1 root root 10 Nov 10 00:00 3472-5482 -> ../../nvme
lrwxrwxrwx 1 root root 10 Nov 10 00:00 3251-54821212 -> ../../nvme2
lrwxrwxrwx 1 root root 10 Nov 10 00:00 3582-6578 -> ../../sda
in this example the UUID you want is 3582-6578 
```
Then in the script called ann add your kindle uuid to the following line.
```
kindle_UUID=
```
```
kindle_UUID=3582-6578
```
Then run chmod +x ann
and then move ann anywhere in your path, now you can just run ann and your kindle will be mounted and ann2html will be ran!



