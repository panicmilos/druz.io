#!/bin/bash

# Recreate config file
rm -rf ./env-config.js
touch ./env-config.js

# Check if NODE_ENV is production
# Using .env for production and .env.local for development
if [[ $NODE_ENV = 'production' ]]
then
  file_name=.env
else
  if [ ! -f .env.local ]; then
    echo ".env.local not found! Reading from .env"
    file_name=.env
  else 
    file_name=.env.local
  fi
fi

# Add assignment 
echo -n "window._env_ = {" >> ./env-config.js

# Read each line in .env file
# Each line represents key=value pairs
while read -r line || [[ -n "$line" ]];
do
  # Split env variables by character `=`
  if echo -n $line | grep -q -e '='; then
    varname=$(echo $line | sed -e 's/=.*//' | tr -d '\n' | tr -d '\r')
    varvalue=$(echo $line | sed -e 's/^[^=]*=//' | tr -d '\n' | tr -d '\r')
  fi

  # Read value of current variable if exists as Environment variable
  value=$(echo ${!varname} | tr -d '\n' | tr -d '\r')
  # Otherwise use value from .env file
  [[ -z $value ]] && value=$(echo $varvalue | tr -d '\n' | tr -d '\r')
  
  # Append configuration property to JS file
  echo -n "  $varname: \"$value\"," >> ./env-config.js
done < $file_name

echo "}" >> ./env-config.js