curl -v -X POST http://localhost:5001/accounts | jq

curl -v -X POST http://localhost:5001/accounts/f9ad3c24-9149-4445-b67c-dafd68f65d67/cashin \
--data-raw '{"amount": 100, "externalId": "000000001"}' | jq

curl -v -X POST http://localhost:5001/accounts/f9ad3c24-9149-4445-b67c-dafd68f65d67/cashout \
--data-raw '{"amount": 100, "externalId": "000000017"}' | jq



#!/bin/bash

TXID=$(uuidgen)

curl -v -X POST http://localhost:5001/accounts/f9ad3c24-9149-4445-b67c-dafd68f65d67/cashin \
--data-raw "{\"amount\": 1, \"externalId\": \"${TXID}\"}" >> /tmp/cashin/${TXID}.json


#!/bin/bash

for i in {1..1000}
do
  echo "Cashin number ${i}"
  ./cashin.sh &
done

echo "Finish"

