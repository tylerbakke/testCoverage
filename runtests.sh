#!/bin/bash
echo 'mode: count' > coverage.txt

for pkg in $(go list ./...);
do
    dir="$GOPATH/src/$pkg"
    len="${#PWD}"
    dir_relative=".${dir:$len}"

    go test -v -covermode=count -coverprofile="$dir_relative/profile.tmp" "$dir_relative"

    if [[ -f "$dir_relative/profile.tmp" ]]
    then
        cat "$dir_relative/profile.tmp" | tail -n +2 >> coverage.txt
        rm "$dir_relative/profile.tmp"
    fi
done

#bash <(curl -s https://codecov.io/bash)
