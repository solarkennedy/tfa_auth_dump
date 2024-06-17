.PHONY: test run clean fmt

test:
	go test -v -bench=.

run:
	go build . 
	./tfa_auth_dump example.db

clean:
	rm tfa_auth_dump

fmt:
	go fmt .

