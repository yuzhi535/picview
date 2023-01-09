FROM golang:latest

WORKDIR /workspace

# Copy the source from the current directory to the Working Directory inside the container
COPY ./ /workspace/

ENV GOPROXY=https://goproxy.io,direct

# install dependencies
RUN go mod download

ENV GIN_MODE=release

CMD [ "go", "run", "main.go" ]



