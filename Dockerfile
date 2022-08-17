# syntax=docker/dockerfile:1
FROM golang:1.18-buster


# using command docker build -f Dockerfile -t proj:backend . to run 

ENV app ~/src/startups/justcorpz/backend
RUN mkdir -p "app"
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
#COPY go.sum ./
#COPY * ./

#COPY graph/generated ./
RUN go mod download

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

COPY *.go ./

# Copy SSH key for git private repos
RUN mkdir -p /root/.ssh
ADD /Users/square8/.ssh/id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

# Use git with SSH instead of https
RUN echo “[url \”git@github.com:\”]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
# Skip Host verification for git
RUN echo “StrictHostKeyChecking no “ > /root/.ssh/config

RUN go get github.com/howstrongiam/backend/graph/generated
RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
