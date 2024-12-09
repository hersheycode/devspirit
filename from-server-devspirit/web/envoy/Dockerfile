FROM envoyproxy/envoy:v1.14.3
COPY web/envoy/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml