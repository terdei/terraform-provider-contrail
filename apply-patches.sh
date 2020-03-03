#!/bin/sh

cd "$(dirname "$0")/../"
TopDir=$(pwd)

for p in $(find ./terraform-provider-contrail/patches/ -type f -name '*.patch' | sort -n); do
	cd "$TopDir"
	target="$(basename "$(dirname "$p")")"
	echo "Found patch file: $p FOR $target"
	cd "$target" || continue
	git apply "$TopDir/$p"
done