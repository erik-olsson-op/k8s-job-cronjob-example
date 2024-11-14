dockertagRead = job-read:latest
dockertagWrite = job-write:latest
dockerTagConsumer = consumer:latest
docerTagProducer = producer:latest
k8sNamespace = my-namespace
# Default target
start: database-start docker-build minikube-load minikube-apply-resources
delete: minikube-delete-resources

# spin up MariaDB
database-start:
	docker compose -f scripts/docker-compose-db.yml up --force-recreate -d

# Build the docker images
docker-build:
	docker build -f consumer/Dockerfile -t $(dockerTagConsumer) .
	docker build -f producer/Dockerfile -t $(docerTagProducer) .
	docker build -f job-read/Dockerfile -t $(dockertagRead) .
	docker build -f job-write/Dockerfile -t $(dockertagWrite) .

# Load local images to minikube cluster
minikube-load:
	minikube image load $(dockerTagConsumer)
	minikube image load $(docerTagProducer)
	minikube image load $(dockertagRead)
	minikube image load $(dockertagWrite)

# Delete k8s definitions
minikube-delete-resources:
	minikube kubectl -- delete configmap my-config --namespace=$(k8sNamespace)
	minikube kubectl -- delete secret my-secret --namespace=$(k8sNamespace)
	minikube kubectl -- delete job job-read --namespace=$(k8sNamespace)
	minikube kubectl -- delete pod consumer --namespace=$(k8sNamespace)
	minikube kubectl -- delete pod producer --namespace=$(k8sNamespace)
	minikube kubectl -- delete cronjob job-write --namespace=$(k8sNamespace)
	minikube kubectl -- delete svc mariadb-external -n $(k8sNamespace)

minikube-read-job:
	minikube kubectl -- apply -f job-read/k8s.yml

# Apply k8s definitions
minikube-apply-resources:
	minikube kubectl -- apply -f scripts/config-k8s.yml
	minikube kubectl -- apply -f scripts/secret-k8s.yml
	minikube kubectl -- apply -f scripts/service-k8s.yml
	minikube kubectl -- apply -f consumer/k8s.yml
	minikube kubectl -- apply -f producer/k8s.yml
	minikube kubectl -- apply -f job-read/k8s.yml
	minikube kubectl -- apply -f job-write/k8s.yml

# Run locally in Docker
docker-run-write:
	docker run -it --network=scripts_my_bridge_network --rm $(dockertagWrite)
docker-run-read:
	docker run -it --network=scripts_my_bridge_network --rm $(dockertagRead)