import { CustomButton } from "../components/Button";
import { CustomInput } from "../components/Input";
import "./LoginPage.scss";

const LoginPage = () => {
  const handleClick = async () => {
    const password = document.querySelector("#password") as HTMLInputElement;
    const username = document.querySelector("#username") as HTMLInputElement;
    if (username && password) {
      try {
        const res = await fetch("http://localhost:4000/api/v1/signin");
        const data = await res.json();
        console.log(data);
      } catch (err) {
        console.log(err);
      }
    }
  };
  return (
    <div className="Login-Page">
      <div className="Login-Card">
        <CustomInput label="Username" id="username" />
        <CustomInput label="Password" id="password" />
        <CustomButton styles={{ width: "10rem" }} onClick={handleClick}>
          Login
        </CustomButton>
      </div>
    </div>
  );
};

export default LoginPage;
