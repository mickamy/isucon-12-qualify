#!/bin/bash
set -euo pipefail

git fetch origin --prune
git pull origin

cd go && make isuports && cd -

sudo systemctl daemon-reload
sudo systemctl restart mysql.service
sudo systemctl restart nginx.service
sudo systemctl restart isuports.service
