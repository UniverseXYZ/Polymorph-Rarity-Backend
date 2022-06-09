# Polymorphs Rarity Backend (Burn To Mint edition)
## Summary
- This is the latest and most up-to-date version of the rarity backend. It is updated for the `Burn to Mint` Polymorphs edition
- Everytime a new polymorphs is burned from V1 contract and minted to V2, all records from the Mongo V1 collections are deleted
- Also adds some concurrency and database optimizations
## Deployment
- Currently, the backend runs as a process in a GCloud Virtual Machine:
  - `go run main.go`
- Notes:
  - When deploying for `production`, take into account all constants in `constants/metadata.go`