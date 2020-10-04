this program was written for Linux/MacOS.
This program does not work on Windows.

if you want to run on windows, edit / to \ of path.

n-leave-ban that is discord bot bans user when user leaves server n times.

rename .env.example to .env and edit it.

```
ENV=production  
TOKEN=aaaaaaaaaaaaaaaaa.aaaaa
LEAVE_MAX_COUNT=3
CLIENT_ID=11922960
```
this is a sample.  
ENV should be set to production or development or test.  
TOKEN is a area of to imput your discord bot token  
LEAVE_MAX_COUNT means when a user leaves the server N times, ban the user.  
CLIENT_ID should be set your discord bot's client id. but this is not used, so you don't know your bot's client id, then you can imput random string.  

### used libs
- github.com/joho/godotenv
- github.com/jinzhu/gorm
- github.com/mattn/go-sqlite3 
