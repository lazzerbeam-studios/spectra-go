## NocoDB

**1. Create Static IP Address:**

    gcloud compute addresses create nocodb-ip --global --project [project]-gcp

**2. Create Subdomain:**

    nocodb.[domain] -> nocodb-ip

**3. Create Metadata Database:**

    Create a small PostgreSQL database for NocoDB metadata only.

**4. Create Secret in `api-stag` :**

Generate a JWT secret first:

    openssl rand -base64 48

    kubectl create secret generic nocodb-secret \
      --from-literal=NC_AUTH_JWT_SECRET='[secure-jwt-secret]' \
      --from-literal=NC_DB='pg://[HOST]:5432?u=postgres&p=[PASSWORD]&d=postgres' \
      --namespace api-stag

**5. Create ArgoCD App:**

    Repository: [github_repo]
    Path: nocodb
    Namespace: api-stag
