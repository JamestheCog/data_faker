# 💾 Mock Data Faker (and Sender)

This is the repository for the 3SI Lab's take-home assessment's mock data sender.  This repository isn't going to be shipped off in the assessment, but will 
contain the source code behind the `./send_data.exe` executable in the assessment's files - in case the assessment ever changes in the future and / or the code needs to
be referenced again at a later date for reasons currently unknown to us.

Plus - to any potential colleagues who might be joining us - welcome to what actually goes on behind the scenes!  If you're here because you have an inquiry about how 
`./send_data.exe` was built and / or want to verify that it's not malware, feel free to browse through the executable's code (which is all written in Golang - the language
Kevin chose when he was building the executable)!  As always, do also reach out to us via HR if you have any other questions pertaining to this assessment in general!

# 🔧 Building the Executable

*NB: If you do not already have Golang installed on your machine, please visit [Go's official webpage](https://go.dev/) to do so first.  Otherwise, your machine will not be 
able to do any of the below steps.*

Building your own `./send_data.exe` is easy no thanks to Go's capabilities:

1. Pull this repository off GitHub and `cd` into it:
   ```
   git pull <https://github.com/JamestheCog/data_faker.git>
   cd data_faker
   ```

2. Build the project with `go_build`:
   ```
   go build -o send_data.exe
   ```

And that's it!  It's that simple!

# 🤔 Usage

Use `send_data.exe` like you would with any other CLI application.
