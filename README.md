### Golang Graphql Layer for fleetster Bookings API

This project is mostly a exploration of possibilities with GQLGen and GraphQL in Go. 
It is not a production ready service. The models to not expose all of the fields you would expect since I don't really needed at this point.

#### Build
Run ... 
```shell script
sh scripts/build
```
#### Start GraphQL Server
Run ... 
```shell script
./bookingql -server TARGET_FEETSTER_SERVER
```

#### Notes
I am still not sure what the best way is to handle the dataloading. The loaders seem to work fine but it does cause an issue for the companies by id endpoint which uses the repo instead of the loaders (API specific requirement)
