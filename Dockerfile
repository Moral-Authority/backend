# syntax=docker/dockerfile:1
FROM golang:1.18-buster AS build


# using command docker build -f Dockerfile -t proj:backend . to run 
ENV app ~/src/startups/justcorpz/backend
RUN mkdir -p "app"
WORKDIR /app

# Download necessary Go modules
COPY go.mod go.sum ./

#COPY graph/generated ./
RUN go mod download

RUN GO111MODULE=on 
ENV GOFLAGS=-mod=mod


COPY *.go ./

ARG SSH_PRIVATE_KEY
RUN mkdir ~/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_ed25519
RUN chmod 600 ~/.ssh/id_ed25519
RUN ssh-keyscan github.com >> ~/.ssh/known_hosts# Print SSH_PRIVATE_KEY 

# Skip Host verification for git
RUN echo “StrictHostKeyChecking no “ > /root/.ssh/config

# RUN go get github.com/howstrongiam/backend/graph/generated


EXPOSE 8080
