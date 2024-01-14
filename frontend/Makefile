init:
	cd web/static && npm install

build: init
	cd web/static && npm run build
	go build -o tmp/main cmd/webapp/main.go
