import { type ButtonProps, resolveVariant } from '@railway/ui-core';

/** Button 的 Web (React) 实现，遵循 @railway/ui-core 契约。 */
export function Button({ label, variant, disabled, onPress }: ButtonProps) {
  return (
    <button
      type="button"
      data-variant={resolveVariant(variant)}
      disabled={disabled}
      onClick={onPress}
    >
      {label}
    </button>
  );
}
