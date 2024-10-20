"use client";

import { FormEvent, useState } from "react";

import Link from "next/link";

import {
  Alert,
  FormControl,
  Input,
  InputLabel,
  InputAdornment,
  IconButton,
  Snackbar,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

import { ThemeProvider } from "@emotion/react";
import { defaultTheme } from "../../assets/defaultTheme";
import { useRouter } from "next/navigation";
import { useAppDispatch, useAppSelector } from "@/lib/hooks";
import { SetGlobalError } from "@/lib/game/globalErrorSlice";

export default function Login() {
  const globalErrorData = useAppSelector((state) => state.globalError);
  const dispatch = useAppDispatch();
  const [open, setOpen] = useState<boolean>(false);
  const [login, setLogin] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const router = useRouter();

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleLogin = async (e: FormEvent) => {
    e.preventDefault();

    try {
      let res = await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ login: login, password: password }),
        mode: "cors",
        credentials: "include",
      });
      let data = await res.json();
      if (res.status < 300 && res.status >= 200) {
        console.log(data);
        router.push("/game");
      } else throw res;
    } catch (err: unknown) {
      if (err instanceof Response) {
        console.error(err);
        dispatch(
          SetGlobalError({
            message: String(err.statusText),
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
          <h3 className="form-title">Login</h3>
          <div className="center-flex-col">
            <form className="center-flex-col" method="POST">
              <FormControl sx={{ width: "16em" }} variant="standard">
                <InputLabel htmlFor="login-login">Login</InputLabel>
                <Input
                  id="login-login"
                  onChange={(e) => setLogin(e.target.value)}
                ></Input>
              </FormControl>
              <FormControl sx={{ width: "16em" }} variant="standard">
                <InputLabel htmlFor="login-password">Password</InputLabel>
                <Input
                  id="login-password"
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
              <p className="form-prompt">
                Don't have an account yet?{" "}
                <Link href="register">Register here</Link>
              </p>
              <button
                type="submit"
                onClick={(e) => handleLogin(e)}
                className="submit-button"
              >
                Login
              </button>
            </form>
          </div>
        </article>
      </main>
    </ThemeProvider>
  );
}
