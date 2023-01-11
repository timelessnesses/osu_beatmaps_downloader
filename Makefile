.PHONY: install_deps

RUN = 
ifeq ($(OS),Windows_NT)
	RUN += 
else
	UNAME = $(shell uname -s)
	ifeq ($(UNAME), Linux)
		RUN += sudo apt-get install libgtk-3-dev gcc
	endif
	ifeq ($(UNAME), Darwin)
		RUN += brew install gtk+3
	endif
endif

install_deps:
	$(RUN)
build:
	go install github.com/fyne-io/fyne-cross@latest
	fyne-cross windows -arch=*
