gcloud compute instances create quoter \
    --image container-vm-v20140925 \
    --image-project google-containers \
    --metadata-from-file google-container-manifest=./containers.yaml \
    --tags http-server \
    --zone us-central1-a \
    --machine-type f1-micro