#!/bin/bash

TXID=$(uuidgen)

curl -s -X POST http://localhost:5001/accounts/f9ad3c24-9149-4445-b67c-dafd68f65d67/cashout \
--data-raw "{\"amount\": 1, \"externalId\": \"${TXID}\"}" >> /tmp/cashin/${1}-${TXID}.json

