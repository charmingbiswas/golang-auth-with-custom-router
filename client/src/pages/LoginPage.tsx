import { useState } from "react";
import { CustomButton } from "../components/Button";
import { CustomInput } from "../components/Input";
import "./LoginPage.scss";
import { Alert, Button, Snackbar, Typography } from "@mui/material";
import axios from "axios";
import { Error } from "../types";

const LoginPage = () => {
  const [showLogin, setShowLogin] = useState<boolean>(false);
  const [openSnackbar, setOpenSnackBar] = useState<boolean>(false);
  const [serverMessage, setServerMessage] = useState<string>("");
  const [isError, setIsError] = useState<boolean>(false);

  const changePage = () => {
    setShowLogin((prev) => !prev);
  };

  const handleClick = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    if (e.currentTarget.innerText === "SIGNUP") {
      const password = document.querySelector("#password") as HTMLInputElement;
      const email = document.querySelector("#email") as HTMLInputElement;
      const username = document.querySelector("#username") as HTMLInputElement;
      if (email && password && username) {
        const dataToSend = JSON.stringify({
          name: username.value,
          email: email.value,
          password: password.value,
        });
        try {
          const res = await axios.post(
            "http://localhost:4000/api/v1/signup",
            dataToSend,
            {
              headers: {
                "Content-Type": "application/json",
              },
            }
          );
          const data = res.data;
          setIsError(false);
          setServerMessage(data);
        } catch (err) {
          console.log(err);
          setIsError(true);
          const error = err as Error;
          setServerMessage(error.response.data);
        } finally {
          setOpenSnackBar(true);
        }
      }
    } else if (e.currentTarget.innerText === "LOGIN") {
      const password = document.querySelector("#password") as HTMLInputElement;
      const email = document.querySelector("#email") as HTMLInputElement;
      if (email && password) {
        try {
          const res = await axios.post(
            "http://localhost:4000/api/v1/signin",
            JSON.stringify({
              email: email.value,
              password: password.value,
            })
          );
          const data = res.data;
          setIsError(false);
          setServerMessage(data);
        } catch (err) {
          console.log(err);
          setIsError(true);
          const error = err as Error;
          setServerMessage(error.response.data);
        } finally {
          setOpenSnackBar(true);
        }
      }
    }
  };
  return (
    <div className="Login-Page">
      <Snackbar
        open={openSnackbar}
        autoHideDuration={5000}
        onClose={() => setOpenSnackBar(false)}
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <Alert
          onClose={() => setOpenSnackBar(false)}
          severity={isError ? "error" : "success"}
          sx={{ width: "100%" }}
        >
          {serverMessage}
        </Alert>
      </Snackbar>
      {showLogin ? (
        <div className="Login-Card">
          <CustomInput label="Email" id="email" />
          <CustomInput label="Password" id="password" />
          <CustomButton styles={{ width: "10rem" }} onClick={handleClick}>
            Login
          </CustomButton>
          <Typography color="white">
            New Customer?{" "}
            <Button variant="contained" size="medium" onClick={changePage}>
              Sign Up
            </Button>
          </Typography>
        </div>
      ) : (
        <div className="Signup-Card">
          <CustomInput label="Username" id="username" />
          <CustomInput label="Email" id="email" />
          <CustomInput label="Password" id="password" />
          <CustomButton styles={{ width: "10rem" }} onClick={handleClick}>
            Signup
          </CustomButton>
          <Button variant="contained" size="medium" onClick={changePage}>
            Back To Login
          </Button>
        </div>
      )}
    </div>
  );
};

export default LoginPage;
