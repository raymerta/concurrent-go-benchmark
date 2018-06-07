#!/bin/bash

go build serial-matmul.go
./serial-matmul

go build parallel-matmul.go
./parallel-matmul


