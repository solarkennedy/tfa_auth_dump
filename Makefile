.PHONY: test run clean fmt

test:
	go test -v -bench=.

tfa_auth_dump:
	go build .

run: tfa_auth_dump
	./tfa_auth_dump example.db

clean:
	rm tfa_auth_dump

fmt:
	go fmt .

adb: tfa_auth_dump
	adb shell su -c cat /data/data/com.google.android.apps.authenticator2/databases/databases > database.sqlite
	./tfa_auth_dump database.sqlite