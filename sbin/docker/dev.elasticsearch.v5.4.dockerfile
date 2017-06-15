# AWS now supports v5.4, so we stick with that... not that it matters to this challenge but I like to keep my planets aligned.
FROM elasticsearch:5.4
# official ES images will be deprecated soon, so need to change to the one below (which comes with X-Pack)
# FROM docker.elastic.co/elasticsearch/elasticsearch:5.2.2

ADD files/elasticsearch.v5.4.yml /usr/share/elasticsearch/config/

# ... because we love all world languages
RUN /usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-icu

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["elasticsearch"]
