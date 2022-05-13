


test:
	go install &&cd example/ && goctl api plugin -plugin goctl-restclient="-filename test.rest" -api test.api -dir .