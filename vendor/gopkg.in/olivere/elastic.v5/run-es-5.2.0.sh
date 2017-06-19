docker run --rm --privileged=true -p 9200:9200 -p 9300:9300 -v "$PWD/config.json:/usr/share/elasticsearch/config.json" -e ES_JAVA_OPTS='-Xms1g -Xmx1g' elasticsearch:5.2.0 elasticsearch
