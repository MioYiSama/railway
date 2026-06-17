import { type ButtonProps, resolveVariant } from '@railway/ui-core';
import { Pressable, Text } from 'react-native';

/** Button 的 React Native 实现，遵循 @railway/ui-core 契约。 */
export function Button({ label, variant, disabled, onPress }: ButtonProps) {
  // resolveVariant 复用核心契约逻辑，后续可据此映射 RN 样式。
  resolveVariant(variant);
  return (
    <Pressable accessibilityRole="button" disabled={disabled} onPress={onPress}>
      <Text>{label}</Text>
    </Pressable>
  );
}
