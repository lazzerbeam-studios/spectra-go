## NocoDB

### GCE

**1. Create Default VPC Network:**

    gcloud compute networks create default --subnet-mode=auto --bgp-routing-mode=regional --project=[project]-gcp

**2. Create Static IP Address:**

    gcloud compute addresses create nocodb-ip --region [region] --project [project]-gcp
    gcloud compute addresses describe nocodb-ip --region [region] --project [project]-gcp

**3. Create VM:**

    gcloud compute instances create nocodb-vm \
      --project=[project]-gcp \
      --zone=[zone] \
      --machine-type=e2-medium \
      --subnet=default \
      --address=nocodb-ip \
      --image-family=ubuntu-2204-lts \
      --image-project=ubuntu-os-cloud \
      --boot-disk-size=20GB \
      --boot-disk-type=pd-ssd \
      --tags=nocodb

**4. Open Firewall:**

    gcloud compute firewall-rules create allow-nocodb-ssh \
      --project=[project]-gcp \
      --direction=INGRESS \
      --priority=1000 \
      --network=default \
      --action=ALLOW \
      --rules=tcp:22 \
      --target-tags=nocodb \
      --source-ranges=0.0.0.0/0

    gcloud compute firewall-rules create allow-nocodb-web \
      --project=[project]-gcp \
      --direction=INGRESS \
      --priority=1000 \
      --network=default \
      --action=ALLOW \
      --rules=tcp:80,tcp:443 \
      --target-tags=nocodb \
      --source-ranges=0.0.0.0/0

**5. Create Subdomain:**

    nocodb.[your-domain.com] -> nocodb-ip

### VM

**1. SSH Into VM:**

    gcloud compute ssh nocodb-vm --zone=[zone] --project=[project]-gcp

**2. Install Docker:**

    sudo apt update
    sudo apt install -y docker.io
    sudo systemctl enable docker
    sudo systemctl start docker

**3. Install Caddy:**

    sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl
    curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
    curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
    sudo apt update
    sudo apt install -y caddy

**4. Create Data Directory:**

    sudo mkdir -p /opt/nocodb/data
    sudo chown -R $USER:$USER /opt/nocodb

**5. Generate Secret And Run NocoDB:**

    openssl rand -base64 48

    sudo docker run -d \
      --name nocodb \
      --restart unless-stopped \
      -p 127.0.0.1:8080:8080 \
      -e NC_AUTH_JWT_SECRET='[GENERATED_SECRET]' \
      -e NC_PUBLIC_URL='https://nocodb.[your-domain.com]' \
      -e NC_TOOL_DIR='/usr/app/data' \
      -e NC_DISABLE_TELE='true' \
      -v /opt/nocodb/data:/usr/app/data \
      nocodb/nocodb:latest

**6. Create Caddyfile:**

```bash
sudo tee /etc/caddy/Caddyfile >/dev/null <<'EOF'
nocodb.[your-domain.com] {
    reverse_proxy 127.0.0.1:8080
}
EOF
```

**7. Restart Caddy:**

    sudo systemctl restart caddy
    sudo systemctl status caddy

### Verify

**1. Confirm App Opens:**

    https://nocodb.[your-domain.com]

**2. Confirm Container Is Running:**

    sudo docker ps
    sudo docker logs nocodb

**3. Confirm Persistence:**

    sudo docker restart nocodb

### Cleanup

**1. Delete VM:**

    gcloud compute instances delete nocodb-vm --zone=[zone] --project=[project]-gcp

**2. Delete Static IP Address:**

    gcloud compute addresses delete nocodb-ip --region [region] --project [project]-gcp

**3. Delete Firewall Rules:**

    gcloud compute firewall-rules delete allow-nocodb-ssh --project=[project]-gcp
    gcloud compute firewall-rules delete allow-nocodb-web --project=[project]-gcp

**4. Delete Default VPC Network:**

    gcloud compute networks delete default --project=[project]-gcp
