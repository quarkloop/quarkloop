#!/usr/bin/bash

mkdir -p ./temp

for file in ./models/*.prisma; do
    sed '/\/\/@relations/,$d' "$file" > "./temp/$(basename "$file")"
done

cat ./root.prisma ./temp/*.prisma > ./schema.prisma

rm -rf ./temp