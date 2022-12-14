import React, { useState } from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { SigninInterface } from "../models/ISignin";
import { Login } from "../services/HttpClientService";
import { InputAdornment } from "@mui/material";

import EmailIcon from '@mui/icons-material/Email';
import PasswordIcon from '@mui/icons-material/Password';
import LoginIcon from '@mui/icons-material/Login';

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

const theme = createTheme();

function SignIn() {
    const [signin, setSignin] = useState<Partial<SigninInterface>>({});
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof signin;
        const { value } = event.target;
        setSignin({ ...signin, [id]: value });
    };

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const submit = async () => {
        let res = await Login(signin);
        if (res) {
            setSuccess(true);
            setTimeout(() => {
                window.location.reload();
            }, 1000);
        } else {
            setError(true);
        }
    };

    return (
        <ThemeProvider theme={theme}>
            <Grid
                item xs={false} sm={8} md={12} container component="main"
                sx={{
                    backgroundImage: "url(https://sv1.picz.in.th/images/2022/10/23/vfPECa.jpg)",
                    backgroundPosition: "center", height: "100vh", width: '100%',
                }} >
                <Snackbar
                    open={success}
                    autoHideDuration={3000}
                    onClose={handleClose}
                    anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                >
                    <Alert onClose={handleClose} severity="success"
                        sx={{ width: 240, height: 60, alignItems: "center", alignSelf: "center", }}>
                        <div className="good-font"> ??????????????????????????????????????????????????? </div>
                    </Alert>
                </Snackbar>
                <Snackbar
                    open={error}
                    autoHideDuration={3000}
                    onClose={handleClose}
                    anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                >
                    <Alert onClose={handleClose} severity="error"
                        sx={{ width: 320, height: 60, alignItems: "center", alignSelf: "center", }}>
                        <div className="good-font"> ????????????????????????????????????????????????????????????????????????????????? </div>
                    </Alert>
                </Snackbar>

                <CssBaseline />
                <Grid item xs={8}> <Paper elevation={0} /> </Grid>
                <Grid item xs={5} sm={8} md={4} component={Paper} elevation={6} square container spacing={2}
                    sx={{ my: 'auto', mx: '-10%', width: 300, height: 420, borderRadius: 2, alignItems: "center", alignSelf: "center", }}>

                    <Typography component="h1" variant="h4" sx={{ m: 'auto' }}>
                        <Avatar sx={{ mx: 'auto', my: '5%', bgcolor: "#03a9f4", width: 56, height: 56 }}> <LockOutlinedIcon /> </Avatar>
                        <div className="good-font"> Sign in to website </div>
                    </Typography>
                    <Box sx={{ mx: 5, display: "flex", flexDirection: "column", alignItems: "center", alignSelf: "center", }}>
                        <Box sx={{ mt: 0 }}>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                id="Email"
                                label="Email Address"
                                name="email"
                                autoComplete="email"
                                autoFocus
                                value={signin.Email || ""}
                                onChange={handleInputChange}
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <EmailIcon />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                name="password"
                                label="Password"
                                type="password"
                                id="Password"
                                autoComplete="current-password"
                                value={signin.Password || ""}
                                onChange={handleInputChange}
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <PasswordIcon />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            <FormControlLabel
                                control={<Checkbox value="remember" color="primary" />}
                                label="Remember me"
                            />
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                sx={{ mt: 2, mb: 6 }}
                                onClick={submit}
                                endIcon={<LoginIcon />}
                            >
                                <div className="good-font" > Sign In </div>
                            </Button>
                        </Box>
                    </Box>
                </Grid>
            </Grid>
        </ThemeProvider>
    );
}

export default SignIn;
