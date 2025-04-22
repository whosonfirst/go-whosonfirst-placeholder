#!/bin/sh

DATABASE=/data/placeholder/store.sqlite3

/usr/local/bin/wof-extract-properties -iterator-uri 'org:///tmp?exclude=properties.edtf:deprecated=.*' $@ > ${DATABASE}

cd /code/pelias/placeholder
npm run build

# copy /data/placeholder/store.sqlite3 to ... ??
