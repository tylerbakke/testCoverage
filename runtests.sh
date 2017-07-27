#!/bin/bash
echo 'mode: count' > coverage.txt

for pkg in $(go list ./...);
do
    go test -v -covermode=count -coverprofile="profile.tmp" ${pkg}

    if [[ -f "profile.tmp" ]]
    then
        cat "profile.tmp" | tail -n +2 >> coverage.txt
        rm "profile.tmp"
    fi
done

#bash <(curl -s https://codecov.io/bash)
