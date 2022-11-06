#!/bin/sh
pname=devLog
rm -f $pname
env GOOS=linux GOARCH=amd64 go build -o $pname
sleep 1
tar -cvf $pname.tar $pname
sleep 1
mv $pname.tar ~/Downloads/
rm -f $pname

