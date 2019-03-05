#!/usr/bin/env bash

mongo -- "$MONGO_DB" <<EOF
var user = '$MONGO_INITDB_ROOT_USERNAME';
var passwd = '$MONGO_INITDB_ROOT_PASSWORD';
var admin = db.getSiblingDB('admin');
admin.auth(user, passwd);
db.createUser({user: "$MONGO_USER", pwd: "$MONGO_PASSWORD", roles: ["readWrite"]});
EOF