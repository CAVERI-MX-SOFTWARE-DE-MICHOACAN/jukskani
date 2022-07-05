config:
	cp data\environ.example.json data\environ.json
install:
	@echo "Building..."
	go build .
	cp .env.example .env
service:
	cp jukskani.service /etc/systemd/system/jukskani.service
	sudo systemctl daemon-reload
	sudo systemctl enable jukskani
	sudo systemctl start jukskani