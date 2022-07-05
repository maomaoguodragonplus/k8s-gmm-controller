FROM alpine:3.14
WORKDIR /app
COPY ./run.sh ./run.sh
COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json
COPY ./src ./src
EXPOSE 3000
CMD ["sh","run.sh"]
