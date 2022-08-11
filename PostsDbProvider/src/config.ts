import dotenv from 'dotenv';

dotenv.config();

export const { RAVENDB_URL, RAVENDB_DATABASE, PORT } = process.env as any;
