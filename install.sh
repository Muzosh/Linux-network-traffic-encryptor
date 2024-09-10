#!/bin/sh

Help()
{
   echo "Usage: ./install.sh [Other gateway network]"
   echo "Network format: x.x.x.x/y"
   echo
}

if [ $# -lt 1 ] || [ $1 = "--help" ] || [ $1 = "-h" ]
then
Help
exit 1
fi

Route_IP=$1

sudo ip tuntap add name tun99 mode tun
sudo ip link set tun99 up
sudo ip addr add 192.168.1.1 peer 192.168.1.2 dev tun99
sudo ip route add $Route_IP via 192.168.1.2

echo "1" | sudo tee /proc/sys/net/ipv4/ip_forward

go build .