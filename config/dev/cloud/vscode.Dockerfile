# FROM deployments_ap as ap
FROM codercom/code-server
RUN sudo apt-get update && sudo apt-get upgrade -y && sudo apt-get install wget
RUN sudo usermod --shell /bin/zsh coder
# COPY --from=ap /workspaces/devspirit/pkg/cli/data/tokens.csv /usr/ap/tokens.csv
# COPY --from=ap /etc/ssl/certs/ca.pem /etc/ssl/certs/ca.pem
# COPY --from=ap /app/ap /usr/bin/ap

# ENV AUTH_ADDRESS=auth:50051
# ENV NODES_ADDRESS=nodes:50058
# ENV CA_FILE=/etc/ssl/certs/ca.pem
# ENV AUTH_TOKEN_FILE=/usr/ap/tokens.csv

#Current stable release of Docker: 
RUN sudo curl -fsSL https://get.docker.com -o get-docker.sh
RUN sudo chmod +x get-docker.sh && ./get-docker.sh

# Current stable release of Docker Compose:
RUN sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN sudo chmod +x /usr/local/bin/docker-compose

# get zsh
RUN sudo sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)"

COPY scripts/cloud/zshrc /home/coder/.zshrc 
RUN sudo chown -h coder:coder /home/coder/.zshrc 

SHELL ["/bin/zsh", "-c"]

ENTRYPOINT ["/usr/bin/entrypoint.sh", "--bind-addr", "0.0.0.0:8080", "."]