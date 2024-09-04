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
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <main className="center-wrapper">
        <article className="form-panel">
          <h3 className="form-title">Login</h3>
          <div className="center-flex-col">
            <FormControl sx={{ width: "16em" }} variant="standard">
              <InputLabel htmlFor="standard-login">Login</InputLabel>
              <Input id="standard-login"></Input>
            </FormControl>
            <FormControl sx={{ width: "16em" }} variant="standard">
              <InputLabel htmlFor="standard-password">Password</InputLabel>
              <Input
                id="standard-password"
                type={showPassword ? "text" : "password"}
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
            <button type="submit" className="submit-button">
              Login
            </button>
          </div>
        </article>
      </main>
    </ThemeProvider>
  );
}
