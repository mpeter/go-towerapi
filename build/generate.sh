#!/bin/bash

# Generates model structs -- WARNING: Hackery going on here. This is a conveience script only.
#
# Prerequisites:
# httpie -- https://github.com/jkbrzt/httpie
# gojson -- go get github.com/ChimeraCoder/gojson/gojson

OPTS=`getopt -o vhns: --long endpoint,username,password: -n 'parse-options' -- "$@"`

if [ $? != 0 ] ; then echo "Failed parsing options." >&2 ; exit 1 ; fi
echo "$OPTS"
eval set -- "$OPTS"

TOWER_USERNAME=admin
TOWER_PASSWORD=TX5KmCUKM5Sp
TOWER_ENDPOINT=http://ansible-tower/api/v1

while true; do
  case "$1" in
    -e | --endpoint ) TOWER_ENDPOINT=true; shift ;;
    -u | --username ) TOWER_USERNAME=true; shift ;;
    -p | --password ) TOWER_PASSWORD=true; shift ;;
    -- ) shift; break ;;
    * ) break ;;
  esac
done

echo TOWER_ENDPOINT=$TOWER_ENDPOINT
echo TOWER_USERNAME=$TOWER_USERNAME
echo TOWER_PASSWORD=$TOWER_PASSWORD

TYPES=( Ping Config Me Dashboard Organizations Users Projects Teams
Credentials Inventories Inventory_Scripts Inventory_Sources Groups Hosts
Job_Templates Jobs Job_Events Ad_Hoc_Commands System_Job_Templates System_Jobs
Schedules Roles Notification_Templates Notifications Labels
Unified_Job_Templates Unified_Jobs Activity_Stream )

# Save the pwd before we run anything
PRE_PWD=`pwd`

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Derive the project name from the directory
PROJECT="$(basename $BUILD_DIR)"

#rm -rf $PWD/towerapi/model
#mkdir -p $PWD/towerapi/model

for t in ${TYPES[*]};
do
  echo Unified_Job_Templates | awk '{ print tolower($1) }'  
  name_uri=`echo "$t" | awk '{ print tolower($1) }'`
  name_struct=`echo $t | sed 's/_//g'`
  #name_ucamel=`echo $name_struct
  name_lcamel=`echo $name_struct | perl -ne 'print lcfirst($_);'`

  echo "Generating model for $name_struct"
  rm -f ../towerapi/services/$name_uri.go
  cp skeleton_service.go ../towerapi/services/$name_uri.go
  sed -i '' -e "s/{{name_uri}}/$name_uri/g" ../towerapi/services/$name_uri.go
  sed -i '' -e "s/{{name_struct}}/$name_struct/g" ../towerapi/services/$name_uri.go
  sed -i '' -e "s/{{name_ucamel}}/$name_ucamel/g" ../towerapi/services/$name_uri.go
  sed -i '' -e "s/{{name_lcamel}}/$name_lcamel/g" ../towerapi/services/$name_uri.go
#  http --auth $TOWER_USERNAME:$TOWER_PASSWORD --body --pretty=format $TOWER_ENDPOINT/$name_uri/ | gojson -name $name_struct -pkg model -o ../towerapi/model/$name_uri.go
  #sleep 2
done

