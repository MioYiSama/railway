/** Button 的统一组件 API（跨 Web / Taro / RN 三端共享契约）。 */
export type ButtonVariant = 'primary' | 'secondary' | 'ghost';

export interface ButtonProps {
  /** 按钮文案 */
  label: string;
  /** 视觉变体，默认 primary */
  variant?: ButtonVariant;
  /** 是否禁用 */
  disabled?: boolean;
  /** 点击/按下回调 */
  onPress?: () => void;
}

/** 解析变体，未指定时回退到 primary，供各端实现复用。 */
export function resolveVariant(variant?: ButtonVariant): ButtonVariant {
  return variant ?? 'primary';
}
