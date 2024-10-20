"use client";

import { FormEvent, useState } from "react";

import Link from "next/link";

import {
  FormControl,
  Input,
  InputLabel,
  InputAdornment,
  IconButton,
  ThemeProvider,
  Snackbar,
  Alert,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

import { defaultTheme } from "../../assets/defaultTheme";
import { useRouter } from "next/navigation";
import { SetGlobalError } from "@/lib/game/globalErrorSlice";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";

export default function Register() {
  const globalErrorData = useAppSelector((state) => state.globalError);
  const dispatch = useAppDispatch();

  const [open, setOpen] = useState<boolean>(false);
  const [login, setLogin] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [passwordAgain, setPasswordAgain] = useState<string>("");

  const [showPassword, setShowPassword] = useState<boolean>(false);
  const [showAgainPassword, setShowAgainPassword] = useState<boolean>(false);

  const router = useRouter();

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClickShowAgainPassword = () => {
    setShowAgainPassword(!showAgainPassword);
  };

  const handleRegister = async (e: FormEvent) => {
    e.preventDefault();

    try {
      let res = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          login: login,
          email: email,
          password: password,
          againPassword: passwordAgain,
        }),
        mode: "cors",
        credentials: "include",
      });

      if (res.ok) {
        const data = await res.json();
        console.log(data);
        router.push("/");
      } else throw res;
    } catch (err: unknown) {
      if (err instanceof Response) {
        console.error(err);
        dispatch(
          SetGlobalError({
            message: err.statusText,
            statusCode: err.status,
          })
        );
        setOpen(true);
      }
    }
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Snackbar open={open} autoHideDuration={5000} onClose={handleClose}>
        <Alert
          onClose={handleClose}
          severity="error"
          variant="filled"
          sx={{ width: "100%" }}
        >
          <b>{globalErrorData.message}</b> - Status code:{" "}
          {globalErrorData.statusCode}
        </Alert>
      </Snackbar>
      <main className="center-wrapper">
        <article className="form-panel">
          <h3 className="form-title">Register</h3>
          <div className="center-flex-col">
            <form className="center-flex-col" method="POST">
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-login">Login</InputLabel>
                <Input
                  id="register-login"
                  autoComplete="username"
                  onChange={(e) => setLogin(e.target.value)}
                ></Input>
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-email">Email</InputLabel>
                <Input
                  id="register-email"
                  autoComplete="email"
                  onChange={(e) => setEmail(e.target.value)}
                ></Input>
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-password">Password</InputLabel>
                <Input
                  id="register-password"
                  type={showPassword ? "text" : "password"}
                  onChange={(e) => setPassword(e.target.value)}
                  endAdornment={
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowPassword}
                      >
                        {showPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  }
                />
              </FormControl>
              <FormControl sx={{ width: "20em" }} variant="standard">
                <InputLabel htmlFor="register-again-password">
                  Password again
                </InputLabel>
                <Input
                  id="register-again-password"
                  type={showAgainPassword ? "text" : "password"}
                  onChange={(e) => {
                    setPasswordAgain(e.target.value);
                  }}
                  endAdornment={
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowAgainPassword}
                      >
                        {showAgainPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  }
                />
              </FormControl>
              <p className="form-prompt">
                Already have an existing account?{" "}
                <Link href="/login">Login here</Link>
              </p>
              <button
                type="submit"
                onClick={(e) => handleRegister(e)}
                className="submit-button"
              >
                Register
              </button>
            </form>
          </div>
        </article>
      </main>
    </ThemeProvider>
  );
}
