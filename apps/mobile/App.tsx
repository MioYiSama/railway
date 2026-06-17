import { Button } from '@railway/ui-native';
import { SafeAreaView, Text } from 'react-native';
import { appConfig } from './src/config';

export default function App() {
  return (
    <SafeAreaView>
      <Text>Railway Mobile</Text>
      <Text>BACKEND_URL: {appConfig.backendUrl}</Text>
      <Button label="开始购票" onPress={() => undefined} />
    </SafeAreaView>
  );
}
