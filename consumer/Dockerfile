# Get the standard golang base image
FROM golang:1.23

# set the working dir
WORKDIR /app
# set env
ENV GOWORK=/app/go.work
# copy the go.work file
COPY go.work ./
# copy our target runner folder
COPY /job-read ./job-read
# copy the folder job for reference to be able to build by using go workspace
COPY /job-write ./job-write
COPY /consumer ./consumer
COPY /producer ./producer
COPY /shared ./shared
# get the packages and verify
RUN go mod download && go mod verify

# build the application
WORKDIR /app/consumer
RUN go build -v -o .

# remove folders after build
RUN rm -rf ./job-write
RUN rm -rf ./producer
RUN rm -rf ./shared
RUN rm -rf go.work

CMD ["./consumer"]