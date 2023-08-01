# use a lightweight base image 
FROM golang:1.17-alpine 

# set the working directory inside the container 
WORKDIR /app 

# copy the source code into the container 
COPY . . 

# build the Go binary 
RUN go build -o tasker 

# Use a minimal base image for running the app 
FROM alpine:latest 

# Set the working directory inside the container 
WORKDIR /app  

# Copty the binary from the previous stage 
COPY --from=0 /app/tasker . 

# Set the entry point for the container 
ENTRYPOINT ["./tasker"]