#!/bin/bash

# This script creates the configuration for a Lachesis testnet with a variable
# number of nodes. It will generate crytographic key pairs and assemble a
# peers.json file in the format used by Lachesis. The files are copied into
# individual folders for each node so that these folders can be used as the
# datadir that Lachesis reads configuration from.

set -e

N=${1:-4}
DEST=${2:-"$PWD/conf"}
IPBASE=${3:-172.77.5.}
PORT=${4:-1337}

for i in $(seq 1 $N)
do
    dest=$DEST/node$i
    mkdir -p $dest
    echo "Generating key pair for node$i"
    docker run --rm \
        -v $dest:/keys \
        go-lachesis \
        keygen --pem /keys/priv_key.pem --pub /keys/key.pub
    echo "$IPBASE$i:$PORT" > $dest/addr
done

PFILE=$DEST/peers.json
echo "[" > $PFILE
for i in $(seq 1 $N)
do
    com=","
    if [[ $i == $N ]]; then
        com=""
    fi

    printf "\t{\n" >> $PFILE
    printf "\t\t\"NetAddr\":\"$(cat $DEST/node$i/addr)\",\n" >> $PFILE
    printf "\t\t\"PubKeyHex\":\"$(cat $DEST/node$i/key.pub)\"\n" >> $PFILE
    printf "\t}%s\n"  $com >> $PFILE

done
echo "]" >> $PFILE

for i in $(seq 1 $N)
do
    dest=$DEST/node$i
    cp $DEST/peers.json $dest/
done
