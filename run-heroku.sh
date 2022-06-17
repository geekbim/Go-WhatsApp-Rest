go build -o bin/go_wa_rest -v cmd/main.go
git add .
git commit -m "initial commit"
git push origin master -f
git push heroku master:master