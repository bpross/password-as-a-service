test_env:
	export PASSWORDWAIT=0

test: test_env
	go test ./handlers ./main ./password
