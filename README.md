# oneTimeLink
A microservice for generating one-time links<br>
## To run
```shell 
docker run \
  -e MONGODB_URI=YOUR_URI \
  -e DB_NAME=YOUR_DB_NAME \
  -e COLLECTION_NAME=YOUR_COLLECTION_NAME \
  -e PORT=8080 \
  -p 8080:8080 \
  onetimelink:VERSION
```

## APIs
POST `/createLink?redirectTo=YOUR_URL` to create new link. The `id` will be returned in json `{"id":"THE_ID"}`<br>

POST `/deleteLink?linkId=LINK_ID` to delete the link. A message will be returned in json, view the `/deleteLink` route in `main.go` to see the messages<br>