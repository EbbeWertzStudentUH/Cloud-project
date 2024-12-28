FROM ruby:latest

WORKDIR /app

COPY . .

RUN bundle install

EXPOSE 3003

ENTRYPOINT ["/app/docker_entrypoint.sh"]
