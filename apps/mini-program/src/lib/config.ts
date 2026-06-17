// Taro 在编译期注入 TARO_APP_* 环境变量（见 .env.*）。
export const appConfig = {
  appUrl: process.env.TARO_APP_APP_URL ?? 'http://localhost:5173',
  backendUrl: process.env.TARO_APP_BACKEND_URL ?? 'http://localhost:8080',
};
