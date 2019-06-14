# Docker and regeax/ Golang

# Intalando o docker e docker-compose!

```sh
$ sudo apt install docker.io
$ sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
$ sudo chmod +x /usr/local/bin/docker-compose
```

# Rodando o código local!

```sh
$ git clone https://github.com/ViniciusBoscardin/regexpandmicroservices.git
$ cd <pasta da aplicacao>
$ cd docker-compose
$ cd docker-compose up --build -d
```

# Requestes com POSTMAN!

```javascript
URL: http://localhost/v1/texto/match-text
{
    "texto": "<seu texto>" 
}

URL: http://localhost/v1/texto/match-number
{
    "texto": "<seu texto>" 
}

URL: http://localhost/v1/texto/find-number
{
    "texto": "<seu texto>" 
}
```
