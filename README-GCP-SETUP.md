## GCP SETUP

### Bucket

**1. Create Service Account:**

    gcloud iam service-accounts create [project]-api --display-name="[project]-api" --project [project]-gcp
    gcloud projects add-iam-policy-binding [project]-gcp --member="serviceAccount:[service_account_email]" --role="roles/storage.admin"
    gcloud iam service-accounts keys create --iam-account=[service_account_email] gcp.json

**2. Convert gcp.json to base64:**

    cat gcp.json | base64

**3. Create Bucket:**

    1. Go to Cloud Storage
    2. Create a new bucket
    3. Set bucket name (e.g. [project]-api-stag)
    4. Choose Location Type (e.g. Multi-region)
    5. Choose Standard default storage class
    6. Choose Uniform access control
    7. Uncheck the box "Enforce public access prevention on this bucket"

**4. Set Bucket Permissions:**

    1. Go to "PERMISSIONS" tab
    2. Click "GRANT ACCESS"
    3. Add "allUsers" in New principals
    4. Select role "Storage Object Viewer"
    5. Click "SAVE"

### Database

**1. Delete default VPC network**

**2. Create new VPC network by running:**

    make gcp-auth
    make gcp-create-vpc

**3. Create PostgreSQL:**

    1. Set Instance ID
    2. Set secure password
    3. Enable both Private IP and Public IP
    4. Select the newly created VPC (vpc-v1) for Private IP

postgres://postgres:[PASSWORD]@[IP_ADDRESS]:5432/postgres?sslmode=disable

### GKE

**1. Create Static IP Address:**

    gcloud compute addresses create api-stag-ip --global --project [project]-gcp

**2. Subdomain**

    Create a subdomain with an A record to the IP

**3. OpenTofu**

    make gcp-auth-terraform
    tofu init
    tofu plan
    tofu apply

**4. GKE**

    make gcp-auth
    make gcp-auth-gke

**5. Argo CD**

    https://argo-cd.readthedocs.io/en/stable/

    make argocd-create
    make argocd-password
    make argocd-port

    kubectl create secret docker-registry container-secret --docker-server=ghcr.io --docker-username=[username] --docker-email=[email] --docker-password=[password] -n [namespace]
