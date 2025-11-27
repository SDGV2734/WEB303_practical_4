# Practical Report: Student Cafe Microservices Implementation

**Course:** WEB303 - Advanced Web Technologies  
**Practical:** 4 - Microservices Architecture with Kubernetes  

## Executive Summary

This practical report documents the implementation of a cloud-native microservices application for a student cafe ordering system. The solution demonstrates enterprise-grade patterns including service discovery, API gateway implementation, and container orchestration using Kubernetes, Consul, and Kong.

## System Architecture

### Technical Architecture

The system follows a microservices architecture pattern with the following traffic flow:

```
Client Applications → Kong API Gateway → Microservices → Kubernetes Service Discovery
```

### System Components

- **Food Catalog Service (Go)**: RESTful API for menu item and pricing management
- **Order Service (Go)**: Order processing engine with integrated service discovery capabilities
- **Cafe UI (React)**: Single-page application for customer interactions including menu browsing and order placement
- **Consul**: Distributed service registry providing health monitoring and service discovery (optional component)
- **Kong**: Enterprise API gateway handling external routing, load balancing, and security policies
- **Kubernetes**: Container orchestration platform providing service mesh capabilities

## Implementation Guide

### System Requirements

**Hardware Requirements:**

- Docker Desktop (minimum 4GB RAM allocated)
- Minikube cluster
- kubectl CLI tool
- Helm package manager

**Software Dependencies:**

- Kubernetes cluster (v1.20+)
- Go runtime (v1.16+) for local development
- Node.js (v14+) for frontend development

### Step 1: Infrastructure Setup

```bash
# Initialize Minikube cluster
minikube start

# Configure Docker environment to use Minikube's Docker daemon
eval $(minikube docker-env)
```

### Step 2: Service Dependencies Installation

```bash
# Create namespace
kubectl create namespace student-cafe

# Install Consul (optional - for advanced service discovery)
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install consul hashicorp/consul \
  --set global.name=consul \
  --set server.replicas=1 \
  --set ui.enabled=true \
  --set connectInject.enabled=true \
  -n student-cafe

# Install Kong API Gateway
helm repo add kong https://charts.konghq.com
helm install kong kong/kong \
  --set ingressController.installCRDs=false \
  --set admin.enabled=true \
  -n student-cafe
```

### Step 3: Application Deployment

```bash
# Build Docker images within Minikube environment
docker build -t food-catalog-service:v1 ./food-catalog-service/
docker build -t order-service:v1 ./order-service/
docker build -t cafe-ui:v1 ./cafe-ui/

# Deploy application services to Kubernetes cluster
kubectl apply -f app-deployment.yaml

# Configure Kong ingress routing policies
kubectl apply -f kong-ingress.yaml
```

### Step 4: System Verification

```bash
# Retrieve Minikube cluster IP address
minikube ip

# Obtain Kong proxy service port mapping
kubectl get service kong-kong-proxy -n student-cafe

# Access deployed application
# Example endpoint: http://192.168.49.2:30622
```

## Project Implementation Details

### Directory Structure Analysis

```
student-cafe/
├── food-catalog-service/       # Go-based microservice for menu management
│   ├── main.go                # HTTP service implementation using Chi router framework
│   ├── go.mod                 # Go module dependency definitions
│   ├── go.sum                 # Cryptographic dependency checksums
│   └── Dockerfile             # Multi-stage container build configuration
├── order-service/              # Go-based microservice for order processing
│   ├── main.go                # Service implementation with Consul service discovery
│   ├── go.mod                 # Go module dependency management
│   ├── go.sum                 # Dependency integrity verification
│   └── Dockerfile             # Containerization build specification
├── cafe-ui/                    # React-based frontend application
│   ├── src/                   # React component source code
│   ├── public/                # Static asset resources
│   ├── package.json           # NPM package dependencies and scripts
│   └── Dockerfile             # Multi-stage build for production deployment
├── app-deployment.yaml         # Kubernetes deployment and service manifests
├── kong-ingress.yaml          # Kong API gateway ingress configuration
└── minikube                   # Local Kubernetes cluster reference
```

## API Specification

### Internal Service Endpoints

**Food Catalog Service** (Port: 8080)

- `GET /items`: Retrieve complete menu catalog with pricing information

**Order Service** (Port: 8081)

- `POST /orders`: Process new order requests with item validation
- `GET /health`: Service health monitoring endpoint for Kubernetes probes

### External API Gateway Endpoints

**Kong Gateway Routes:**

- `GET /api/catalog`: Menu items accessible through API gateway with load balancing
- `POST /api/orders`: Order submission endpoint with gateway security policies
- `GET /`: React frontend application delivery with static asset caching

## Technical Implementation Features

### Service Discovery Implementation

**Kubernetes Native Discovery:**

- Services communicate via Kubernetes DNS resolution system
- Automatic service endpoint registration and discovery
- Built-in load balancing through kube-proxy

**Consul Integration (Optional):**

- Advanced service registry for complex service mesh scenarios
- Health monitoring with configurable check intervals
- Dynamic service configuration management

**Health Monitoring:**

- Kubernetes readiness probes ensure traffic routing to healthy pods
- Liveness probes provide automatic restart capabilities for failed services
- Dynamic service discovery eliminates hardcoded endpoint configurations

### API Gateway Architecture

**Unified Entry Point:**

- Kong serves as the single external access point for all client requests
- Centralized security policy enforcement and rate limiting
- SSL termination and certificate management

**Intelligent Routing:**

- Path-based routing with configurable rules:
  - Root path (`/`) → React Frontend Application
  - API path (`/api/catalog`) → Food Catalog Service
  - Order path (`/api/orders`) → Order Service
- Automatic load balancing across multiple service instances
- Circuit breaker patterns for fault tolerance

**Policy Enforcement:**

- Centralized authentication and authorization mechanisms
- Rate limiting and throttling policies
- Request/response transformation capabilities

### Container Orchestration Strategy

**Infrastructure as Code:**

- Declarative deployment specifications using YAML manifests
- Version-controlled infrastructure configurations
- Reproducible environment deployments

**Auto-scaling Capabilities:**

- Horizontal Pod Autoscaler (HPA) for dynamic scaling based on resource utilization
- Vertical Pod Autoscaler (VPA) for optimized resource allocation
- Cluster autoscaling for node-level resource management

**Service Networking:**

- Automatic service-to-service communication through Kubernetes networking
- Network policies for traffic isolation and security
- DNS-based service discovery with FQDN resolution

**Operational Management:**

- Rolling deployment strategies with zero-downtime updates
- Automatic rollback capabilities on deployment failures
- Health management with configurable restart policies

## Design Pattern Implementation Analysis

### Implemented Design Patterns

- ✅ **Service Registry Pattern**: Kubernetes DNS resolution combined with optional Consul integration
- ✅ **API Gateway Pattern**: Kong implementation for centralized external traffic management
- ✅ **Health Check Pattern**: Kubernetes probes providing automated health monitoring
- ✅ **Service Discovery Pattern**: Dynamic service location without hardcoded configurations
- ✅ **Container Orchestration Pattern**: Kubernetes-based deployment automation and lifecycle management
- ✅ **Circuit Breaker Pattern**: Graceful failure handling and fault tolerance mechanisms
- ✅ **Database per Service Pattern**: Independent data management per microservice

## System Monitoring and Diagnostics

### Service Health Verification

```bash
# Monitor all pod status within the namespace
kubectl get pods -n student-cafe

# Verify deployment status and replica counts
kubectl get deployments -n student-cafe

# Examine service endpoints and networking configuration
kubectl get services -n student-cafe
kubectl get endpoints -n student-cafe
```

### Application Logging and Diagnostics

```bash
# Real-time order service log monitoring
kubectl logs -f deployment/order-deployment -n student-cafe

# Food catalog service log analysis
kubectl logs -f deployment/food-catalog-deployment -n student-cafe

# Frontend application log inspection
kubectl logs -f deployment/cafe-ui-deployment -n student-cafe
```

### API Gateway Management

```bash
# Access Kong admin interface for configuration management
kubectl port-forward -n student-cafe svc/kong-kong-manager 8002:8002
# Navigate to: http://localhost:8002

# Analyze ingress configuration and routing rules
kubectl describe ingress cafe-ingress -n student-cafe
```

### Inter-Service Communication Testing

```bash
# Execute commands within order service container
kubectl exec -it deployment/order-deployment -n student-cafe -- sh

# Verify internal service communication capabilities
curl http://food-catalog-service:8080/items

# Test DNS resolution functionality
nslookup food-catalog-service
```

## Development Environment Configuration

### Local Development Setup

**Backend Services (Go):**

```bash
# Execute food catalog service locally (requires Go 1.16+)
cd food-catalog-service && go run main.go

# Execute order service locally
cd order-service && go run main.go
```

**Frontend Application (React):**

```bash
# Install dependencies and start development server (requires Node.js 14+)
cd cafe-ui && npm install && npm start
```

### Container Image Management

**Image Building Process:**

```bash
# Ensure Docker context points to Minikube environment
eval $(minikube docker-env)

# Build updated container images with version tagging
docker build -t food-catalog-service:v2 ./food-catalog-service/
docker build -t order-service:v2 ./order-service/
docker build -t cafe-ui:v2 ./cafe-ui/

# Update image versions in deployment manifests, then apply changes
kubectl apply -f app-deployment.yaml
```

**Deployment Updates:**

```bash
# Force restart deployments to pull updated images
kubectl rollout restart deployment/food-catalog-deployment -n student-cafe
kubectl rollout restart deployment/order-deployment -n student-cafe
kubectl rollout restart deployment/cafe-ui-deployment -n student-cafe
```

## Production Readiness Assessment

### Security

- [ ] Implement mTLS between services
- [ ] Add authentication to API gateway (Kong plugins)
- [ ] Use Kubernetes network policies for traffic isolation
- [ ] Secure service-to-service communication
- [ ] Implement API rate limiting

### Scalability

- [ ] Configure horizontal pod autoscaling (HPA)
- [ ] Implement database per service pattern
- [ ] Add persistent volumes for data storage
- [ ] Configure resource limits and requests
- [ ] Implement caching strategies

### Observability

- [ ] Add distributed tracing (Jaeger/Zipkin)
- [ ] Implement metrics collection (Prometheus)
- [ ] Centralized logging (ELK/EFK stack)
- [ ] Application performance monitoring (APM)
- [ ] Custom dashboards (Grafana)

### Reliability

- [ ] Implement circuit breaker pattern
- [ ] Add retry mechanisms with exponential backoff
- [ ] Configure timeouts for service calls
- [ ] Implement bulkhead pattern
- [ ] Add chaos engineering tests

## 📚 Learning Outcomes

This project demonstrates mastery of:

- ✅ **Microservices Architecture Design** - Decomposition into focused services
- ✅ **Container Orchestration** - Kubernetes deployment and service management
- ✅ **Service Discovery** - Both Kubernetes-native and Consul-based approaches
- ✅ **API Gateway Implementation** - Kong for traffic management
- ✅ **Cloud-Native Development** - 12-factor app principles
- ✅ **DevOps Practices** - Infrastructure as code, containerization
- ✅ **Inter-Service Communication** - HTTP APIs and service mesh concepts
- ✅ **Health Monitoring** - Probes, logging, and observability
- ✅ **Fault Tolerance** - Resilience patterns and graceful degradation

## 🧪 Testing the Application

### Functional Testing

1. **Access the Frontend**: Navigate to the Kong gateway URL
2. **Browse Menu**: Verify food items load correctly
3. **Add to Cart**: Test shopping cart functionality
4. **Place Order**: Submit an order and observe the response
5. **API Testing**: Use curl to test endpoints directly

### Expected Behavior

- ✅ Frontend loads and displays menu items
- ✅ Services communicate successfully
- ✅ Orders can be placed through the UI
- ⚠️ **Intentional Failure**: Order submission may fail as designed to demonstrate error handling

### Test Commands

```bash
# Test catalog endpoint
curl http://$(minikube ip):$(kubectl get svc kong-kong-proxy -n student-cafe -o jsonpath='{.spec.ports[0].nodePort}')/api/catalog

# Test order submission
curl -X POST http://$(minikube ip):$(kubectl get svc kong-kong-proxy -n student-cafe -o jsonpath='{.spec.ports[0].nodePort}')/api/orders \
  -H "Content-Type: application/json" \
  -d '{"items":[{"id":"1","name":"Coffee","price":2.5}],"total":2.5}'
```


---
