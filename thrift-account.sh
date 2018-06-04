#!/usr/bin/env bash

Dir="./.thrift/thrift-account"
if [ -d $Dir ];
then
    cd $Dir
    git checkout .
    git pull
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-account.git $Dir
fi

thrift -out ./thrift -r --gen go $Dir/thrift/accountService.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go