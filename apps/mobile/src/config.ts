// React Native 没有 import.meta.env，这里用 __DEV__ 区分 development / production。
// 真实项目可换成 react-native-config 读取 .env.* 文件（见同目录 .env.*）。
declare const __DEV__: boolean | undefined;

const isDev = typeof __DEV__ !== 'undefined' ? __DEV__ : process.env.NODE_ENV !== 'production';

export const appConfig = {
  appUrl: isDev ? 'http://localhost:5173' : 'https://web.railway.local',
  backendUrl: isDev ? 'http://localhost:8080' : 'https://api.railway.local',
};
