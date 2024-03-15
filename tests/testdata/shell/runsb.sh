#!/bin/bash

COMMAND="/root/.ya/bin/sing-box run -c /root/.ya/subscribe/conf.json"

PIDS=$(pgrep -f "$COMMAND")

if [ -z "$PIDS" ]; then
    kill $PIDS
    sleep 2
fi

nohup $COMMAND > /dev/null 2>&1 &
