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
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Autocomplete from "@mui/material/Autocomplete";

import { EmployeeInterface } from "../../models/IEmployee";
import { MemberInterface } from "../../models/theerawat/IMember";
import { OrderInterface } from "../../models/Natthapon/IOrder";
import { StocksInterface } from "../../models/methas/IStock";
import { CartInterface } from "../../models/Natthapon/ICart";
import { StatusInterface } from "../../models/Natthapon/IStatus";
import { GetCurrentEmployee } from "../../services/HttpClientService";


const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function OrderCreate() {
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");

    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [member, setMember] = React.useState<MemberInterface[]>([]);
    const [shelving, setShelving] = React.useState<StocksInterface[]>([]);
    const [status, setStatus] = React.useState<StatusInterface[]>([]);
    const [order, setOder] = React.useState<OrderInterface>({});
    const [cart, setCart] = React.useState<CartInterface>({});
    const [stock, setStock] = React.useState<StocksInterface[]>([]);

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
        const id = event.target.id as keyof typeof OrderCreate;
        const { value } = event.target;
        setOder({ ...order, [id]: value });
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof order;
        setOder({
            ...order,
            [name]: event.target.value,
        });
    };

    const getShelving = async () => {
        fetch(`${apiUrl}/stock`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
            if (res.data) {
                console.log(res.data)
                setShelving(res.data);
            }
            else { console.log("NO DATA") }
            });
    };

    const getMember = async () => {
        fetch(`${apiUrl}/member`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
            if (res.data) {
                console.log(res.data)
                setMember(res.data);
            }
            else { console.log("NO DATA") }
            });
    };

    const getStatus = async () => {
        fetch(`${apiUrl}/status`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
            if (res.data) {
                console.log(res.data)
                setStatus(res.data);
            }
            else { console.log("NO DATA") }
            });
      };

    const getEmployee = async () => {
        let res = await GetCurrentEmployee();
        cart.Employee_ID = res.ID;
        if (res) {
            setEmployee(res);
            console.log(res)
        }
    };

    useEffect(() => {
        getEmployee();
        getMember();
        getShelving();
        getStatus();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function addcart() {
        let data = {
            Total: typeof cart.Total === "string" ? parseInt(cart.Total) : 0,
            Status_ID: 1,
            Member_ID: convertType(cart.Member_ID),
            Employee_ID: convertType(cart.Employee_ID),
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

        fetch(`${apiUrl}/carts`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("")
                } else {
                    setError(true);
                    setErrorMessage(res.error)
                }
            });
    }


    async function submit() {
        let data = {
            Quantity: typeof order.Quantity === "string" ? parseInt(order.Quantity) : 0,
            Shelving_ID: convertType(order.Shelving_ID),
            Shopping_Cart_ID: convertType(order.Shopping_Cart_ID),
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

        fetch(`${apiUrl}/orders`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess(true);
                    setErrorMessage("")
                } else {
                    setError(true);
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
                        เพิ่มสินค้าลงตะกร้าแล้ว
                    </div>
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    <div className="good-font">
                        เพิ่มสินค้าไม่สำเร็จ
                    </div>
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
                            <div className="good-font">
                                ตะกร้าสินค้า
                            </div>
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รหัสสมาชิก</p>
                            <Autocomplete
                                disablePortal
                                id="Member_ID"
                                getOptionLabel={(item: MemberInterface) => `${item.ID}`}
                                options={member}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { cart.Member_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="รหัสสมาชิก" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">พนักงานที่บันทึก</p>
                            <Select
                                native
                                value={cart.Employee_ID + ""}
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

                    <Grid item xs={12}>
                        <Button
                            style={{ display: "flex", justifyContent: "center", margin: "0 auto" }}
                            onClick={addcart}
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                สร้างตะกร้า
                            </div>
                        </Button>
                    </Grid>
                </Grid>

                    
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รายการสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Stock_ID"
                                getOptionLabel={(item: StocksInterface) => `${item.Name} ราคา ${item.Price}`}
                                options={stock}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { order.Shelving_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกสินค้า" />}
                            />
                        </FormControl>
                    </Grid>


                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">จำนวน</p>
                            <TextField
                                id="Quantity"
                                variant="outlined"
                                type="number"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 , max: 50}}}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={order.Quantity || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={3}>
                        <Button
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ใส่ตะกร้า
                            </div>
                        </Button>
                    </Grid>
    
                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/Discount" variant="contained">
                            <div className="good-font-white">
                                กลับ
                            </div>
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default OrderCreate;
