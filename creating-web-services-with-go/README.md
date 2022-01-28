# creating-web-services-with-go
   
Used following plural sight course : https://app.pluralsight.com/library/courses/creating-web-services-go . 

## inventory-mgmt project
This project is copied from exercise file section, its written in angular to test our GO APIs.

### Required packages
```
$ sudo npm install -g @angular/cli
```

### Install Project dependencies

```
# go to root of the project
$ cd inventory-mgmt

$ npm install
```

### Run the server
```
$ ng serve --open
```


## run MySQL via docker:

```
$ docker run --rm -it -d \
    -v ${pwd}/data:/var/lib/mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=password123 \
    mysql:latest
```