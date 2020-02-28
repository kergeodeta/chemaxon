
instances:
	@echo "Creating AWS EC2 instances with Terraform"
	@if [ -f "ansible/inventory" ]; then rm -f ansible/inventory; fi
	@terraform apply terraform

setup:
	@echo "Run containerized applications in cloud"
	@ansible-playbook \
		 --private-key=ssh/chemkey \
		 --inventory ansible/inventory \
		 --extra-vars="ansible_agent_password=`openssl rand -base64 32` account_private_key=$LETSENCRYPT_ACC" \
		 --ssh-extra-args="-o StrictHostKeyChecking=no" \
		 ansible/site.yml
