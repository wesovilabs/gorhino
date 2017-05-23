#!/bin/bash

clear


curl http://localhost:9546/api/v0/domains \
-H "Content-Type:application/json" \
--data '{ "UID":"taurus",  "Name":"Cron Application",  "Settings":{  "DomainType": "MongoDB",  "Sandbox":"true",  "Credentials":{  "Username": "root",  "Password": "secret"},  "Hostname": "localhost:27047", "Database":"test", "Collection": "taurus-app-demo"}}' \
-i -v