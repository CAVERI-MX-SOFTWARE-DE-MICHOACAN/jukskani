config:
	cp data\environ.example.json data\environ.json
install:
	echo "Building..."
	go build .
	cp jukskani.service /etc/systemd/system/jukskani.service
service:
	sudo systemctl daemon-reload
	sudo systemctl enable jukskani
	sudo systemctl start jukskani