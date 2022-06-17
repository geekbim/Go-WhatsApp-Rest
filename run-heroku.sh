# heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go
go build -o bin/go_wa_rest -v cmd/main.go
git add .
git commit -m "initial commit"
git push origin master
git push heroku master:master