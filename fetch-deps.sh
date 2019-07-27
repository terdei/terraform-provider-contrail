#!/bin/sh

cd "$(dirname "$0")/../"
TopDir=$(pwd)
Repos="https://github.com/Juniper/contrail-generateDS https://github.com/Juniper/contrail-controller"

for repo in $Repos; do
	Dir=${repo##*/}
	echo "$Dir -> $repo"
	[ -d "$Dir" ] && echo "Skipping fetching $repo - directory $Dir already exists" && continue
	git clone "$repo"
done
for p in $(find ./terraform-provider-contrail/patches/ -type f -name '*.patch' | sort -n); do
	cd "$TopDir"
	target="$(basename "$(dirname "$p")")"
	echo "Found patch file: $p FOR $target"
	cd "$target" || continue
	git apply "$TopDir/$p"
done