#!/bin/sh

cd ../ && protoc --go_out=./ ./protocol/message.proto
