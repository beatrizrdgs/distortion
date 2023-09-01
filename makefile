up:
	docker compose up -d
dev-hot-reload:
	nodemon --exec go run . --signal SIGTERM