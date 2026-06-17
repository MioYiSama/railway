import type { ReactNode } from 'react';

/** 测试用的 react-native 轻量替身（映射到 DOM 元素）。 */
export function Pressable(props: {
  children?: ReactNode;
  disabled?: boolean;
  onPress?: () => void;
  accessibilityRole?: string;
}) {
  return (
    <button type="button" disabled={props.disabled} onClick={props.onPress}>
      {props.children}
    </button>
  );
}

export function Text(props: { children?: ReactNode }) {
  return <span>{props.children}</span>;
}
