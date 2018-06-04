#!/usr/bin/env bash

Dir="./.thrift/thrift-business"
if [ -d $Dir ];
then
    cd $Dir
    git checkout .
    git pull
else
    git clone git@git.bisale.org:backend/thrift-business.git $Dir
    cd $Dir
    git checkout dev
fi

cd ../../
thrift -out ./thrift -r --gen go:package_prefix=bisale/bisale-console-api/thrift/ $Dir/thrift/reformationActivity.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go