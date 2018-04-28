#
# Makefile
# mantas_marcinkus, 2018-04-28 00:00
#

run-service:
	go run main.go

dev:
	echo '[bash] Kill processes by using CTRL+C'
	trap 'kill %1; kill %2' SIGINT;
	echo '[bash] Running webpack (building typescript react/redux app) AND Starting HTTP go service'
	npm run-script build-watch | sed -e 's/^/[webpack] /' & CompileDaemon -command="./algis" -exclude-dir=".git" -exclude-dir="node_modules" -exclude-dir="src" | sed -e 's/^/[go] /'

deps:
	npm install
	glide install
	go get github.com/githubnemo/CompileDaemon

build:
	mkdir publish | cd publish | mkdir dist service
	go build

