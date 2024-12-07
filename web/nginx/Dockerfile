FROM nginx:alpine 

COPY conf.d /etc/nginx/conf.d/
# COPY ./certs/datadrivenpath.io/public_html /public_html

#CERTS FOR apppathway.com
COPY ./certs/apppathway.com/dhparam /etc/nginx/dhparam/apppathwaycom
COPY ./certs/apppathway.com/certbot/conf/ /etc/nginx/ssl/apppathwaycom/
COPY ./certs/apppathway.com/certbot/data /usr/share/nginx/html/letsencrypt/apppathwaycom/

#CERTS FOR code.apppathway.com
COPY ./certs/code.apppathway.com/dhparam /etc/nginx/dhparam/code.apppathwaycom
COPY ./certs/code.apppathway.com/certbot/conf/ /etc/nginx/ssl/code.apppathwaycom/
COPY ./certs/code.apppathway.com/certbot/data /usr/share/nginx/html/letsencrypt/code.apppathwaycom/

#CERTS FOR api.apppathway.com
COPY ./certs/api.apppathway.com/dhparam /etc/nginx/dhparam/api.apppathwaycom
COPY ./certs/api.apppathway.com/certbot/conf/ /etc/nginx/ssl/api.apppathwaycom/
COPY ./certs/api.apppathway.com/certbot/data /usr/share/nginx/html/letsencrypt/api.apppathwaycom/

#CERTS FOR next.apppathway.com
COPY ./certs/next.apppathway.com/dhparam /etc/nginx/dhparam/next.apppathwaycom
COPY ./certs/next.apppathway.com/certbot/conf/ /etc/nginx/ssl/next.apppathwaycom/
COPY ./certs/next.apppathway.com/certbot/data /usr/share/nginx/html/letsencrypt/next.apppathwaycom/

#CERTS FOR nhershey.com
COPY ./certs/nhershey.com/dhparam /etc/nginx/dhparam/nhersheycom
COPY ./certs/nhershey.com/certbot/conf/ /etc/nginx/ssl/nhersheycom/
COPY ./certs/nhershey.com/certbot/data /usr/share/nginx/html/letsencrypt/nhersheycom/