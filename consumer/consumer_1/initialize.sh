#!/bin/sh

kubectl exec -it svc/redpanda -- rpk topic create onTestConsumer -c cleanup.policy=delete -r 1 -p 1