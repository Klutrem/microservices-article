import { Inject, Injectable, Logger } from '@nestjs/common';
import { ClientKafka } from '@nestjs/microservices';

import { firstValueFrom, tap } from 'rxjs';

export const OUTPUT_TOPICS_TOKEN = 'OUTPUT_TOPICS';
export const KAFKA_CLIENT_TOKEN = 'KAFKA_CLIENT';

@Injectable()
export default class KafkaService {
  constructor(
    @Inject(KAFKA_CLIENT_TOKEN) private readonly client: ClientKafka,
    @Inject(OUTPUT_TOPICS_TOKEN) private readonly topics: string[],
    private logger: Logger,
  ) {
    this.logger = new Logger();
  }

  /**
   * Subscibes topics to client and initiates connection
   */
  async onModuleInit() {
    this.topics.forEach((topic) => this.addTopic(topic));
    await this.client.connect();
  }

  /**
   * Adds new topics to the service
   * @param topic topic to add
   */
  addTopic(topic: string) {
    this.logger.debug(`subscribing to topic ${topic}.reply`);
    this.client.subscribeToResponseOf(topic);
  }

  /**
   * Sends the message to the topic
   * @param topic The topic
   * @param data The data
   * @returns The response
   */
  send<TOutput = unknown>(
    topic: string,
    data: unknown,
  ): Promise<TOutput | undefined> {
    this.logOutgoing(topic, data);
    return firstValueFrom(
      this.client.send(topic, JSON.stringify(data)).pipe(
        tap({
          next: (r) => this.logResponse(topic, r),
          error: (err) => this.handleResponseError(topic, err),
        }),
      ),
    );
  }

  private logOutgoing(topic: string, data: unknown) {
    this.logger.debug(`Message sent to '${topic}' with data:`);
    this.logger.debug(JSON.stringify(data));
  }

  private logResponse(topic: string, response: unknown) {
    this.logger.debug(`[${topic}] Response came: ${JSON.stringify(response)}`);
  }

  private handleResponseError(topic: string, err: Error) {
    this.logger.error(`[${topic}] Error came: ${JSON.stringify(err)}`);
    throw err;
  }
}
