FROM golang:1.15-buster AS builder
ENV CGO_ENABLED=0
ARG COMPILE_FLAGS
WORKDIR /root/sentinel-hub
COPY . /root/sentinel-hub
RUN make build

FROM debian:buster AS sentinel-hub
ARG GENESIS_URL
ENV GENESIS_URL=${GENESIS_URL}
ARG GENESIS_CHECKSUM
ENV GENESIS_CHECKSUM=${GENESIS_CHECKSUM}
RUN adduser --uid 1000 --disabled-password --gecos '' --home /srv/sentinel-hub sentinel-hub
RUN apt-get -yq update \
        && DEBIAN_FRONTEND=noninteractive apt-get install -y \
                unattended-upgrades \
                # ssl certs to external services
                ca-certificates \
                curl \
        && rm -rf /var/lib/apt/lists/* \
        && apt-get clean
COPY --from=builder /root/sentinel-hub/bin/sentinel-hub-cli /usr/bin/
COPY --from=builder /root/sentinel-hub/bin/sentinel-hubd /usr/bin/
COPY entrypoint /usr/bin/
USER sentinel-hub
RUN mkdir -p /srv/sentinel-hub/.sentinel-hubd
VOLUME /srv/sentinel-hub/.sentinel-hubd
ENTRYPOINT ["entrypoint"]
CMD ["start"]
