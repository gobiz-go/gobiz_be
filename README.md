# Golang template for CI/CD in alwaysdata.com

This is a simple Golang template using Go Fibre as the base framework. It is compatible with alwaysdata.com CI/CD deployment.

## Go Boilerplate Folder Structure

This boilerplate has several folders with different functions, such as:
* URL: same as routes, this folder acts to route URL in the browser into the controller
* config: all apps configuration like database, API, token.
* controller: all of the endpoint function
* model: all of the type struct used in this app
* helper: package folder with function only called by others

## Alwaysdata.com CI/CD setting

Follow this instruction:
1. Open the menu Web>Sites>Modify
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/a95bce70-f0fc-4a74-abfa-51ba3dd543d4)
2. In the Configuration section edit command and Environment
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/d88f8fe6-08a3-4efe-9705-3ad5016b80ee)  
   Please set the environment variable in your system:
   ```sh
   MONGOSTRING=YOURMONGOSTRINGACCESS
   WAQRKEYWORD=yourkeyword
   WEBHOOKSECRET=yoursecret
   WAAPIQR=https://api.wa.my.id/api/whatsauth/request
   WAAPITXTMSG=https://api.wa.my.id/api/send/message/text
   ```
   In this menu, you will see an APPID in the title, shown as a number.
3. Go to menu REmote Access>SSH>Modify, set a very strong password and tick enable password-based login
4. Set APIKEY in Customer Area>Profile >Managing Tokens>Generate a token
5. Add sshhost, sshusername, sshpassword, sshport, apikey, appid in your github secret>action variabel
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/bc223d09-bf7a-4f6c-83f4-3a273ecce1ad)
   
## Rename Apps

If you want to rename apps, please delete (go.mod) and (go.sum) files first, then type the command in your terminal or cmd :

```sh
go mod init gocroot
go mod tidy
```

Then replace package import with the name (gocroot/...) in the (main.go), (controller.go) dan (url.go)

## Github Action

inside the folder .github/workflows, you might choose your cloud provider and remove the .template extension :
1. fly.io: fly.yml
2. alwaysdata.com: alwaysdata.yml

