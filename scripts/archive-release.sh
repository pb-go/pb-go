#!/bin/bash
for FILE in output/*
do
  if [[ $FILE == *"exe"* ]]; then
    CMD="zip -rm ${FILE}.zip ${FILE}"
  else
    CMD="xz ${FILE}"
  fi
  echo "$CMD"
  $CMD
done