cd cmd/supercute

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

cd ../..

mkdir -p output/cmd/supercute

cp cmd/supercute/supercute output/cmd/supercute

cp -rf web output

tar -zcvf supercute.tar.gz output/