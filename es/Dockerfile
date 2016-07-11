FROM elasticsearch:2.3.4
RUN sed -i '/listen \[::\]:9200/s/\(listen \[::\]:9200.*\)/#\1/' /etc/nginx/sites-enabled/default
