# Stage 1: Install Node.js with nvm
FROM debian:bullseye-slim
#
# this file mirrors the build params used in the GitHub Actions and enables
# reproducible builds for downstream forks for Ziti contributors 
#

ARG TARGETARCH
ARG golang_version=1.21.3
ARG go_distribution_file=go${golang_version}.linux-${TARGETARCH}.tar.gz
ARG go_path=/usr/share/go
ARG go_root=/usr/local/go
ARG go_cache=/usr/share/go_cache
ARG uid=1000
ARG gid=1000
RUN apt-get -y update \
    && apt-get -y install \
        gcc-arm-linux-gnueabihf g++-arm-linux-gnueabihf gcc-aarch64-linux-gnu \
        wget build-essential

RUN wget -q https://go.dev/dl/${go_distribution_file}
RUN tar -xzf ${go_distribution_file} -C /usr/local/

RUN wget -qO- https://deb.nodesource.com/setup_18.x | bash \
    && apt-get -y update \
    && apt-get -y install \
        nodejs

RUN mkdir ${go_path} ${go_cache}
RUN chown -R ${uid}:${gid} ${go_path} ${go_cache}

COPY ./linux-build.sh /usr/local/bin/

USER ${uid}:${gid}
ENV TARGETARCH=${TARGETARCH}
ENV GOPATH=${go_path}
ENV GOROOT=${go_root}
ENV GOCACHE=${go_cache}
ENV PATH=${go_path}/bin:${go_root}/bin:$PATH

RUN go install github.com/mitchellh/gox@latest
WORKDIR /mnt
ENTRYPOINT ["linux-build.sh"]
