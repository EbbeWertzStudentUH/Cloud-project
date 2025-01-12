FROM gradle:7.6.2-jdk17

WORKDIR /app

# java wil niet werken in een goede docker container
# dus ik doe het gewoon op de dirty manier
# (kopieer alle build files)
# (en run 'gradle run')
COPY . .

RUN gradle build --no-daemon 

EXPOSE 3007

ENTRYPOINT ["/app/docker_entrypoint.sh"]
