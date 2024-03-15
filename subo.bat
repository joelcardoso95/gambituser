git add .
git commit -m "Last Commit"
git push
set GOOS=linux
set GOARCH=arm64
go build main.go
del main.zip
tar.exe -a -cf main.zip bootstrap