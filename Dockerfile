
###################################
#Build stage
FROM gitea/gitea:1.6

COPY entrypoint.sh /usr/local/bin
COPY setup /etc/s6/gitea

EXPOSE 22 3000

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
