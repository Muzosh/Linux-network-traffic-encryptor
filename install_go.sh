#!/bin/sh

wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
rm -r go1.23.1.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc
