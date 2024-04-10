readonly service="$1"

cd "./internal/$service"
ls
find .env
env $(cat "./.env") go run .
