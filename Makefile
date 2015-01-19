.PHONY: test run deps clean fmt

test:
	go test -v -bench=.

run:
	go build . 
	./qr_auth_dump example.db

deps:
	godep get
	godep save

clean:
	rm qr_auth_dump

fmt:
	go fmt .

