# creating-web-services-with-go
   
Used following plural sight course : https://app.pluralsight.com/library/courses/creating-web-services-go . 



## run MySQL via docker:

```
    docker run --rm -it -d \
        -v ${pwd}/data:/var/lib/mysql \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=password123 \
        mysql:latest
```