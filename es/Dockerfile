FROM tutum/elasticsearch
RUN sed -i '/listen \[::\]:9200/s/\(listen \[::\]:9200.*\)/#\1/' /etc/nginx/sites-enabled/default
