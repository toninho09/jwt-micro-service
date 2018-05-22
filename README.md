# JWT-Micro-Service

Api utilizada para criar e verificar JWT


## Configurações

Copie os certificados para `/root/cert/app.rsa` e `/root/cert/app.rsa.pub`, o serviço já tem 
os certificados para exemplo, mas é recomendado que seja utilizado certificados diferentes 
no ambiente de produção

## Docker

Para criar a imagem da aplicação utilizar o comando

````shell
docker build -t jwt-micro-service .
````

Para rodar utilizando a imagem do docker
````shell
docker run -d --name jwt-micro-service -p 80:80 -v /you/cert/:/root/cert/ jwt-micro-service
````
 
Como alternativa pode-se usar a imagem do docker hub `toninho09/jwt-micro-service`, que pode ser baixada utilizando o comando
 
````shell
    docker pull toninho09/jwt-micro-service
````
 
 
# Serviços


### [POST|GET] /verify

Verifica se o token e válido, caso o token não seja valido, o HTTP_STATUS de retorno será 401,
caso seja válido, será retornado os claims

##### request

````json
{
	"token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNeVNlcnZpY2UiLCJkYXRhIjp7InJvdGUiOiJhZG1pbiIsInVzZXIiOjEyM30sImV4cCI6MTUyNzAwNTM2NiwiaWF0IjoxNTI3MDAxNzY2LCJpc3MiOiJBcGkiLCJzdWIiOiJNeVNlcnZpY2UifQ.SAtmA8Yf-oegPIer7LSWWZT6lR9h0oJkikhmAkneYNSt6an_D1WBUduJlJSLJDoAL86NHNfzx6-PNWV_hQfwubg95U_keEBcBiPPYSjEOHtH6n3f6duW66OgFjxQLXlB4FNhTEZod_cD5pCnjZs2s55-nVaepeZngy2Npog_3dw"	
}
````

##### response

````json
{
	"aud": "MyService",
	"data": {
		"rote": "admin",
		"user": 123
	},
	"exp": 1527005366,
	"iat": 1527001766,
	"iss": "Api",
	"sub": "MyService"
}
````

### [POST|GET] /create

Cria um jwt através dos claims enviados, os dados adicionais devem estar dentro da propriedade `data`, todos
os campos são opcionais

##### Request

````json
{
	"sub":"MyService",
	"iss":"Api",
	"aud":"MyService",
	"data":{
		"user":123,
		"rote":"admin"
	}
}
````

##### Response

````json
{
	"token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNeVNlcnZpY2UiLCJkYXRhIjp7InJvdGUiOiJhZG1pbiIsInVzZXIiOjEyM30sImV4cCI6MTUyNzAwNTM2NiwiaWF0IjoxNTI3MDAxNzY2LCJpc3MiOiJBcGkiLCJzdWIiOiJNeVNlcnZpY2UifQ.SAtmA8Yf-oegPIer7LSWWZT6lR9h0oJkikhmAkneYNSt6an_D1WBUduJlJSLJDoAL86NHNfzx6-PNWV_hQfwubg95U_keEBcBiPPYSjEOHtH6n3f6duW66OgFjxQLXlB4FNhTEZod_cD5pCnjZs2s55-nVaepeZngy2Npog_3dw"
}
````

