FROM fedora:32
MAINTAINER lburgazzoli

ARG CONTAINER_USER
ARG CONTAINER_USER_GID
ARG CONTAINER_USER_UID

ENV CONTAINER_USER_HOME  /home/$CONTAINER_USER
ENV HOME                 $CONTAINER_USER_HOME
RUN \
    groupadd \
        --gid $CONTAINER_USER_GID \
        $CONTAINER_USER \
    && useradd \
        --create-home \
        --home-dir $CONTAINER_USER_HOME \
        --uid $CONTAINER_USER_GID \
        --gid $CONTAINER_USER_UID \
        $CONTAINER_USER \
    && rm -f \
        /etc/yum.repos.d/fedora-cisco-openh264.repo \
        /etc/yum.repos.d/fedora-updates-testing.repo \
    && dnf --assumeyes --refresh --verbose --allowerasing update \
    && dnf --assumeyes --verbose --allowerasing install \
        dnf-plugins-core \
        curl \
        unzip \
        wget \
        tar \
        bzip2 \
        xz \
    && dnf --assumeyes --verbose clean all

