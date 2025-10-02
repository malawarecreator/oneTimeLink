# oneTimeLink
A microservice for generating one-time links<br>
## to run
`docker run --env MONGODB_URI=YOUR_URI DB_NAME=YOUR_DB_NAME COLLECTION_NAME=YOUR_COLLECTION_NAME -p YOUR_PORT:8080 onetimelink:VERSION`

## APIs
POST `/createLink?redirectTo=YOUR_URL` to create new link. The `id` will be returned in json `{"id":"THE_ID"}`<br>

POST `/deleteLink?linkId=LINK_ID` to delete the link. A message will be returned in json, view the `/deleteLink` route in `main.go` to see the messages<br>