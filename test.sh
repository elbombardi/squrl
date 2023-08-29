#!/bin/bash

for i in {1..1000} ; 
do
    curl -X 'POST'   'http://[::]:8080/v1/accounts'   -H 'accept: application/json'   -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMTM1MzEsImlzcyI6InNxdXJsIiwidXNlciI6ImFkbWluIn0.8QEXhgWrLALs5Tt2VGSD9W0SMX-HNvEhquxW2STm1vs'   -H 'Content-Type: application/json'   -d '{"email": "account0@gmail.com","username": "account'$i'"}' &
done
echo "done!"