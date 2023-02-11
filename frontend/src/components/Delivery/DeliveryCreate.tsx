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
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Autocomplete from "@mui/material/Autocomplete";

import { DeliveryInterface } from "../../models/thanadet/IDelivery"
import { CarInterface } from "../../models/thanadet/ICar"
import { PaymentInterface } from "../../models/Natthapon/IPayment"
import { EmployeeInterface } from "../../models/IEmployee"


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function DeliveryCreate() {
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");
    const [message, setAlertMessage] = React.useState("");

    const [employee, setEmployee] = React.useState<EmployeeInterface[]>([]);
    const [car, setCar] = React.useState<CarInterface[]>([]);
    const [payment, setPayment] = React.useState<PaymentInterface[]>([]);
    const [delivery, setDelivery] = React.useState<DeliveryInterface>({
        Delivery_date: new Date(),
    });

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

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof DeliveryCreate;
        const { value } = event.target;
        setDelivery({ ...delivery, [id]: value });
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof delivery;
        setDelivery({
            ...delivery,
            [name]: event.target.value,
        });
    };

    const getCar = async () => {
        fetch(`${apiUrl}/cars`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setCar(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getPayment = async () => {
        fetch(`${apiUrl}/payments`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setPayment(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getEmployee = async () => {
        fetch(`${apiUrl}/employees`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setEmployee(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    useEffect(() => {
        getEmployee();
        getCar();
        getPayment();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            Location: delivery.Location ?? "",
            Customer_name: delivery.Customer_name ?? "",
            Delivery_date: delivery.Delivery_date,
            Employee_ID: convertType(delivery.Employee_ID),
            Car_ID: convertType(delivery.Car_ID),
            Payment_ID: convertType(delivery.Payment_ID),
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

        let res = await fetch(`${apiUrl}/deliveries`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("")
                    return { status: true, message: res.data };
                } else {
                    setError(true);
                    setErrorMessage(res.error)
                    return { status: false, message: res.error };
                }
            });

        if (res.status) {
            setAlertMessage("บันทึกข้อมูลสำเร็จ");
            setSuccess(true);
            setTimeout(() => {
                window.location.reload();
            }, 0);
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar
                id="success"
                open={success}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    {message}
                </Alert>
            </Snackbar>
            <Snackbar
                id="error"
                open={error}
                autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    {message}
                </Alert>
            </Snackbar>
            <Paper>
                <Box
                    display="flex"
                    sx={{
                        marginTop: 2,
                    }}
                >
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            เพิ่มรายการการสั่งซื้อ
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>

                    <Grid item xs={6}>
                        <p className="good-font">ชื่อ - นามสกุล ของลูกค้า</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField
                                id="Customer_name"
                                variant="outlined"
                                type="string"
                                size="medium"
                                value={delivery.Customer_name || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <p className="good-font">สถานที่จัดส่งสินค้า</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField
                                id="Location"
                                variant="outlined"
                                type="string"
                                size="medium"
                                value={delivery.Location || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รถยนต์ที่ใช้</p>
                            <Autocomplete
                                disablePortal
                                id="Car_ID"
                                getOptionLabel={(item: CarInterface) => `${item.Car_Model} ${item.Registation_Number}`}
                                options={car}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { delivery.Car_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกรถยนต์ที่ใช้" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รายการสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Payment_ID"
                                getOptionLabel={(item: PaymentInterface) => `${item.ID}`}
                                options={payment}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { delivery.Payment_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกรายการสินค้า" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">พนักงานที่ส่งสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Employee_ID"
                                getOptionLabel={(item: EmployeeInterface) => `${item.Name}`}
                                options={employee}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { delivery.Employee_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกพนักงานที่ส่งสินค้า" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">วันที่ส่งสินค้า</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={delivery.Delivery_date}
                                    onChange={(newValue) => {
                                        setDelivery({
                                            ...delivery,
                                            Delivery_date: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/Delivery" variant="contained">
                            กลับ
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            บันทึก
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default DeliveryCreate;
