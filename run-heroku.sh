go build -o bin/go-wa-rest -v cmd/main.go
git add .
git commit -m "build(go-rest-wa): add bin file"
git push origin master
git push heroku master:main