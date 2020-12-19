#!/bin/sh
curl --trace-ascii /dev/stdout -X  POST \
	 -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials&client_id=${client_id}&client_secret=${client_secret}"\
    "https://${host}/oxauth/restv1/token" \

#curl -v -X POST \
#	-H "Content-Type: application/x-www-form-urlencoded" \
#    -d "grant_type=client_credentials&client_id=${client_test}&client_secret=${client_secret}"\
#    "https://${host}/oxauth/restv1/token" \
#
	
#
#
#tion/x-www-form-urlencoded" \
#dentials&client_id=${client_id}&client_secret=${client_secret}" \
#o/auth/realms/practable-develop/protocol/openid-connect/token" \
#
