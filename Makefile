.PHONY: test run deps clean fmt

test:
	go test -v -bench=.

run:
	go build . 
	./tfa_auth_dump example.db

deps:
	go get -t .

clean:
	rm tfa_auth_dump

fmt:
	go fmt .

