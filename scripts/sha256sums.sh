#!/bin/bash

SHA256SUMS="Binary Assets Checksum: \n"
cd output || exit
for FILE in *
do
    SHA256SUMS+="$(sha256sum "${FILE}")"
    SHA256SUMS+="\n"
done
cd ..
echo -e "${SHA256SUMS}" | tee output/sha256sums.txt
