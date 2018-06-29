#!/usr/bin/env bash

thrift -out ./thrift -r --gen go:package_prefix=bisale/bisale-console-api/thrift/ ./Wallet.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go