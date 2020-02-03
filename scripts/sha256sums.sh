#!/bin/bash

SHA256SUMS=""
cd output || exit
for FILE in *
do
    SHA256SUMS="${SHA256SUMS} \n $(sha256sum ${FILE})"
done
cd ..
echo -e "${SHA256SUMS}" | tee output/sha256sums.txt
