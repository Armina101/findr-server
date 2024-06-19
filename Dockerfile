# user the latest golang image of the standard library
FROM golang:1.21.0

# create and set the working directory for the docker images to build
WORKDIR /app

# copy the go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# download all dependencies and verify them
RUN go mod download && go mod verify

# copy the source from the current directory to the working directory of the container
COPY . .

# build the application from the main package of the source code(./cmd)
RUN CGO_ENABLED=0 go build -o testfindr .

# make the build binary executable by changing the permission mode of the file
RUN chmod +x ./testfindr

CMD [ "/app/testfindr" ]