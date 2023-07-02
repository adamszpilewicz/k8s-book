build_push:
	docker build -t 871201/app-status:v1 ./app
	docker login -u 871201 -p Kub16ala!1
	docker push 871201/app-status:v1

	docker build -t 871201/app-checker:v1 ./checker
	docker login -u 871201 -p Kub16ala!1
	docker push 871201/app-checker:v1

deploy:
	kubectl apply -f deploy.yml  -f service.yml

delete:
	kubectl delete -f deploy.yml  -f service.yml

get_pods:
	kubectl get pods

get_logs:
	kubectl logs -n k8s-book -f $(kubectl get pods -n k8s-book -l app=go-app -o jsonpath="{.items[0].metadata.name}") go-checker

api_resources:
	kubectl api-resources

namespaces:
	kubectl get namespaces

describe_default:
	kubectl describe namespace default

list_svc_kube_system:
	kubectl get svc --namespace kube-system

list_svc_default:
	kubectl get svc --namespace default

list_all_default:
	kubectl get all --namespace default

create_ns_k8s_book:
	kubectl create namespace k8s-book

list_namespaces:
	kubectl get namespaces

describe_deployment:
	kubectl describe deploy go-app-deployment -n k8s-book

describe_rs:
	kubectl get rs -n k8s-book

describe_rs_number:
	kubectl describe rs go-app-deployment-748bc895c4 -n k8s-book

describe_rs_number_current:
	kubectl describe rs go-app-deployment-7c4889b4b6 -n k8s-book

get_rollout_history:
	kubectl rollout history deployment go-app-deployment -n k8s-book

undo_rollout:
	kubectl rollout undo deployment go-app-deployment -n k8s-book --to-revision=3
