#Create a client certificate

rm -rf client.crt client.csr client.key

openssl genpkey -algorithm RSA -out client.key -aes256
openssl req -new -key client.key -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365
