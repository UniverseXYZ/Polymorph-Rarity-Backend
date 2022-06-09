# Polymorphs Rarity Backend (Burn To Mint edition)
## Summary
- This is the latest and most up-to-date version of the rarity backend. It is updated for the `Burn to Mint` Polymorphs edition
- Everytime new polymorph(s) is/are burned from V1 contract and minted to V2, all records from the Mongo V1 collections are deleted
- Also adds some concurrency and database optimizations
## Deployment
- Currently, the backend runs as a process in a GCloud Virtual Machine:
  - `go run main.go`
- Notes:
  - When deploying for `production`, take into account all constants in `constants/metadata.go`
  - This backend only queries the contract ~ 15 seconds for specific events, and updates the Mongo DB accordingly
  - The mongo collections themselves are queried using the following gcloud function: `https://github.com/UniverseXYZ/Polymorph-Rarity-Cloud`