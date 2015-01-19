## tfa_auth_dump

[![Build Status](https://travis-ci.org/solarkennedy/tfa_dump_auth.svg?branch=master)](https://travis-ci.org/solarkennedy/tfa_dump_auth)

`tfa_auth_dump` is a command line utility to take a Two-Factor-Auth application database
and dump it into QR codes you can re-import into another phone.

This is useful for:

 * Migrating to a new phone
 * Copying codes from one phne to a backup phone
 * Printing and saving the QR codes in a safe place? (not everyone supports one-time codes)

## Screenshot

![tfa_auth_dump_screenshot](https://raw.githubusercontent.com/solarkennedy/ether_house/master/tfa_auth_dump_screenshot.png)

## Installation

Assuming you have your GOPATH setup and stuff.

    go get github.com/solarkennedy/tfa_auth_dump

## How to Get Your Database

### Google Authenticator

    scp root@android:/data/data/com.google.android.apps.authenticator2/databases/databases /tmp/android_ga.db
    tfa_auth_dump /tmp/android_ga.db

## Tab Completion

For no good reason (doesn't work yet):

    source tfa_auth_dump.bash
