import { useState } from "react";
import { CustomButton } from "../components/Button";
import { CustomInput } from "../components/Input";
import "./LoginPage.scss";
import { Button, Typography } from "@mui/material";

const LoginPage = () => {
  const [showLogin, setShowLogin] = useState<boolean>(false);

  const changePage = () => {
    setShowLogin((prev) => !prev);
  };

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
      {showLogin ? (
        <div className="Login-Card">
          <CustomInput label="Email" id="email" />
          <CustomInput label="Password" id="password" />
          <CustomButton styles={{ width: "10rem" }} onClick={handleClick}>
            Login
          </CustomButton>
          <Typography color="white">
            New Customer?{" "}
            <Button variant="text" size="medium" onClick={changePage}>
              Sign Up
            </Button>
          </Typography>
        </div>
      ) : (
        <div className="Signup-Card">
          <CustomInput label="Username" id="username" />
          <CustomInput label="Email" id="username" />
          <CustomInput label="Password" id="password" />
          <CustomButton styles={{ width: "10rem" }} onClick={handleClick}>
            Signup
          </CustomButton>
          <Button variant="text" size="medium" onClick={changePage}>
            Back To Login
          </Button>
        </div>
      )}
    </div>
  );
};

export default LoginPage;
