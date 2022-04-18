source deploy-config.sh

DST_BACKUP_POSTGRES="$BOLETIA_GITIGNORE/postgres/"
#Delete backup
mkdir -p $DST_BACKUP_POSTGRES
sudo rm -r $DST_BACKUP_POSTGRES
#Copy backup
mkdir -p $DST_BACKUP_POSTGRES
cp -r misc/postgres/* $DST_BACKUP_POSTGRES