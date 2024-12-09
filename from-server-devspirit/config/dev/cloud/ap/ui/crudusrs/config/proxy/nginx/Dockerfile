FROM nginx:latest

RUN apt-get update && apt-get upgrade -y && apt-get install -y inotify-tools
RUN mkdir -p /var/local/step
# COPY site.conf /etc/nginx/conf.d/
COPY conf.d/default.conf /etc/nginx/conf.d/
COPY certwatch.sh /
COPY entrypoint.sh /

# Certificate watcher and nginx
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
