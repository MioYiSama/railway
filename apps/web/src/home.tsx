import { Button } from '@railway/ui-web';
import { useQuery } from '@tanstack/react-query';
import { env } from './env';

export function Home() {
  const { data } = useQuery({
    queryKey: ['health'],
    // 骨架阶段直接返回 mock，真实实现会请求 `${env.backendUrl}/healthz`。
    queryFn: async () => ({ status: 'ok' }),
  });

  return (
    <main>
      <h1>Railway Web</h1>
      <p>APP_URL: {env.appUrl}</p>
      <p>BACKEND_URL: {env.backendUrl}</p>
      <p>health: {data?.status ?? 'loading'}</p>
      <Button label="开始购票" onPress={() => undefined} />
    </main>
  );
}
