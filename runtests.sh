#!/bin/bash
echo 'mode: count' > coverage.txt

for pkg in $(go list ./...);
do
    echo $pkg
    dir="$GOPATH/src/$pkg"
    echo $dir
    len="${#PWD}"
    echo $len
    dir_relative=".${dir:$len}"
    echo $dir_relative

    go test -v -covermode=count -coverprofile="$dir_relative/profile.tmp" "$dir_relative"

    if [[ -f "$dir_relative/profile.tmp" ]]
    then
        cat "$dir_relative/profile.tmp" | tail -n +2 >> coverage.txt
        rm "$dir_relative/profile.tmp"
    fi
done

#bash <(curl -s https://codecov.io/bash)
