RUN = ""
ifeq ($(OS),Windows_NT)
	RUN += ""
else
	UNAME = $(shell uname -s)
	ifeq ($(UNAME), Linux)
		RUN += "apt-get install libgtk-3-dev gcc"
	endif
	ifeq ($(UNAME), Darwin)
		RUN += "brew install gtk+3"
endif

install_deps:
	$(RUN)
build:
	go build -ldflags -H=windowsgui -tags hint -o bin/osu_downloader.exe . 