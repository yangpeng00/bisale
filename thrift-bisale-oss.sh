#!/bin/bash

Dir="./.thrift/thrift-bisale-oss"
if [ -d $Dir ];
then
    cd $Dir
    git add .
    git stash
    git checkout balance-account-20180725
    git pull
    cd ../../
else
    git clone git@git.bisale.org:backend/thrift-bisale-oss.git $Dir
fi

thrift -out ./thrift -r --gen go $Dir/thrift/finance/depositWithdraw.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/user/userKyc.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/user/user.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/content/content.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/engine/orders.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/system/system.thrift

thrift -out ./thrift -r --gen go $Dir/thrift/balanceAccount/balanceAccount.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/balanceAccount/transferrequest.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/content/appVersion.thrift

thrift -out ./thrift -r --gen go $Dir/thrift/engine/accountTransfers.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/engine/tradeDetails.thrift
thrift -out ./thrift -r --gen go $Dir/thrift/finance/accountStatement.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go