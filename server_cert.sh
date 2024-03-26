#Create a server certificate

rm -rf server.crt server.csr server.key

echo "Removed old server certificates"

openssl genpkey -algorithm RSA -out server.key -aes256
openssl req -new -key server.key -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
