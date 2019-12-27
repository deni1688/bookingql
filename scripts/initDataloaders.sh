#!/bin/bash

cd loaders;
rm ./*_gen.go
go run github.com/vektah/dataloaden UserLoader string '*github.com/deni1688/bookingql/models.User';
go run github.com/vektah/dataloaden CompanyLoader string '*github.com/deni1688/bookingql/models.Company';
go run github.com/vektah/dataloaden VehicleLoader string '*github.com/deni1688/bookingql/models.Vehicle';
go run github.com/vektah/dataloaden LocationLoader string '*github.com/deni1688/bookingql/models.Location';