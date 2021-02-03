#!/bin/sh

printf "\n"
printf "Removing remaining socket and previous rose DB on the filesystem..."
printf "\n"
rm /tmp/rose.sock
rm -r /Users/macbook/.rose_db

cd /Users/macbook/go/src/rose/main

printf "Starting server..."
printf "\n"
go run main.go &

printf "\n"
printf "Preparing to run the tests..."
printf "\n"

sleep 2

cd /Users/macbook/go/src/rose/server

printf "\n\n"
go test

printf "\n\n"
printf "Killing the server process..."
printf "\n"

SERVER_PID=$(/bin/ps -e | grep "/var/folders/.*/exe/main" | grep -v 'grep' | awk '{print $1}')
kill -9 $SERVER_PID

printf "Tests finished!\n"

exit 0