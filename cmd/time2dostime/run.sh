#!/bin/sh

dostime=$( ./time2dostime '2019-05-01T00:00:00.0+09:00' )

echo dostime: ${dostime}

test -x ~/dostime2time || exec sh -c 'echo dostime2time missing.; exit 1'

printf 'time:    '
~/dostime2time ${dostime}
