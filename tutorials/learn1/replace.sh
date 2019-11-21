#!/bin/bash
for file in $( find . -name "*.go" | grep -v "_test" )
do
     sed -i "" 's/github\.com\/micro/micro_learn\/micro/g' $file  
done

## mac sed 命令