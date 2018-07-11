all:
	go build -buildmode=plugin -o plugins/plugin.so plugins/plugin.go
	go build app.go

clean:
	rm -f app
