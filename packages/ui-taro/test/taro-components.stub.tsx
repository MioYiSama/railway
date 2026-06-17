import type { ReactNode } from 'react';

/** 测试用的 @tarojs/components 轻量替身。 */
export function Button(props: {
  children?: ReactNode;
  disabled?: boolean;
  onClick?: () => void;
  'data-variant'?: string;
}) {
  return (
    <button type="button" disabled={props.disabled} onClick={props.onClick} data-variant={props['data-variant']}>
      {props.children}
    </button>
  );
}
