FROM scratch

MAINTAINER @Jeadie

COPY datediff_docker main

ENV GIN_MODE=release
EXPOSE $PORT
ENTRYPOINT ["/main", "server", "-p", "$PORT"]
