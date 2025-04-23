#!/bin/sh

PLACEHOLDER="/code/pelias/placeholder"
DATA="${PLACEHOLDER}/data"

EXTRACT="${DATA}/wof.extract"
DATABASE="${DATA}/store.sqlite3"

TARGET=

ITERATOR_URI="org:///tmp?_dedupe=true&_exclude_alt=true&exclude=properties.edtf:deprecated=.*"

while getopts "T:" opt; do
    # echo "-$opt = $OPTARG"
    case "$opt" in
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

echo /usr/local/bin/wof-extract-properties -verbose -iterator-uri "${ITERATOR_URI}" ${SOURCES} ">" ${EXTRACT}
/usr/local/bin/wof-extract-properties -verbose -iterator-uri "${ITERATOR_URI}" ${SOURCES} > ${EXTRACT}

if [ "$?" -ne 0 ]; then
    echo "wof-extract-properties exited with non-zero status"
    exit 1;
fi

echo "build database"
cd /code/pelias/placeholder
npm run build

if [ "$?" -ne 0 ]; then
    echo "build command exited with non-zero status"
    exit 1;
fi

echo /usr/local/bin/copy -source file://${DATABASE} -target ${TARGET}
/usr/local/bin/copy -source file://${DATABASE} -target ${TARGET}

if [ "$?" -ne 0 ]; then
    echo "copy exited with non-zero status"
    exit 1;
fi

exit 0;
