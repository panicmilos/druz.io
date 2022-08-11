import { Server } from 'http';
import { IInstaller } from "./contracts/IInstaller";
import { AppContainer } from "./AppContainer";
import { Application, Nimble, NimblyApi } from 'nimbly-api';
import { PostsController } from './controllers/PostsController';
import { MissingEntityError } from './errors/MissingEntityError';

const PORT = 80;

main();

function main() {
  installMiddleware();

  const nimble = new Nimble().addLocalServices(PostsController);
  new NimblyApi({
    app: AppContainer.get<Application>('Application'),
    nimblyInfo: {
      name: "PostsDbProvider"
    }
  })
  .withErrors({
    [MissingEntityError.name]: 404,
    Error: 500,
  })
  .from(nimble);

  startServer();
}

function installMiddleware() {
  const installers: IInstaller[] = AppContainer.getAll<IInstaller>("IInstaller");
  installers.forEach(i => i.install());
}

function startServer(): void {
  const httpServer = AppContainer.get<Server>(Server);
  httpServer.listen(PORT);
  
  console.log(`Server is listening on port ${PORT}.`);
}