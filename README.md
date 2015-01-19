## How to Get Your Database

### Google Authenticator

    scp root@android:/data/data/com.google.android.apps.authenticator2/databases/databases /tmp/android_ga.db
    tfa_auth_dump /tmp/android_ga.db

## Tab Completion

For no good reason (doesn't work yet):

    source tfa_auth_dump.bash
