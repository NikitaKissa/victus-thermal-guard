BINARY    := thermal-guard
SERVICE   := thermal-guard.service
BUILD_DIR := build
BINDIR    := /usr/local/bin
UNIT_DIR  := /etc/systemd/system

.PHONY: build test install uninstall

run:
	@go run main.go

test:
	go vet ./...
	go test -race ./...

build:
	CGO_ENABLED=0 go build -trimpath -o $(BUILD_DIR)/$(BINARY) .

# full install command 'make build && sudo make install && rm -rf build' 
install:
	@test -f $(BUILD_DIR)/$(BINARY) || { echo "binary not found: run 'make' first, then 'sudo make install'"; exit 1; }
	install -Dm755 $(BUILD_DIR)/$(BINARY) $(BINDIR)/$(BINARY)
	install -Dm644 deploy/$(SERVICE) $(UNIT_DIR)/$(SERVICE)
	systemctl daemon-reload
	systemctl enable $(SERVICE)
	systemctl restart $(SERVICE)

uninstall:
	-systemctl disable --now $(SERVICE)
	rm -f $(UNIT_DIR)/$(SERVICE) $(BINDIR)/$(BINARY)
	systemctl daemon-reload