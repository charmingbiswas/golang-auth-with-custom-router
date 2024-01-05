import { CSSProperties, InputHTMLAttributes } from "react";
import "./Input.scss";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  styles?: CSSProperties;
  label: string;
}

const Input: React.FC<InputProps> = ({ type, styles, label, ...props }) => {
  return (
    <div className="Custom-Input">
      <input type={type} style={styles} required {...props} />
      <span>{label}</span>
      <i></i>
    </div>
  );
};

export default Input;
