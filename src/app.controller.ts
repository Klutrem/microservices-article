import { Controller, Get } from '@nestjs/common';

@Controller()
export class AppController {
  constructor() {}

  @Get('test1')
  test1() {
    return Promise.resolve('hoy');
  }

  @Get('test2')
  test2() {
    return Promise.resolve('huy');
  }

  @Get('test3')
  test3() {
    return Promise.resolve('pizds');
  }
}
