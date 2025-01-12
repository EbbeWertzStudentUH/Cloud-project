FROM python:3.13.1 AS base

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt

FROM python:3.13.1 AS source

WORKDIR /app
COPY --from=base /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=base /usr/local/bin /usr/local/binCOPY
COPY . .

EXPOSE 3008

ENTRYPOINT ["/app/docker_entrypoint.sh"]
