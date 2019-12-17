# JCA
This is a cryptographic algorithm written by me.

## Inspiration
I've always wanted to create a cryptographic algorithm. So when I found golang, I said to myself now is the time to do it. So I started learning Go and created my cryptographic algorithm whilst doing so.

## jca
The jca binary file can be used to encrypt and decrypt any file with the given key. To do so, you just need to execute the following commands.
#### Encryption
```sh
jca -in YOUR_INPUT_FILE -key YOUR_KEY_FILE -out YOUR_OUTPUT_FILE
```
If you want to use a custom number of encryption rounds you can do so by using the -rounds switch.
```sh
jca -in YOUR_INPUT_FILE -key YOUR_KEY_FILE -out YOUR_OUTPUT_FILE -rounds YOUR_ROUND_COUNT
```
Please note, that if you change the round count you need to do so for the decryption too.
#### Decryption
```sh
jca -in YOUR_INPUT_FILE -key YOUR_KEY_FILE -out YOUR_OUTPUT_FILE -decrypt
```
If you used a custom number of rounds for the encryption you need to use the -rounds switch.
```sh
jca -in YOUR_INPUT_FILE -key YOUR_KEY_FILE -out YOUR_OUTPUT_FILE -rounds YOUR_ROUND_COUNT -decrypt
```

## jca-keygen
At this point, the jca-keygen supports keys of size 2048 bytes, 4096 bytes, 8192 bytes, 16384 bytes, 32768 bytes, 65536 bytes, 131072 bytes, and 262144 bytes.

To generate a key use:
```sh
jca-keygen -size YOUR_KEY_SIZE -out YOUR_OUT_FILE
```

If you don't specify the size the program is gonna use 8192 bytes by default.

## jca-integrity
This tool generates the sha256 and md5 sums for a given file. 
#### Attention
I might remove this tool in the future and add a decryption check to the cryptographic algorithm, that way it can check if the decryption was successful or not.

## Dependencies
The binaries do not need any libraries. If you want to build the tool by yourself tho you need to install golang.
### Ubuntu
```sh
sudo apt-get install golang
```
### Windows
Go to the golang.org site and download the newest version.

## Information for the interested
For my algorithm, I used the substitution box from AES, due to me being too lazy to make my own. 

## Issues
If you have any issues please specify examples like exit code and console output, no support if not given!!!

## Collaborators
 - Jooarye