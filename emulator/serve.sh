#!/bin/bash

set -euo pipefail

export data_directory="../emulator/emulator-data"

#cd /usr/src/firebase/functions
#( npm i 2>&1 )

cd /usr/src/firebase/firestore
( npm i 2>&1 )

#cd /usr/src/firebase/storage
#( npm i 2>&1 )

# install firebase and run the tests in the correct folder for rules etc

cd /usr/src/firebase
( npm i 2>&1 )

( firebase emulators:start --project=$FIREBASE_PROJECT --import=$data_directory --export-on-exit=$data_directory --only=$EMULATORS_USED) &
firebase_pid=$!

# sleep as emulators need to start or the tests crash a bit
sleep 20s

( nginx ) &
nginx_pid=$!

( npm run start 2>&1 ) &
npm_pid=$!

:stop() {
    # remove old firestore data and recreate, export on exit was not working properly nor this without clearing folder first
    ( rm -rf $data_directory && firebase emulators:export $data_directory )
}

#Execute command
"${@}" &

#Wait
wait $!

# fire stop on container exit
trap :stop INT TERM SIGTERM

wait $firebase_pid $nginx_pid $npm_pid