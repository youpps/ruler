## Ruler is a telegram bot for remote managing your Windows computer.

> ### Ruler allow you to open files, download photo and video from telegram, look at system directories, check files content.

<br/>

## How to try it

#### 1. Clone repository: `git clone https://github.com/youpps/ruler`

#### 2. Run building: `go build -o build/main.exe cmd/main.go`

#### 3. Then run `main.exe` from build directory

#### 4. Follow instructions of the bot

<br/>

### You can use these comands to manage your Windows system:

      /ls {PATH}                          - to get all file from PATH
      /open_file {ABSOLUTE_PATH}          - to open a file
      /get_file_body {ABSOLUTE_PATH}      - to get file content
      /shutdown                           - to shutdown system
      /destroy                            - to close bot deleting all its data

<br/>

### If you send photo or file to Ruler, it will put it at `USER_HOME_DIR/jupiter` directory. If there is not the such directory in your system, it will create it.

<br/>

### You could also send a message to Ruler. Then, it will execute it like a command in the console.

<br/>

## Additional info

> #### `CTRL + C` and `/destroy` command is equal. Execution either of them will delete USER_HOME_DIR/jupiter

> #### When you send a photo or a video to bot your file should match `FILENAME.EXT`.
