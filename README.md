## 💬 Introduction
E-Procurement for Vhiweb Test. The API endpoints is accessible [here](https://www.postman.com/techtoniclabs/workspace/vhiweb-test).

## 🚀 Quickstart
1. Create a postgres instance.
2. Create new file called `.env` and fill exactly like the `.env.example`.
3. Download the go dependencies using `go mod tidy`
4. Run migration to create tables needed using `go run migrations/migrate.go`.
5. Run the server using `go run main.go`