#!/bin/bash

echo "Copiling started..."
go build main.go
echo "Copiling complete."
echo "Trying to launch program"
./main.exe
echo "Program exited"