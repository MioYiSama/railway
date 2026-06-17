import { Button as TaroButton } from '@tarojs/components';
import { type ButtonProps, resolveVariant } from '@railway/ui-core';

/** Button 的 Taro 小程序实现，遵循 @railway/ui-core 契约。 */
export function Button({ label, variant, disabled, onPress }: ButtonProps) {
  return (
    <TaroButton data-variant={resolveVariant(variant)} disabled={disabled} onClick={onPress}>
      {label}
    </TaroButton>
  );
}
