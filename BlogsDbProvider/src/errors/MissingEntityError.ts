export class MissingEntityError extends Error {

  constructor(message: string) {
    super(message);
    this.name = "MissingEntityError";
  }

}