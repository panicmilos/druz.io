cd ../EmailService
docker build -t druzio-email-service .

cd ../ImageService
docker build -t druzio-image-service .

cd ../PostsDbProvider
docker build -t druzio-posts-db-provider .

cd ../PostsService
docker build -t druzio-posts-service .