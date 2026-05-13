# 💾 Mock Data Faker (and Sender)

This is the repository for the 3SI Lab's take-home assessment's mock data sender.  This repository isn't going to be shipped off in the assessment, but will 
contain the source code behind the `send-<system>.exe` executable in the assessment's files - in case the assessment ever changes in the future and / or the code needs to
be referenced again at a later date for reasons currently unknown to us.

Plus - to any potential colleagues who might be joining us - welcome to what actually goes on behind the scenes!  If you're here because you have an inquiry about how 
`send-<system>.exe` was built and / or want to verify that it's not malware, feel free to browse through the executable's code (which is all written in Golang - the language
Kevin chose when he was building the executable)!  As always, do also reach out to us via the HR personnel who initially contacted you if you have any other questions pertaining to this assessment in general!

# 🔧 Building the Executable

*NB: If you do not already have Golang installed on your machine, please visit [Go's official webpage](https://go.dev/) to do so first.  Otherwise, your machine will not be 
able to do any of the below steps*

The steps might differ based on the machine that you're using, but rest assured - it's still simple enough across the board no thanks to Golang's capabilities (and the fact that this repository doesn't use CGO)!  Either way, ensure that you've pulled the repository onto your machine first:

However, note that if you *are* using a Linux or a MacOS machine, you will need to change the file's permissions with `chmod +x <name of your program>` for it to be executable!

```
git pull https://github.com/JamestheCog/data_faker.git
cd data_faker
```

And only then - follow the below commands depending on the operating system that you're using!

## 🪟 ...for Windows Computers

You can use Golang's `go build` like so (don't forget the `.exe` extension):

```
go build -o windows-send.exe
```

## 🍎 ...for MacOS Machines

This is a little trickier as Apple machines have different chip generations: Intel-based ones and Apple Silicon's M1 to M3 chips.  So, you'll need to adjust Go's compilation settings accordingly before running `go build`:

1. **Intel Chip MacOS-es**
   ```
   $env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o send-mac
   ```

2. **M1 to M3 MacOS-es**
   ```
   $env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o send-mac-m13
   ```

And you should be able to run those said binaries on your machine!

## 🐧 ...for Linux Machines

You will need to re-configure Go's compilation environment like MacOS users would, but the simplicity still stands:

```
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o send-linux
```

# 🤔 Usage

Use `send-<system>.exe` like you would with any other CLI application.  Though, running `send-<system>.exe` as is will merely bring up a help pane before exiting with status 1.
