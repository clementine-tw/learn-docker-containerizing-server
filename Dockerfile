FROM debian:stable-slim

# COPY source destination
COPY server /bin/server

ENV PORT=8991

CMD ["/bin/server"]

