import { Test, TestingModule } from '@nestjs/testing';
import { AppController } from './app.controller';
import { KafkaService } from './app.service';

describe('AppController', () => {
  let appController: AppController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [AppController],
      providers: [KafkaService],
    }).compile();

    appController = app.get<AppController>(AppController);
  });

  describe('root', () => {
    it('should return "Hello World!"', () => {
      expect(appController.test1()).toBe('Hello World!');
    });
  });
});
