FROM python:3.13.1

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 3004
EXPOSE 3005

ENTRYPOINT ["/app/docker_entrypoint.sh"]
