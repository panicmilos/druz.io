import { Container, inject, injectable } from "inversify";
import DocumentStore from "ravendb";
import { RAVENDB_DATABASE, RAVENDB_URL } from "../config";
import { IInstaller } from "../contracts/IInstaller";


@injectable()
export class DocumentStoreInstaller implements IInstaller {

  constructor(@inject("AppContainer") private appContainer: Container) { }

  install(): void {
    const store = new DocumentStore(RAVENDB_URL, RAVENDB_DATABASE);
    store.initialize();
    this.appContainer.bind<DocumentStore>("DocumentStore").toConstantValue(store);
  }
}
