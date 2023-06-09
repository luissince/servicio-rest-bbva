env=".env"

if [ -e "$env" ]; then
    echo "El archivo .env existe"
else
    cp .env.example .env
fi

mkdir logs

docker stop servicio-bbva && docker rm servicio-bbva

docker image rm servicio-bbva

docker build -t servicio-bbva .

docker run -d \
--restart always \
--name servicio-bbva \
--net=upla \
-p 8889:80 \
-v $(pwd)/logs:/etc/push \
servicio-bbva