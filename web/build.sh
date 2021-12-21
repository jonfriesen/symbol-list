#!/bin/bash

# This file copies in the latest daily symbol lists and is the build command used on App Platform.

echo VITE_DATA_DATE=$(ls -tr ../data | grep -E '[0-9]\.json$' | tail -n 1 | sed 's/\.[^.]*$//') > .env

cp ../data/$(ls -tr ../data | grep -E '[0-9]\.json$' | tail -n 1) static/data/daily-symbol-list.json && \
 cp ../data/$(ls -tr ../data | grep -E '[0-9]\.csv$' | tail -n 1) static/data/daily-symbol-list.csv && \
 cp ../data/$(ls -tr ../data | grep -E '[0-9]\-crypto\.json$' | tail -n 1) static/data/daily-crypto-symbol-list.json && \
 cp ../data/$(ls -tr ../data | grep -E '[0-9]\-crypto\.csv$' | tail -n 1) static/data/daily-crypto-symbol-list.csv && \
 npm run build