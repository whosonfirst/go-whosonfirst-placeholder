#!/bin/sh

DATABASE=/data/placeholder/store.sqlite3
TARGET=
ACCESS_TOKEN=

ITERATOR_URI=org:///tmp?_exclude_alt=true&exclude=properties.edtf:deprecated=.*&dedupe=true

while getopts "A:T:" opt; do
    # echo "-$opt = $OPTARG"
    case "$opt" in
	A)
	    ACCCESS_TOKEN=$OPTARGS
	    ;;
	T)
	    TARGET=$OPTARG
	    ;;
	: )
	    echo "WHAT"
	    ;;
    esac
done

shift $((OPTIND-1))
SOURCES=$@

/usr/local/bin/wof-extract-properties -iterator-uri "${ITERATOR_URI}" -access-token "${ACCESS_TOKEN}" ${SOURCES} > ${DATABASE}
# echo /usr/local/bin/wof-extract-properties -iterator-uri "${ITERATOR_URI}" -access-token "${ACCESS_TOKEN}" ${SOURCES} 

cd /code/pelias/placeholder
npm run build

/usr/local/bin/copy file://${DATABASE} ${TARGET}
