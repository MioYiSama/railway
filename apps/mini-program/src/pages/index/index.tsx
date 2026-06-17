import { Button } from '@railway/ui-taro';
import { Text, View } from '@tarojs/components';
import { appConfig } from '../../lib/config';

export default function Index() {
  return (
    <View>
      <Text>Railway 小程序</Text>
      <Text>BACKEND_URL: {appConfig.backendUrl}</Text>
      <Button label="开始购票" onPress={() => undefined} />
    </View>
  );
}
