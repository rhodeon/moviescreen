#!/bin/bash
# setup script for Ubuntu 20.04 on AWS EC2

# raise an error and exit the shell if an unset parameter is used
set -eu

# --- VARIABLES ---
TIMEZONE=Africa/Lagos
USERNAME=moviescreen

# prompt for admin username and project database password
read -p "Enter password for moviescreen DB user: " DB_PASSWORD

# force all output to be presented in en_US for the duration of this script
export LC_ALL=en_US.UTF-8

# --- SCRIPT LOGIC ---
# enable the universe repository
sudo add-apt-repository --yes universe

# update all packages
sudo apt update
sudo apt --yes -o Dpkg::Options::="--force-confnew" upgrade

# set timezone and install all locales
sudo timedatectl set-timezone ${TIMEZONE}
sudo apt --yes install locales-all

# create new user with root privileges and ensure password is required on login
sudo useradd --create-home --shell "/bin/bash" --groups sudo "${USERNAME}"
sudo passwd --delete "${USERNAME}"
sudo chage --lastday 0 "${USERNAME}"

# copy SSH keys from the root to the new user
sudo rsync --archive --chown=${USERNAME}:{USERNAME} /home/ubuntu/.ssh /home/${USERNAME}

# --- FIREWALL ---
# configure and enable firewall
sudo ufw allow 22/tcp # SSH
sudo ufw allow 80/tcp # HTTP
sudo ufw allow 443/tcp # HTTPS
sudo ufw --force enable

# install Fail2Ban
sudo apt --yes install fail2ban

# install database migrate tool (https://github.com/golang-migrate/migrate)
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.deb --output migrate.deb
sudo dpkg -i migrate.deb
rm migrate.deb

# --- DATABASE ---
# install PostgreSQL (https://www.postgresql.org/download/linux/ubuntu/)
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" | sudo tee /etc/apt/sources.list.d/pgdg.list
sudo apt update
sudo apt --yes install postgresql

# set up moviescreen database
sudo -i -u postgres psql -c "CREATE DATABASE moviescreen"
sudo -i -u postgres psql -d moviescreen -c "CREATE EXTENSION IF NOT EXISTS citext"
sudo -i -u postgres psql -d moviescreen -c "CREATE ROLE moviescreen WITH LOGIN PASSWORD '${DB_PASSWORD}'"

# add project database to environment variables
echo "MOVIESCREEN_DB_DSN='postgres://moviescreen:${DB_PASSWORD}@localhost/moviescreen'" | sudo tee -a /etc/environment

# install Caddy (https://caddyserver.com/docs/install#debian-ubuntu-raspbian)
sudo apt --yes install debian-keyring debian-archive-keyring apt-transport-https
curl -sL 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo tee /etc/apt/trusted.gpg.d/caddy-stable.asc
curl -sL 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt --yes install caddy

echo "Script complete. Rebooting..."
sudo reboot
