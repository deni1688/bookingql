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

#### DataLoaders
The project was initially using DataLoaders (still present in dev branch) but I have removed them in favor of data hydration on the API side. 
