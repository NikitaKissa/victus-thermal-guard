MODULE_NAME := github.com/NikitaKissa/victus-thermal-guard

BINARY    := thermal-guard
SERVICE   := thermal-guard.service
CONFIG    := default-config.cfg
BUILD_DIR := build
BINDIR    := /usr/local/bin
UNIT_DIR  := /etc/systemd/system

CONFIG_PATH := /etc/thermal-guard.cfg

.PHONY: build test install uninstall

run:
	@go run main.go

test:
	go vet ./...
	go test -race ./...

build:
	CGO_ENABLED=0 go build --ldflags "-X main.ConfigPath=$(CONFIG_PATH)" -trimpath -o $(BUILD_DIR)/$(BINARY) .

# full install command 'make build && sudo make install && rm -rf build' 
install:
	@test -f $(BUILD_DIR)/$(BINARY) || { echo "binary not found: run 'make' first, then 'sudo make install'"; exit 1; }
	install -Dm755 $(BUILD_DIR)/$(BINARY) $(BINDIR)/$(BINARY)
	install -Dm644 deploy/$(SERVICE) $(UNIT_DIR)/$(SERVICE)
	install -Dm644 deploy/$(CONFIG) $(CONFIG_PATH)
	systemctl daemon-reload
	systemctl enable $(SERVICE)
	systemctl restart $(SERVICE)

uninstall:
	-systemctl disable --now $(SERVICE)
	rm -f $(UNIT_DIR)/$(SERVICE) $(BINDIR)/$(BINARY)
	systemctl daemon-reload