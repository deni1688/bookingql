### Golang Graphql Layer for fleetster Bookings API

This project is mostly a exploration of possibilities with GQLGen and GraphQL in Go. 
It is not a production ready service. The models do not expose all of the fields you would expect since I don't really needed at this point.


#### Build
You will need to have GO installed to build the executable. Run ... 
```bash
sh scripts/build.sh
```
#### Start GraphQL Server
Comes with GraphQL Playground. Pass auth token through headers. Run ... 
```bash
./bookingql -server TARGET_FLEETSTER_SERVER
```

#### Notes
I am still not sure what the best way is to handle the dataloading. The loaders seem to work fine but it does cause an issue for the companies by id endpoint which uses the repo instead of the loaders (API specific requirement)