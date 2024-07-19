PID      = /tmp/awesome-golang-project.pid
GO_FILES = $(wildcard *.go)
APP      = ./app

serve: restart
	@fswatch -o . | xargs -n1 -I{} make restart

kill:
	@kill `cat $(PID)` || true

before:
	@echo "TODO: inject your hooks here"

$(APP): $(GO_FILES)
	@go build -o $@ $?

restart: kill before $(APP)
	@$(APP) & echo $$! > $(PID)

.PHONY: serve restart kill before
