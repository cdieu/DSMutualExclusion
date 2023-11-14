# DSMutualExclusion

Run one line for one terminal:
go run node/node.go --id=1 --hastoken=true --ownport=5454 --holderport=5555
go run node/node.go --id=2 --hastoken=false --ownport=5555 --holderport=5656
go run node/node.go --id=3 --hastoken=false --ownport=5656 --holderport=5454
