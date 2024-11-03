GOOS=linux GOARCH=amd64 go build  -o index main.go 

scp -i ~/Downloads/MA.pem -r index  ubuntu@ec2-52-23-145-149.compute-1.amazonaws.com:~/deployment/ma-back/

# go build  -o index main.go 
# gcloud compute scp --recurse index tech@crm-temp:~/deployment/flowpod-server-testing --zone "us-central1-a"
# gcloud compute ssh --zone "us-central1-a" "crm-temp" -- 'sudo -S  supervisorctl stop ma-server '
# gcloud compute ssh --zone "us-central1-a" "crm-temp" -- 'sudo cp /home/tech/deployment/flowpod-server-testing/index /home/tech/deployment/admin-server/ '
# gcloud compute ssh --zone "us-central1-a" "crm-temp" -- 'sudo -S  supervisorctl restart ma-server  '


# scp -i ~/Documents/cert/suresh_aws_project.pem -r ~/Documents/sample.sql ubuntu@3.144.23.39:/home/ubuntu/deployment/


# scp -i ~/Documents/cert/suresh_aws_project.pem -r index ubuntu@3.144.23.39:/home/ubuntu/deployment/server/


# ssh -i ~/Downloads/MA.pem ubuntu@ec2-52-23-145-149.compute-1.amazonaws.com 