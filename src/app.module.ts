import { Logger, Module } from '@nestjs/common';
import {
  ClientProviderOptions,
  ClientsModule,
  Transport,
} from '@nestjs/microservices';
import * as dotenv from 'dotenv';
import { readFileSync } from 'fs';
import { AppController } from './app.controller';
import {
  KAFKA_CLIENT_TOKEN
} from './app.service';

const envFile = process.env.NODE_ENV ? `.env.${process.env.NODE_END}` : '.env';
dotenv.config({ path: envFile });

const { name: appName } = JSON.parse(
  readFileSync('package.json', {
    encoding: 'utf-8',
  }),
);

@Module({
  imports: [ClientsModule.register([getMicroserviceClientSettings()])],
  controllers: [AppController],
  providers: [Logger],
})
export class AppModule {}

function getMicroserviceClientSettings(): ClientProviderOptions {
  return {
    name: KAFKA_CLIENT_TOKEN,
    transport: Transport.KAFKA,
    options: {
      client: {
        brokers: [`${process.env.KAFKA_HOST}:${process.env.KAFKA_PORT}`],
        clientId: `${appName}`,
      },
      consumer: {
        groupId: `${appName}`,
        retry: {
          restartOnFailure: (err) => {
            throw err;
          },
        },
      },
    },
  };
}
