# takegur
steal all the images for reposts! takegur can save wallpapers, black twitter, mobile wallpapers,
dickbutt images, epic fails, reactions (gifs or images), Stay Classy, and Darwin Awards from the Imgur galleries.

One thing to note, it doesn't save anything specific. You run it, `takegur wallpapers`, and it will
look in User Sub and the FP for whatever you are trying to save.

## To download

1. Look above, there's some text that says "Releases" with a number before it. Click it.
2. For the highest version number:
    * If you are using Windows 64-bit, download the takegur-windows-amd64 version.
    * if you are using Windows 32-bit, download the takegur-windows-386 version.
    * If you are using OSX, download the takegur-darwin-amd64 version.
    * If you are using a _REALLY_ old version of OSX, you might need the takegur-darwin-386 version.
    * Linux peeps, you know which is yours.

## General Instructions

1. Unzip the file to the directory you want to save reposts to.
2. Open your command line. Windows, hit WinKey+R, type `cmd`. OSX, look for the Terminal application.
3. In your command line, change to the location where you saved it.
    * Windows: `cd c:\location\you\saved\it`
    * OSX: `cd ~/location/you/saved/it`
4. Save something! `./takegur wallpapers` will save all the wallpapers on the front page.
5. Profit.

If you have issues/problems/feature requests, please file an issue. You can find the link to the issues
on the right hand side.

## Usage

```bash
picturesarenice$ ./takegur --help
NAME:
   takegur - used to download the awesomeness of imgur.

USAGE:
   takegur [global options] command [command options] [arguments...]
   
VERSION:
   1.0
   
AUTHOR(S):
   picturesarenice emi8ly 
   
COMMANDS:
   black-twitter        when u tell her 2 stop n she keeps suckin...
   wallpapers           ALL THE DESKTOP BACKGROUNDS ARE BELONG TO YOU.
   mobile               for when you need to keep looking at your phone to avoid meetings.
   dickbutt             please don't ever use this.
   stay-classy          you're a sick bastard.
   darwin-awards        let's watch some stupid people!
   fails                ouch.
   mrw                  your reaction when...
   help, h              Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```


### Advanced

If you already have Go installed, you can use `go get https://github.com/mynameismevin/takegur`. Easiest way
to go, and will help you live through upgrades.

To call `takegur` from anywhere (if you didn't install via `go get`), just add it to your path!

## Contributing

takegur uses regex in order to query imgur gallery titles. To add a new image series (like Michael Cera),
just add a regex that would match anything with Michael Cera in the title, and then add it to the `cli.[]Commands{}`
section.

## To Do List

* Add functionality to read user accounts so Michael Cera photos can be downloaded.
* Make the `Reader()` extensible.
* Optimise the absolutely terrible regex's to catch more matches.

## Authors

Made with <3 by picturesarenice and emi8ly.