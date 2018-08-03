#!/bin/bash

Dir="./.thrift/thrift-business"

thrift -out ./thrift -r --gen go $Dir/thrift/reformationActivity.thrift

sed -i '' 's/oprot.Flush()/oprot.Flush(ctx)/g' ./thrift/**/*.go