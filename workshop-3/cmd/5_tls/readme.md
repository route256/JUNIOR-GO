#### Создать сертификат
````
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
````

#### Примеры

````
curl -k --cacert ./server.crt https://localhost:9000
curl -k -v localhost:9000
````
