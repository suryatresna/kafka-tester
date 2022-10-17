#!/bin/sh

kubectl exec -ti redpanda-0 -c redpanda -- rpk --brokers=redpanda-0.redpanda.default.svc.cluster.local.:9093 topic create onTestConsumer -c cleanup.policy=delete -r 3 -p 1