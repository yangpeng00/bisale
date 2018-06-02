#!/usr/bin/env bash

Dir="./.thrift/thrift-storage"
if [ -d $Dir ];
then
    cd $Dir
    git checkout .
    git pull
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-storage.git $Dir
fi

thrift -out ./thrift -r --gen go $Dir/thrift/storageService.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go