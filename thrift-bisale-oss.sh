#!/bin/bash

Dir="./.thrift/thrift-bisale-oss"
if [ -d $Dir ];
then
    cd $Dir
    git checkout dev
    git pull
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-bisale-oss.git $Dir
fi

thrift -out ./thrift -r --gen go $Dir/thrift/finance/depositWithdraw.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/user/userKyc.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/user/user.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go