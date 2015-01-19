test:
	go test -v -bench=.

run:
	go build . 
	@run

deps:
	godep get
	godep save

clean:
	rm qr_auth_dump

fmt:
	go fmt .

