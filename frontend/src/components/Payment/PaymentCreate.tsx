import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs, { Dayjs } from 'dayjs';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Autocomplete from "@mui/material/Autocomplete";

import { EmployeeInterface } from "../../models/IEmployee"
import { CartInterface } from "../../models/Natthapon/ICart"
import { Payment_methodInterface } from "../../models/Natthapon/IPayment_method"
import { PaymentInterface } from "../../models/Natthapon/IPayment"
import { GetCurrentEmployee } from "../../services/HttpClientService";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function PaymentCreate() {
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");
    const [message, setAlertMessage] = React.useState("");

    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [methods, setMethod] = React.useState<Payment_methodInterface[]>([]);
    const [cart, setCart] = React.useState<CartInterface[]>([]);
    const [payment, setPayment] = React.useState<PaymentInterface>({
        Time: new Date(),
    });
    const [latestCartId, setLatestCartId] = React.useState(0);
    const [sumprice, setSumprice] = React.useState(0);

    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
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

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof payment;
        setPayment({
            ...payment,
            [name]: event.target.value,
        });
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof PaymentCreate;
        const { value } = event.target;
        setPayment({ ...payment, [id]: value });
    };

    const getPayment_method = async () => {
        fetch(`${apiUrl}/payment_methods`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setMethod(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getCart = async () => {
        fetch(`${apiUrl}/unpaids`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setCart(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getEmployee = async () => {
        let res = await GetCurrentEmployee();
        payment.Employee_ID = res.ID;
        if (res) {
            setEmployee(res);
            console.log(res)
        }
    };

    const getLatestCartId = async () => {
        fetch(`${apiUrl}/unpaids`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                // Find the cart with the highest ID
                let latestCart = res.data.reduce((prev: any, current: any) => {
                    return (prev.ID > current.ID) ? prev : current
                });
                setLatestCartId(latestCart.ID);
            }
        });
    }

    useEffect(() => {
        getEmployee();
        getPayment_method();
        getCart();
        getLatestCartId();
    }, []);

    fetch(`${apiUrl}/ordersum/${latestCartId}`, requestOptions)
        .then((response) => response.json())
        .then(data => {
            let sumPrices = data.sumPrices;
            setSumprice(sumPrices);
            console.log(sumPrices)
            // Use the sumPrices variable as needed
            
    });

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            Paytotal: typeof payment.Paytotal === "string" ? parseFloat(payment.Paytotal) : sumprice,
            Time: payment.Time,
            Note: payment.Note ?? "",
            Shopping_Cart_ID: latestCartId,
            Payment_method_ID: convertType(payment.Payment_method_ID),
            Employee_ID: convertType(payment.Employee_ID),
        };

        console.log(data)

        const requestOptions = {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        let res = await fetch(`${apiUrl}/payments`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("")
                    pay()
                    return { status: true, message: res.data };
                } else {
                    setError(true);
                    setErrorMessage(res.error)
                    return { status: false, message: res.error };
                }
            });
            if (res.status) {
                setAlertMessage("บันทึกสำเร็จ");
                setSuccess(true);
              } else {
                setAlertMessage(res.message);
                setError(true);
              }
    }
    async function pay() {
        let data = {
            Status_ID: 2,
        };
        console.log(data)

        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        fetch(`${apiUrl}/cart/${latestCartId}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setErrorMessage("")            
                } else {
                    setErrorMessage(res.error)
                }
            });

    }


    return (
        <Container maxWidth="md">
            <Snackbar
                open={success}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    <div className="good-font">
                    {message}
                    </div>
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
                <Alert onClose={handleClose} severity="error">
                    <div className="good-font">
                    {message}
                    </div>
                </Alert>
            </Snackbar>
            <Paper>
                <Box display="flex" sx={{marginTop: 2}}>
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom

                        >
                            <div className="good-font">
                                ชำระสินค้า
                            </div>
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}> 
                <Grid item xs={6}>
                        <p className="good-font">ตะกร้าสินค้า</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField
                            id="Shopping_Cart_ID"
                            variant="outlined"
                            type="string"
                            size="medium"
                            value={latestCartId}
                            onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>
                    
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">พนักงานที่บันทึก</p>
                            <Select
                                native
                                value={payment.Employee_ID + ""}
                                onChange={handleChange}
                                disabled
                                inputProps={{
                                    name: "Employee_ID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    เลือก
                                </option>
                                <option value={employee?.ID} key={employee?.ID}>
                                    {employee?.Name}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">ช่องทางการชำระสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Payment_method_ID"
                                getOptionLabel={(item: Payment_methodInterface) => `${item.Method}`}
                                options={methods}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { payment.Payment_method_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกช่องทางการชำระ" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">เวลาที่ชำระสินค้า</p>
                            <LocalizationProvider dateAdapter={AdapterDayjs}>
                                <DateTimePicker
                                    renderInput={(props) => <TextField {...props} />}
                                    value={payment.Time}
                                    onChange={(newValue) => {
                                        setPayment({
                                            ...payment,
                                            Time: newValue,
                                        });
                                    }}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    
                    <Grid item xs={6}>
                        <p className="good-font">ยอดชำระ</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField
                            id="Total"
                            variant="outlined"
                            type="string"
                            size="medium"
                            value={payment.Paytotal || sumprice}
                            onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <p className="good-font">หมายเหตุ</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField
                            id="Note"
                            variant="outlined"
                            type="string"
                            size="medium"
                            value={payment.Note || ""}
                            onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>


                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/Payment" variant="contained">
                            <div className="good-font-white">
                                กลับ
                            </div>
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ยืนยันการชำระ
                            </div>
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default PaymentCreate;
