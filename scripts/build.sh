#!bin/bash

rm ./bookingql*
cd server;
go build -o bookingql.exe;
mv bookingql ../
