#!/bin/bash

set -e

if [ "$#" -ne 2 ]; then
	echo "Usage: $0 <app name> <buildpack url>"
fi

heroku apps:destroy --confirm $1 --app $1 || true
heroku apps:create -s cedar -b $2 $1
git push heroku master
