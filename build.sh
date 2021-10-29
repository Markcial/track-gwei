rm -Rf dist
mkdir dist
# create data icon for systray
rm icon/main.go
echo "//+build linux darwin" >> icon/main.go
cat images/gas_4062.png | 2goarray Data icon >> icon/main.go

env GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-arch x86_64" CGO_ENABLED=1 go build -o dist/gwei .
cd dist/ && appify -icon ../images/gas_4062.png -name "GweiTracker" ./gwei
 