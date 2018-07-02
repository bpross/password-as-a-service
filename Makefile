PACKAGES = ./handlers ./main ./password ./stats
MAIN = ./main/paas.go
BINARY = ./paas

test_env:
	export PASSWORDWAIT=0

prod_env:
	export PASSWORDWAIT=5

test: test_env
	go test -v $(PACKAGES) 

test_cov: test_env
	go test -v -covermode atomic $(PACKAGES) 

build:
	go build -v $(MAIN) 

serve: prod_env
	./paas

build_and_serve: build serve

part_one_proof:
	go test ./password -run TestHashPassword_AngryMonkey
