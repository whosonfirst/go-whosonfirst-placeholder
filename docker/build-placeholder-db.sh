#!/bin/sh

DATABASE=/data/placeholder/store.sqlite3

ITERATOR_URI=org:///tmp?_exclude_alt=true&exclude=properties.edtf:deprecated=.*&dedupe=true

/usr/local/bin/wof-extract-properties -iterator-uri "${ITERATOR_URI}" $@ > ${DATABASE}

cd /code/pelias/placeholder
npm run build

# copy /data/placeholder/store.sqlite3 to ... ??
