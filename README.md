# JWT-Micro-Service

Api utilizada para criar e verificar JWT


## Configurações

As configurações são feitas utilizando variáveis de ambiente 

* PORT : Configura a porta que o serviço estará disponível, default  `8080` 
* ENV : Define o ambiente em que o sistema está rodando, Valores `DEV` , `PROD`

Copie o arquivo de autenticação do firebase para o arquivo `api.json`

## Docker

Para criar a imagem da aplicação utilizar o comando

````shell
docker build -t auth .
````

Para rodar utilizando a imagem do docker
````shell
docker run -d --name micro-auth-firebase -p 80:8080 -v /you/api.json:/root/api.json auth
````
 
 Como alternativa pode-se usar a imagem do docker hub `toninho09/firebase-auth-microservice`, que pode ser baixada utilizando o comando
 
 ````shell
 docker pull toninho09/firebase-auth-microservice
 ````
 
 
# Serviços


### [POST] /verify

Verifica se o token e válido, caso o token não seja valido, o HTTP_STATUS de retorno será 401

##### request

````json
{
	"token":"eyJhbGciOiJSUzI1NiIsImt......"	
}
````

##### response

````json
{
	"iss": "https://securetoken.google.com/xxxxxxx",
	"aud": "xxxxx",
	"exp": 1513713402,
	"iat": 1513709802,
	"sub": "qPD5IBB....",
	"uid": "qPD5IBB...."
}
````

### [POST] /get-user

Obtém as informações do usuário pelo firebase

##### Request

````json
{
	"token":"eyJhbGciOiJSUzI1NiIsImt......"	
}
````

##### Response

````json
{
	"displayName": "Display name",
	"email": "email@email.com",
	"photoUrl": "https://.../photo.jpg",
	"localId": "localId",
	"CustomClaims": null,
	"Disabled": false,
	"EmailVerified": true,
	"ProviderUserInfo": [
		{
			"displayName": "Display name",
			"email": "email@email.com",
			"photoUrl": "https://.../photo.jpg",
			"providerId": "google.com"
		}
	],
	"UserMetadata": {
		"CreationTimestamp": 1513099763000,
		"LastLogInTimestamp": 1513100854000
	}
}
````

