#!/bin/bash

Dir="./.thrift/thrift-business"
if [ -d $Dir ];
then
    cd $Dir
    git checkout .
    git pull
    git checkout dev
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-business.git $Dir
    git checkout dev
fi

thrift -out ./thrift -r --gen go $Dir/thrift/reformationActivity.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go