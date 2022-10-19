package contracts

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.10.25 --abi=IdRegistry.json --pkg=contracts --type=IdRegistry --out=IdRegistry.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.10.25 --abi=NameRegistry.json --pkg=contracts --type=NameRegistry --out=NameRegistry.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen@v1.10.25 --abi=BundleRegistry.json --pkg=contracts --type=BundleRegistry --out=BundleRegistry.go
