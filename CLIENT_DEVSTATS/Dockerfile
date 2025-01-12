FROM php:7.4-apache

RUN apt-get update && apt-get install -y libxml2-dev \
    && docker-php-ext-install soap

RUN a2enmod rewrite

WORKDIR /var/www/html

COPY . /var/www/html/proxy/soap

EXPOSE 80
