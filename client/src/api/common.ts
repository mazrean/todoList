export class Error {
  code: number;
  error: string;

  constructor(code: number, error: string) {
    this.code = code;
    this.error = error;
  }
}

export class Message {
  message: string;
}
