import { ButtonHTMLAttributes, CSSProperties, PropsWithChildren } from "react";
import "./Button.scss";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  styles?: CSSProperties;
}

export const Button: React.FC<PropsWithChildren<ButtonProps>> = ({
  children,
  styles,
  ...props
}) => {
  return (
    <button className="Custom-Button" style={styles} {...props}>
      {children}
    </button>
  );
};
