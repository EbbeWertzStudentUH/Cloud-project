FROM ruby:latest

WORKDIR /app

COPY . .

RUN bundle install

EXPOSE 4567

ENTRYPOINT ["/app/docker_entrypoint.sh"]
