# Deploy with skaffold

## Tools utilisés 
---

* Golang
* Skafold
* Kubernetes (et Docker)
* Metrics Server

## But : 
---

* Déployer une application avec Skafold 
* Tester la scalabilité de mon application en fonction de ressources allouées et de ressources limites *(Dans ce projet la limite d'utilisation est de 70% d'utilisation alloué du programme CPU)* -> avec **hpa.yaml**

## Fichiers 
--- 
* app
    * Dockerfile
        * Créer une image docker
    * server.go
        * Mon programme
* k8s
    * deployment.yaml
        * Configuration de déploiement de l'application *(container, port, variable d'environement, ressources…)*
    * hpa.yaml
        * Configuration du scaling de l'application en fonction du CPU *(ou de la RAM)*
    * service.yaml
        * Point d'entrée *(url)* qui permet d'arriver dans le deployement *(dépend du selector et va chercher notre configuration de deployment)*
* skaffold.yaml
    * Permet de déployer facilement et évite de devoir build manuellement et push l'image docker puis de re-déployé les fichiers kub


## Isntall dependencies

### Helm

```bash
brew install helm
 
helm repo add stable https://charts.helm.sh/stable
helm repo update

```

### metrics server
```bash
helm install metrics-server stable/metrics-server --namespace kube-system --set args={"--kubelet-insecure-tls=true"}
```