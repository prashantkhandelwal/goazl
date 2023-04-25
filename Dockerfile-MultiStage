# syntax=docker/dockerfile:1 

FROM golang:1.20 AS build-stage 

# Set destination for COPY 
WORKDIR /app 

# Download Go modules 
COPY go.mod go.sum ./ 
RUN go mod download 

COPY . ./ 

# Build 
RUN go build -o /goazl 

FROM gcr.io/distroless/base-debian11 AS release-stage 

WORKDIR /app 

COPY --from=build-stage /goazl /goazl 

EXPOSE 8989 

# Run 
ENTRYPOINT ["/goazl"]