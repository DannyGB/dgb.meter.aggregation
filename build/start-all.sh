#!/bin/bash
nohup ./agg/cmd/meter.aggregation &
nohup ./elec/cmd/meter.readings/meter.readings &
nohup ./oil/cmd/meter.readings.oil/meter.readings.oil &
nohup ./notifications/cmd/meter.notifications/meter.notifications &
nohup ./bot/cmd/meter.readings.bot/meter.readings.bot
