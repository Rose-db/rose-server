#!/bin/sh

printf "\n"
printf "Removing remaining socket and previous rose DB on the filesystem..."
printf "\n"
rm /tmp/rose.sock
rm -r /gouser/.rose_db

cd /gouser/go/src/rose/main

printf "Starting server..."
printf "\n"
go run main.go &

printf "\n"
printf "Preparing to run the tests..."
printf "\n"

sleep 2

cd /gouser/go/src/rose/server

printf "\n\n"
go test

printf "\n\n"
printf "Killing the server process..."
printf "\n"

SERVER_PID=$(/bin/ps -elf | grep "/tmp/.*/exe/main" | grep -v 'grep' | awk '{print $4}')
kill -9 $SERVER_PID

printf "Tests finished!\n"

exit 0
