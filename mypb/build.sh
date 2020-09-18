#!/bin/bash

#### grpc
protoc -I . *.proto --go_out=plugins=grpc:.