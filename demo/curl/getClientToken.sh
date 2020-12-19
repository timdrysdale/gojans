#!/bin/sh
client_token=$(curl -s -X POST\
	 -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials&client_id=${client_id}&client_secret=${client_secret}"\
    "https://${host}/oxauth/restv1/token" \
	| jq -r '.access_token')
echo $client_token

