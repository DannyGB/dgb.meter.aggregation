FROM golang:1.19.1 AS build

RUN mkdir -p /app/agg

COPY . /app/agg

WORKDIR /app/agg/cmd

RUN go build -o meter.aggregation .

FROM golang:1.19.1 AS final

RUN mkdir -p /app/agg/cmd
WORKDIR /app/agg/cmd

COPY --from=build /app/agg/cmd/meter.aggregation .

ARG ENVIRONMENT=dev
ENV AGGREGATION_READINGS_ENV $ENVIRONMENT

ARG HTTP_PORT=80
ENV AGGREGATION_HTTP_PORT $HTTP_PORT

ARG ELEC_BACKEND=
ENV AGGREGATION_ELEC_BACKEND $ELEC_BACKEND

ARG OIL_BACKEND=
ENV AGGREGATION_OIL_BACKEND $OIL_BACKEND

ARG CORS_CLIENTS="*"
ENV AGGREGATION_CORS_CLIENTS $CORS_CLIENTS

CMD [ "/app/agg/cmd/meter.aggregation" ]
