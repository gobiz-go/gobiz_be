# Golang template for CI/CD in alwaysdata.com

This is a simple Golang template using Go Fibre as the base framework, and mongodb.com as database host. It is compatible with alwaysdata.com CI/CD deployment.

## MongoDB Preparation

The first thing to do is prepare a Mongo database using this template:
1. Sign up for mongodb.com and create one instance of Data Services of mongodb.
2. Download mongo-compass, connect with your mongo string URI from mongodb.com
3. Create database name iteung and collection reply
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/23ccddb7-bf42-42e2-baac-3d69f3a919f8)  
4. Import [this json](https://whatsauth.my.id/webhook/iteung.reply.json) into reply collection.
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/7a807d96-430f-4421-95fe-1c6a528ba428)  

## Go Boilerplate Folder Structure

![image](https://github.com/gocroot/alwaysdata/assets/11188109/aa3a8162-3aa9-4a55-be6c-2e0caf5dcfef)  

This boilerplate has several folders with different functions, such as:
* .github: GitHub Action yml configuration.
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
   ```
   In this menu, you will see an APPID in the title, shown as a number and a home folder used in the github secrets variable.
3. Go to menu REmote Access>SSH>Modify, set a very strong password and tick enable password-based login
4. Set APIKEY in Customer Area>Profile >Managing Tokens>Generate a token
5. Add sshhost, sshusername, sshpassword, sshport, apikey, appid and folder in your GitHub secret>action variable
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/5cc1e831-49d5-47d1-9486-d6f0f748a963)  

   
## Rename Apps

If you want to rename apps, please delete (go.mod) and (go.sum) files first, then type the command in your terminal or cmd :

```sh
go mod init gocroot
go mod tidy
```

Then replace package import with the name (gocroot/...) in the (main.go), (controller.go) dan (url.go)

## Others Provider Github Action

inside the folder .github/workflows, you might choose your cloud provider and remove the .template extension :
1. fly.io: fly.yml
2. alwaysdata.com: alwaysdata.yml

