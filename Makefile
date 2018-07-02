test_env:
	export PASSWORDWAIT=0

test: test_env
	go test ./handlers ./main ./password

build:
	go build ./main/paas.go

serve:
	./paas

build_and_serve: build serve
