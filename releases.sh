#!/usr/bin/env bash


if [ -d releases ]; then rm -rf releases/; fi

echo -e "clearing current builds."
if [ -e takegur ]
    then
    rm takegur;
fi

if [ -e takegur.exe ]
    then
    rm takegur.exe;
fi

echo -e "starting linux/osx builds."
for GOOS in darwin linux freebsd openbsd netbsd; do
    for GOARCH in 386 amd64; do
        echo "building $GOOS-$GOARCH"
        env GOOS=$GOOS GOARCH=$GOARCH go build -o releases/takegur takegur.go
        cd releases
        zip -9 takegur-$GOOS-$GOARCH.zip takegur
        rm takegur
        cd ..
    done
done

echo -e "starting windows builds."
for GOOS in windows; do
    for GOARCH in 386 amd64; do
        echo "building $GOOS-$GOARCH"
        env GOOS=$GOOS GOARCH=$GOARCH go build -o releases/takegur.exe takegur.go
        cd releases
        zip -9 takegur-$GOOS-$GOARCH.zip takegur.exe
        rm takegur.exe
        cd ..
    done
done


echo "done."
exit $?