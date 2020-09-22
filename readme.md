## devops-deploy
基于ansible的支持多版本多产品的自动话部署页面，离线部署。demo   

包括了前端demo及后端部分实现   

前端基于vue3.0+typescript 

后台基于beego


### dev

```shell script
git clone https://github.com/lafer-m/devops-deploy.git
# vendor依赖
cd devops-deploy
go mod vendor
go run main.go

# dev web-src
cd devops-deploy/web-src
yarn install
yarn serve
```


License MIT