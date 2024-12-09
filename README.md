# App Pathway

## App Pathway is an app builder.

Start from scratch on codespace:

## Setup docker
### login to docker
> docker login -u datadrivenpath
> docker pull datadrivenpath/devspace:latest
> docker swarm init
> sudo docker network create --driver=overlay --attachable app_pathway_network

### Create New Dev Container: 
> sudo docker run --name=devspace         -v /etc/timezone:/etc/timezone:ro         -v /etc/localtime:/etc/localtime:ro         -v /var/run/docker.sock:/var/run/docker.sock         -v /workspaces/devspirit:/workspaces/devspirit         -v /var/lib/docker:/var/lib/docker         -v /workspaces/devspirit/config/dev/cloud/init.vim:/workspaces/devspirit/config/dev/cloud/init.vim          --network=app_pathway_network          --rm -it datadrivenpath/devspace:latest

## Put these in ~/.bashrc 

echo "alias dev='echo "entering devspace" && sudo docker exec -it devspace zsh && echo "exited devspace"'" >> ~/.bashrc

git tok: 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

> Replace  => /home/nate/code/app-pathway With  => /workspaces/devspirit


> sudo docker-compose up -f third_party.yml code
rm -r ./../*.gitignore

git clone https://hersheycode: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@github.com/hersheyapps/client.git

git clone https://hersheycode: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@github.com/hersheyapps/db_api.git
git clone https://github.com/hersheycode/devspirit.git