FROM fedora:32

RUN \
    rm -f \
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