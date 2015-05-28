#
# nkcraddock/mervis
#
# https://github.com/nkcraddock/mervis
FROM scratch
MAINTAINER Nathan Craddock "nkcraddock@gmail.com"

ADD build/mervis /mervis

VOLUME ["/data"]
WORKDIR /data

EXPOSE 443

ENTRYPOINT ["/mervis"]
