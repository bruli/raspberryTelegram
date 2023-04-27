#!/usr/bin/env bash

make decryptVault
ansible-playbook -i devops/ansible/inventories/production/hosts devops/ansible/deploy.yml --vault-id raspberry_telegram@devops/ansible/password
make encryptVault
