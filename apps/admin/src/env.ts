/** 运行时环境变量（区分 development / production 由 Vite mode 决定）。 */
export const env = {
  appUrl: import.meta.env.VITE_APP_URL ?? '',
  backendUrl: import.meta.env.VITE_BACKEND_URL ?? '',
  mode: import.meta.env.MODE,
};
