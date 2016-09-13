#!/bin/bash
set -e
set -o
set -x

# Generates model structs -- WARNING: Hackery going on here. This script is bad, and you should feel bad.
#
# Prerequisites:
# httpie -- https://github.com/jkbrzt/httpie
# gojson -- go get github.com/ChimeraCoder/gojson/gojson

OPTS=`getopt -o eup: --long endpoint,username,password: -n 'parse-options' -- "$@"`

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

# AuthToken
LIST_TYPES=(Activity_Stream Dashboard Ping Config Me Ad_Hoc_Commands
Organizations Users Projects Teams Credentials Inventories Inventory_Scripts
Inventory_Sources Groups Hosts Job_Templates Jobs Job_Events
System_Job_Templates System_Jobs Schedules Roles Notification_Templates
Notifications Labels Unified_Job_Templates Unified_Jobs )

#LIST_TYPES=( Dashboard )

# Save the pwd before we run anything
PRE_PWD=`pwd`

# Determine the build script's actual directory, following symlinks
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

# Derive the project name from the directory
PROJECT="$(basename $BUILD_DIR)"

base_cmd="http --session=`date "+%Y%m%d"` --auth $TOWER_USERNAME:$TOWER_PASSWORD --body --pretty=format http://ansible-tower"

for t in ${LIST_TYPES[*]};
do
  name_uri=`echo "$t" | awk '{ print tolower($1) }'`
  name_struct=`echo $t | sed 's/_//g'`
  name_lcamel=`echo $name_struct | perl -ne 'print lcfirst($_);'`

  echo "Generating model for $name_struct"
  # Generate JSON Example for LIST type
  entity_list_file="../towerapi/model/${name_uri}.json"
  entity_list_json=`$base_cmd/api/v1/$name_uri/?page_size=1\&page=1\&order_by=id | jq .`
  # re-create the output file
  rm -f $entity_list_file && echo $entity_list_json > $entity_list_file
  # Derive the URL and type of the result entry in list
  entity_url=`jq 'try .results[0].url catch ""' $entity_list_file | tr -d '"'`
  entity_type=`jq 'try .results[0].type catch ""' $entity_list_file | tr -d '"'`
  # Derive struct name
  entity_struct=`echo $entity_type | sed 's/_/ /g' | perl TitleCase.pl | sed 's/ //g'`
  # Default null values to string
  sed -i '' -e "s/null/\"string\"/g" $entity_list_file

  # Generate JSON Example for CRUD type based on List type
  entity_go_file="../towerapi/model/${name_uri}.go"
  rm -f $entity_go_file
  if [[ "${entity_type}" == "null" ]]; then
    gojson -input $entity_list_file -name ${name_struct} -pkg towerapi -o $entity_go_file
  else
    jq 'del(.results[])' $entity_list_file | gojson -name ${name_struct} -pkg towerapi | sed -e "s/\[\]interface{}/\[\]${entity_struct}/" > $entity_go_file
    $base_cmd$entity_url | jq . | sed -e "s/null/\"string\"/g" | gojson -name $entity_struct -pkg towerapi | sed -e 's/package towerapi//' >> $entity_go_file
  fi
done

