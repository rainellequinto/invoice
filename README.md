[//]: # (SPDX-License-Identifier: CC-BY-4.0)

## Hyperledger Fabric Samples

Please visit the [installation instructions](http://hyperledger-fabric.readthedocs.io/en/latest/install.html)
to ensure you have the correct prerequisites installed. Please use the
version of the documentation that matches the version of the software you
intend to use to ensure alignment.

## Download Binaries and Docker Images

The [`scripts/bootstrap.sh`](https://github.com/hyperledger/fabric-samples/blob/release-1.3/scripts/bootstrap.sh)
script will preload all of the requisite docker
images for Hyperledger Fabric and tag them with the 'latest' tag. Optionally,
specify a version for fabric, fabric-ca and thirdparty images. Default versions
are 1.4.0, 1.4.0 and 0.4.14 respectively.

```bash
./scripts/bootstrap.sh [version] [ca version] [thirdparty_version]
```

## License <a name="license"></a>

Hyperledger Project source code files are made available under the Apache
License, Version 2.0 (Apache-2.0), located in the [LICENSE](LICENSE) file.
Hyperledger Project documentation files are made available under the Creative
Commons Attribution 4.0 International License (CC-BY-4.0), available at http://creativecommons.org/licenses/by/4.0/.
# hyperledger-training-lab
# invoice

Hyperledger Training
Use Case: Invoice

Installed all the of required development. It is found in the https://hyperledger.github.io/composer/latest/installing/installing-prereqs.html.

Add installation in golang:
1. Go to https://golang.org/dl/ for the compatibility of the device using.
2. After downloading the tar file, extract it in the terminal and paste this command: sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
3. Next, add a directory path environment in .profile. Open the terminal and type nano ~/.profile, then paste this command and save: export PATH=$PATH:/usr/local/go/bin;
4. Lastly, type this command to refresh the .profile file: source ~/.profile.

Also install a postman for the testing part. If you are using the ubuntu, just downloaded it in the ubuntu software.

Download or clone this repository.
Clone also the repository of https://github.com/hyperledger/fabric-samples.
Step 1. Open the terminal and type: nvm use 8. This command is for the compatibility of node in hyperledger.

Step 2. Go to the folder that you clone earlier and go to the directory of invoice.
> hyperledger@ubuntu:~$ cd fabric-samples/invoice

Step 3. Start the fabric server.
> hyperledger@ubuntu:~/fabric-samples/invoice$./startFabric.sh
//This procedure is the starting of the chaincode.
//If there is an error, just type this command:
> $docker kill $(docker ps -q)
> $docker rm $(docker ps -aq)
> $docker rmi $(docker images dev-* -q)
//This command is for stopping and deleting the docker images and container that containing the fabric chaincode

Step 4. Next, type this command in the terminal:
> node enrollAdmin.js
//This is the enrolling of admin in the hyperledger. It will give a certificate that is needed.

Step 5. Next, add a user.
> node registerUser.js
Guide of users:
IBM - Supplier
Lotus - OEM
UBP - Bank
//This is the adding or register the three user. The supplier, bank and the OEM or the Object Efficiency Management.

Step 6. Start the app or the api.
> node app.js

Step 7. Open the postman and put the method in Get and type the localhost:3000. Click the body. Add a username and Supplier or any user that we register earlier in the key and value respectively.
//Supplier, Bank or OEM are the registered user.
//This will give the initiate record in the function initLedger in the invoice.go file. It is located in the invoice folder by the chaincode folder.

Step 8. Adding a invoice.
> Change the GET method into POST and type the localhost:3000/invoice.
> Click the body and add the these key and values:
key			value
username		Lotus
invoiceId		INV1 //stands for invoice1
invoiceNumber		001
billedTo		Bank
invoiceDate		02/12/19
invoiceAmount		5000
itemDescription		Case //Description of the item
gr			No //GR is Good Received
isPaid			No
paidAmount		0
repaid			No
repaymentAmount		0
> Click send
//If you put Bank or OEM in the username value, it will have a error because only the supplier can add a invoice.
For viewing the new invoice: just do the step 7.

Step 9. Bank paying the Supplier
> Change the POST method into PUT and dont change the localhost:3000/invoice.
> In postman there is a checkbox for disabling the key values. Disable all the key except the invoiceId, username and paidAmount.
> Change the username value of UBP, invoiceId is INV1 and put a any value in the paidAmount. Paid Amount value should be lowered in the invoiceAmount. Since we put the invoiceAmount in 5000, you should input a value for paidAmount less than 5000.
> Click Send
//There is also an error if you put Supplier or OEM in the username value. Only the bank should pay the supplier.
For viewing the new invoice: just do the step 7.

Step 10. OEM paying the Bank
> Don't change anything in the method and localhost.
> In postman there is a checkbox for disabling the key values. Disable all the key except the invoiceId, username and repayAmount. Make sure that the paidAmount is disable.
> Change the username value of OEM, invoiceId is INV1 and put a any value in the repaymentAmount. Repayment Amount value should be greater than the paidAmount.
> Click Send
//There is also an error if you put Supplier or Bank in the username value. Only the bank should pay the supplier.
For viewing the new invoice: just do the step 7.

Step 11. Goods Received.
> Don't change anything in the method and localhost.
> Disabled the fields except for the username, invoiceId and gr.
> Change the values into Supplier, INV1, Y respectively.
> Click Send
//There is also an error if you put Supplier or Bank in the username value. Only the bank should pay the supplier.
For viewing the new invoice: just do the step 7.

Step 12. Transaction Log
> Change the method into GET and localhost:3000
> Disable all the key except the username and invoiceId.
> Put the respectively value: UBP and INV1.
//Any username may do.
> Click send.
//You will see the transaction that we did.
//Any username will do.
