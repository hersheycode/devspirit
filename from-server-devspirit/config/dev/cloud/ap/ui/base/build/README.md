
# Opened deployments_vscode_dev in vscode
# 2)  cd code-server && yarn upgrade @typescript-eslint/eslint-plugin@4.31.0
# 3) yarn upgrade @typescript-eslint/parser@4.31.0
# 4) yarn release
# 5) yarn release:standalone
# 6) yarn add minimist
# RUN curl -L https://github.com/a8m/envsubst/releases/download/v1.2.0/envsubst-`uname -s`-`uname -m` -o envsubst
# RUN chmod +x envsubst
# RUN mv envsubst /usr/local/bin
# yarn package
# left and commited changes: d commit 3b75c7a570ac deployments_vscode_dev
FROM deployments_vscode_dev as build 

# FROM ubuntu:latest
# ENV DEBIAN_FRONTEND=noninteractive
# RUN apt-get update && apt-get upgrade -y 
# RUN apt-get install -y \ 
#     apt-utils \
#     sudo \
#     openssh-server \
#     openssh-client \
#     apt-transport-https \
#     ca-certificates \
#     curl \
#     wget \
#     gnupg \
#     lsb-release \
#     git \
#     neovim \
#     openssh-client \
#     dumb-init \
#     zsh \
#     htop \
#     locales \
#     man \
#     nano \
#     git-lfs \
#     procps \
#     openssh-client \
#     vim.tiny \
#     lsb-release \
#   && git lfs install \
#   && rm -rf /var/lib/apt/lists/*

FROM debian:11

RUN apt-get update \
 && apt-get install -y \
    curl \
    dumb-init \
    zsh \
    htop \
    locales \
    man \
    nano \
    git \
    git-lfs \
    procps \
    openssh-client \
    sudo \
    vim.tiny \
    lsb-release \
  && git lfs install \
  && rm -rf /var/lib/apt/lists/*

# https://wiki.debian.org/Locale#Manually
RUN sed -i "s/# en_US.UTF-8/en_US.UTF-8/" /etc/locale.gen \
  && locale-gen

ENV LANG=en_US.UTF-8

RUN adduser --gecos '' --disabled-password coder && \
  echo "coder ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/nopasswd
# RUN groupadd --gid $USER_GID $USERNAME \
#     && useradd -p "$(openssl passwd -1 covsdffdvflvfrdealjksvnkjefdbbSxbd2bdfb42)" --uid $USER_UID --gid $USER_GID -m $USERNAME \
#     && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
#     && chmod 0440 /etc/sudoers.d/$USERNAME

RUN ARCH="$(dpkg --print-architecture)" && \
    curl -fsSL "https://github.com/boxboat/fixuid/releases/download/v0.5/fixuid-0.5-linux-$ARCH.tar.gz" | tar -C /usr/local/bin -xzf - && \
    chown root:root /usr/local/bin/fixuid && \
    chmod 4755 /usr/local/bin/fixuid && \
    mkdir -p /etc/fixuid && \
    printf "user: coder\ngroup: coder\n" > /etc/fixuid/config.yml

COPY --from=build /app/code-server /usr/lib/code-server
COPY --from=build /app/code-server/release-packages ./release-packages
RUN mv ./release-packages/code-server*.deb /tmp/

COPY third_party/code-server/ci/release-image/entrypoint.sh /usr/bin/entrypoint.sh
RUN dpkg -i /tmp/code-server*$(dpkg --print-architecture).deb && rm /tmp/code-server*.deb

USER 1000
ENV USER=coder
WORKDIR /home/coder
ENTRYPOINT ["/usr/bin/entrypoint.sh", "--bind-addr", "0.0.0.0:8080", "."]