PROJECT_NAME=tag
BIN_DIR=bin

${BIN_DIR}/${PROJECT_NAME}: ${PROJECT_NAME}.go
	go build -o ${BIN_DIR}/${PROJECT_NAME} ${PROJECT_NAME}.go
	cp ${BIN_DIR}/${PROJECT_NAME} .

.PHONY: clean
clean:
	@rm -rf ${BIN_DIR}
	@rm -f ${PROJECT_NAME}
	@echo "cleaned"

# create a simlink to /usr/local/bin
.PHONY: install
install: ${BIN_DIR}/${PROJECT_NAME}
	@sudo ln -sf ${PWD}/${BIN_DIR}/${PROJECT_NAME} /usr/local/bin/${PROJECT_NAME}
	@echo "installed ${PROJECT_NAME} to /usr/local/bin/${PROJECT_NAME}"

.PHONY: uninstall
uninstall:
	@sudo rm -f /usr/local/bin/${PROJECT_NAME}
	@echo "uninstalled ${PROJECT_NAME} from /usr/local/bin/${PROJECT_NAME}"