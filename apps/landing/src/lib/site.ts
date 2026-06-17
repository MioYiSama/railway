/** 落地页站点配置（骨架）。 */
export const site = {
  title: 'Railway 12306（学习用途）',
  appUrl: import.meta.env.PUBLIC_APP_URL ?? '',
  backendUrl: import.meta.env.PUBLIC_BACKEND_URL ?? '',
};

/** 拼接进入主应用的链接。 */
export function appLink(path = '/'): string {
  const base = site.appUrl.replace(/\/$/, '');
  return `${base}${path}`;
}
