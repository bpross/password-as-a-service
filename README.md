# Password-as-a-service
Small HTTP server written in Go to serve hashed passwords

# Endpoints
- `/hash`: Returns URL encoded SHA512 hash of a string. 
Ex: `curl -d "password=angryMonkey" http://localhost:8080/hash`
Leaves socket open for 5 seconds before responding, as per the spec. This is configurable by setting the ENV variable `PASSWORDWAIT`

- `/shutdown`: Gracefully shutsdown the service
Ex: `curl -X POST -d http://localhost:8080/shutdown`

- `/stats`: Retrieve stats about hash requests
Ex: `curl http://localhost:8080/stats`

# Running, tests, etc
- To build the server: `make build`
- To build and run the server: `make build_and_serve`
- To run the server `make serve` (must before before running)
- To run tests: `make test`

## Verification of Specs
- Part 1: 
	- `make part_one_proof`
- Part 2: 
	- `make build_and_serve`
	- then run `curl -d "password=angryMonkey" http://localhost:8080/hash` in another window. 
	- To test multiple requests, I used [apache bench](https://httpd.apache.org/docs/2.4/programs/ab.html). 
	- Sample `ab` command: `ab -b post_data.txt -T application/json -c 100 -n 100 http://localhost:8080/hash` `post_data.txt` is included in the root directory of this repo.
- Part 3: 
	- `make_build_and_serve`
	- then run the ab command from above
	- followed by `curl -X POST http://localhost:8080/shutdown` in another window.
- Part 4: 
	- `make_build_and_serve`
	- then run the ab command from above
	- followed by `curl http://localhost:8080/stats`

	
# Resources used
No code is written in a vaccuum. Here are a list of resources that I used for this project (in no particular order):

- [graceful shutdown example](https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2)
- [stats middleware] (https://github.com/thoas/stats)
- [http handler testing] (https://blog.questionable.services/article/testing-http-handlers-go/)