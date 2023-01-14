FROM mcr.microsoft.com/windows/nanoserver:ltsc2022
LABEL title="Windows 20H2 build"
# install go from scoop since im lazy
# though install scoop first tbh
ARG target_repo
ENV target_repo=$target_repo
RUN iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
RUN scoop install go git
RUN git clone $env:target_repo .
# build it
CMD go build -o main.exe