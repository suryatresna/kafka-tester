load('ext://helm_remote', 'helm_remote')

helm_remote(
    'redpanda',
    repo_url='https://charts.redpanda.com',
    allow_duplicates=True
)

include("consumer/consumer_1/Tiltfile")

include("producer/producer_1/Tiltfile")


k8s_yaml([
    # 'resource/redpanda.yaml',
    'resource/redpanda-console.yaml',
],  allow_duplicates=True)

k8s_resource(
    workload='redpanda:statefulset',
    port_forwards=[
        port_forward(8081,8081),
        port_forward(8082,8082),
        port_forward(9092,9092),
        port_forward(29092,29092),
    ],
    labels=['resource']
)

k8s_resource(
    workload='redpanda-console',
    port_forwards=[
        port_forward(8080, 8080)
    ],
    resource_deps=[
        "redpanda"
    ],
    labels=['monitor']
)
