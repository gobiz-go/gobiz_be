# boilerplate

boilerplate for beginner. Using go fiber as base framework. This bolerplate compatible with fly.io CI/CD deployment.

## env setting
Please set environtment variabel in your system:
```sh
PUBLICKEY=YOURPUBLICKEY
MONGOSTRING=YOURMONGOSTRINGACCESS
URLAPI=YOUR3rdPARTYURLAPI
CGO_ENABLED=0
```


## Folder Structure

This boilerplate have several folder with different function, such as:
* url : same as routes, this folder act to route URL in browser into controller
* config : all of apps configuration like: database, api, token.
* contoller : all of endpoint function
* model : all of type struct used in this apps
* helper : package folder with function only called by others


## Rename Apps

If you want to rename apps, please delete go.mod and go.sum file first, then type command in your terminal or cmd :

```sh
go mod init gocroot
go mod tidy
```

Then replace package import with the name gocroot/... in the main.go, controller.go dan url.go

## Github Action

inside folder .github/workflows, you might chose your cloud provider and remove .template ekstension :
1. fly.io : fly.yml | set FLY_API_TOKEN in secrets menu
2. alwaysdata.com : alwaysdata.yml | set server, username, and password in secrets menu
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/27fb3fec-be96-4868-ac94-dc243a8d4e78)  

![image](https://github.com/gocroot/alwaysdata/assets/11188109/bc223d09-bf7a-4f6c-83f4-3a273ecce1ad)
