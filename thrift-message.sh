#!/usr/bin/env bash

Dir="./.thrift/thrift-message"
if [ -d $Dir ];
then
    cd $Dir
    git checkout .
    git pull
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-message.git $Dir
fi

thrift -out ./thrift -r --gen go $Dir/thrift/messageService.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go