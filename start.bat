mkdir backend-project
cd backend-project
mkdir cmd internal internal\adapters internal\adapters\auth internal\adapters\profile internal\adapters\events internal\adapters\dashboard internal\adapters\providers internal\adapters\repository internal\domain internal\domain\entities internal\domain\services internal\infrastructure internal\infrastructure\database internal\infrastructure\logging internal\infrastructure\config internal\usecases internal\usecases\ports internal\usecases\interactor internal\api internal\api\routes internal\api\middlewares pkg pkg\framework
cd cmd && type nul > main.go && cd ..
type nul > go.mod
type nul > README.md