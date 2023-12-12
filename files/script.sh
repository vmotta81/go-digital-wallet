curl -v -X POST http://localhost:5001/accounts | jq

curl -v -X POST http://localhost:5001/accounts/f9ad3c24-9149-4445-b67c-dafd68f65d67/cashin \
--data-raw '{"amount": 100, "externalId": "000000001"}'