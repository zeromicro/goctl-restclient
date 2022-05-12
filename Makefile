


test:
	go install &&cd example/ && goctl api plugin -plugin goctl-restclient="-filename test.json" -api test.api -dir .