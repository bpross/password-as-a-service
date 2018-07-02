test_env:
	export PASSWORDWAIT=0

prod_env:
	export PASSWORDWAIT=5

test: test_env
	go test ./handlers ./main ./password

build:
	go build ./main/paas.go

serve: prod_env
	./paas

build_and_serve: build serve
