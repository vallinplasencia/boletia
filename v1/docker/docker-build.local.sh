go mod vendor
# export GITHUB_TOKEN="token-github-access-repos-privados-golang"
docker build --tag vallinplasencia/boletia:dev-xxx -f dockerfile.local ../.
rm -r ../vendor