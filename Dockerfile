FROM nginx:1-alpine

WORKDIR /app

RUN wget https://github.com/Conflux-Chain/sirius-auto-release/releases/latest/download/scan.zip -O scan.zip && \
    wget https://github.com/Conflux-Chain/sirius-auto-release/releases/latest/download/scan-eth.zip -O scan-eth.zip && \
    unzip scan.zip -d scan && unzip scan-eth.zip -d scan-eth && rm -rf scan.zip scan-eth.zip

COPY ./nginx/* /etc/nginx
COPY ./scan.conf /etc/nginx/conf.d/scan.conf
COPY ./scan_eth.conf /etc/nginx/conf.d/scan_eth.conf



RUN mkdir -p /data/scan && mkdir -p /data/scan-eth


COPY ./entrypoint.sh .

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]