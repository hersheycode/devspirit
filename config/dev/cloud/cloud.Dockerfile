FROM devspace:latest 
RUN sudo usermod --shell /bin/zsh nate
COPY zshrc /home/nate/.zshrc

RUN curl -o- https://deb.packages.mattermost.com/repo-setup.sh | sudo bash
RUN sudo apt install mattermost-omnibus

COPY /home/nate/.ssh /home/nate/.ssh
ENTRYPOINT [ "zsh" ]