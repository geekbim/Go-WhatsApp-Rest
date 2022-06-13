go build -o bin/go_wa_rest -v cmd/main.go
git add .
git commit -m "build(go-rest-wa): add bin file"
git push origin master
git push heroku development:master