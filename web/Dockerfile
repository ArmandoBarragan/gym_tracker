# First layer: create binaries
FROM golang as build-go

WORKDIR /gym_tracker

COPY go.mod go.sum /gym_tracker/

RUN go mod download
COPY . .

EXPOSE 3000
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/gym_tracker .

# # Second layer: run binaries
# FROM alpine:latest

# RUN addgroup -S gym_tracker_group && adduser -S gym_tracker_user -G gym_tracker_group
# USER gym_tracker_user

# WORKDIR /gym_tracker
# COPY --from=build-go /bin/gym_tracker .
# RUN ls -la /gym_tracker  

# EXPOSE 3000
# ENTRYPOINT [ "/gym_tracker/gym_tracker" ]
