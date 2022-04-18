source deploy-config.sh # load env
docker-compose -p boletia -f docker-compose.yml up -d

echo "Wait 17 seconds db service really ready"
sleep 17s

docker restart boletia_sboletia_1