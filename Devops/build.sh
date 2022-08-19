cd ../EmailService
docker build -t druzio-email-service .

cd ../ImageService
docker build -t druzio-image-service .

cd ../PostsDbProvider
docker build -t druzio-posts-db-provider .

cd ../PostsService
docker build -t druzio-posts-service .

cd ../UserService
docker build -t druzio-user-service .

cd ../UserRelationsService
docker build -t druzio-user-relations-service .

cd ../ChatService
docker build -t druzio-chat-service .

cd ../FrontendService
docker build -t druzio-frontend-service .