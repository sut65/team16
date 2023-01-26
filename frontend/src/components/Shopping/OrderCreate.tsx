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
import { IShelving } from "../../models/methas/IShelving";
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
    const [success1, setSuccess1] = React.useState(false);
    const [error1, setError1] = React.useState(false);
    const [success2, setSuccess2] = React.useState(false);
    const [error2, setError2] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");

    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [member, setMember] = React.useState<MemberInterface[]>([]);
    const [shelving, setShelving] = React.useState<IShelving[]>([]);
    const [status, setStatus] = React.useState<StatusInterface[]>([]);
    const [order, setOder] = React.useState<OrderInterface>({});
    const [cart, setCart] = React.useState<CartInterface>({});

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
        setSuccess1(false);
        setError1(false);
        setSuccess2(false);
        setError2(false);
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
        fetch(`${apiUrl}/Shelving`, requestOptions)
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
        fetch(`${apiUrl}/members`, requestOptions)
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
        fetch(`${apiUrl}/statuses`, requestOptions)
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
                    setSuccess1(true);
                    setErrorMessage("")
                } else {
                    setError1(true);
                    setErrorMessage(res.error)
                }
            });
    }


    async function addproduct() {
        let data = {
            Quantity: typeof order.Quantity === "string" ? parseInt(order.Quantity) : 0,
            Shelving_ID: convertType(order.Shelving_ID),
            //Shopping_Cart_ID: convertType(order.Shopping_Cart_ID),
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
                    const cartId = res.data.ID;
                    console.log("Last inserted Shopping Cart ID: " + cartId);
                    setSuccess2(true);
                    setErrorMessage("")
                } else {
                    setError2(true);
                    setErrorMessage(res.error)
                }
            });
    }

    return (
        <Container maxWidth="md">
            <Snackbar
                open={success1}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    <div className="good-font">
                        เพิ่มตะกร้าสำเร็จ
                    </div>
                </Alert>
            </Snackbar>
            <Snackbar open={error1} 
                autoHideDuration={6000} 
                onClose={handleClose} 
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
                <Alert onClose={handleClose} severity="error">
                    <div className="good-font">
                        เพิ่มตะกร้าไม่สำเร็จ
                    </div>
                </Alert>
            </Snackbar>
            <Snackbar
                open={success2}
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
            <Snackbar open={error2} 
                autoHideDuration={6000} 
                onClose={handleClose} 
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
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
                            <p className="good-font">เบอร์โทรศัพท์</p>
                            <Autocomplete
                                disablePortal
                                id="Member_ID"
                                getOptionLabel={(item: MemberInterface) => `${item.Mem_Tel}`}
                                options={member}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.Mem_Tel === value.Mem_Tel}
                                onChange={(e, value) => { 
                                    // find the member that matches the selected phone number
                                    let selectedMember = member.find(member => member.Mem_Tel === value?.Mem_Tel)
                                    // set the cart.Member_ID with the ID of the selected member
                                    cart.Member_ID = selectedMember?.ID 
                                }}
                                renderInput={(params) => <TextField {...params} label="เบอร์โทรศัพท์" />}
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
                    {/* <Grid item xs={9}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รายการสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Product"
                                getOptionLabel={(item: IShelving) => `${item.Stock.Name} ราคา ${item.Stock.Price}`}
                                options={shelving}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) => option.Stock_ID === value.Stock_ID}
                                onChange={(e, value) => { order.Shelving_ID = value?.Stock_ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกสินค้า" />}
                            />
                        </FormControl>
                    </Grid> */}
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">สินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Stock_ID"
                                getOptionLabel={(item: IShelving) => `${item.ID}`}
                                options={shelving}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { order.Shelving_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกสินค้า" />}
                            />
                        </FormControl>
                    </Grid>


                    <Grid item xs={3}>
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

                    <Grid item xs={12}>
                        <Button
                            style={{ display: "flex", justifyContent: "center", margin: "0 auto" }}
                            onClick={addproduct}
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ใส่ตะกร้า
                            </div>
                        </Button>
                    </Grid>
    
                </Grid>

            </Paper>
            <Paper>
            <Grid container spacing={20} sx={{ padding: 2 }}>
                <Grid item xs={6}>
                    <Button component={RouterLink} to="/Cart" variant="contained">
                        <div className="good-font-white">
                            ตะกร้าสินค้า
                        </div>
                    </Button>
                </Grid>
                <Grid item xs={6}>
                    <Button component={RouterLink} to="/Order" variant="contained" style={{ float: "right" }}>
                        <div className="good-font-white">
                            รายการสินค้า
                        </div>
                    </Button>
                </Grid>
            </Grid>
            </Paper>
        </Container>
    );
}

export default OrderCreate;
