## tfa_auth_dump

[![Build Status](https://travis-ci.org/solarkennedy/tfa_auth_dump.svg?branch=master)](https://travis-ci.org/solarkennedy/tfa_auth_dump)

`tfa_auth_dump` is a command line utility to take a Two-Factor-Auth application database
and dump it into QR codes you can re-import into another phone.

This is useful for:

 * Migrating to a new phone
 * Copying codes from one phne to a backup phone
 * Printing and saving the QR codes in a safe place? (not everyone supports one-time codes)

## Screenshot

![tfa_auth_dump_screenshot](https://raw.githubusercontent.com/solarkennedy/tfa_auth_dump/master/tfa_auth_dump_screenshot.png)

## Installation

Assuming you have your GOPATH setup and stuff.

    go get github.com/solarkennedy/tfa_auth_dump

## How to Get Your Database

### Google Authenticator

    scp root@android:/data/data/com.google.android.apps.authenticator2/databases/databases /tmp/android_ga.db
    tfa_auth_dump /tmp/android_ga.db

### Duo Security

# New Style

    scp root@android:/data/data/com.duosecurity.duomobile/files/duokit/accounts.json /tmp/duo_accounts.json
    tfa_auth_dump /tmp/duo_accounts.json

or

    adb root
    adb pull /data/data/com.duosecurity.duomobile/files/duokit/accounts.json /tmp/duo_accounts.json
    tfa_auth_dump /tmp/duo_accounts.json

# Old Style

    scp root@android:/data/data/com.duosecurity.duomobile/databases/databases /tmp/duo.db
    tfa_auth_dump /tmp/duo.db
