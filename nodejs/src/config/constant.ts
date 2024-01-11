import { config } from 'dotenv';

config();
export const PORT: number | string = 5000 || (process.env.PORT as string);
