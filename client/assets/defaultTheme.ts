import { createTheme } from "@mui/material";

export const defaultTheme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "rgb(100, 57, 255)",
    },
  },
  typography: {
    fontFamily: "Space Grotesk",
  },
  components: {
    MuiFormControl: {
      styleOverrides: {
        root: {
          margin: "1em 0",
          padding: 0,
        },
      },
    },
    MuiTypography: {
      styleOverrides: {
        root: {
          fontFamily: '"Space Grotesk", sans-serif',
        },
      },
    },
  },
});
