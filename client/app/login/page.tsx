"use client";

import { useState } from "react";

import Link from "next/link";

import {
  FormControl,
  Input,
  InputLabel,
  InputAdornment,
  IconButton,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

import { ThemeProvider } from "@emotion/react";
import { defaultTheme } from "../../assets/defaultTheme";

export default function Login() {
  const [login, setLogin] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleLogin = async () => {
    try {
      await fetch("http://localhost:8000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ login: login, password: password }),
        mode: "cors",
        credentials: "include",
      });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <ThemeProvider theme={defaultTheme}>
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
                onClick={() => handleLogin()}
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
