#!/usr/bin/env fish
if test -e rice-box.go
    rm rice-box.go
end

rice embed-go

cd core

if test -e rice-box.go
    rm rice-box.go
end

rice embed-go

cd ..

cd cmd
go build -v -o ../build/ogwc