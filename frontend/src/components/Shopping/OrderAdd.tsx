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
import Autocomplete from "@mui/material/Autocomplete";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { OrderInterface } from "../../models/Natthapon/IOrder";
import { ShelvingsInterface } from "../../models/methas/IShelving";
import PaymentIcon from '@mui/icons-material/Payment';



const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function OrderCreate() {

    const [success2, setSuccess2] = React.useState(false);
    const [error2, setError2] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");
    const [message, setAlertMessage] = React.useState("");

    const [shelving, setShelving] = React.useState<ShelvingsInterface[]>([]);
    const [order, setOrder] = React.useState<OrderInterface>({});

    const [num, setNum] = React.useState(0);                    //จำนวนสินค้า input
    const [sumprice, setSumprice] = React.useState(0);          //รวมราคาในตะกร้า
    const [shevID, setShevID] = React.useState(0);              //ID ชั้นวาง
    const [amounts, setAmounts] = React.useState(0);            //จำนวนสินค้าที่ชั้นวาง
    const [shevprice, setShevprice] = React.useState(0);        //ราคาสินค้าที่ชั้นวาง


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
        setSuccess2(false);
        setError2(false);
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof OrderCreate;
        const { value } = event.target;
        setOrder({ ...order, [id]: value, Prices: value * shevprice });
        setNum(value)  
    };
    let total = order.Prices

    const handleInputPrice = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof OrderCreate;
        const { value } = event.target;
        setOrder({ ...order, [id]: value });
        console.log("Price: " + total);
        console.log("carttotal: " + sumprice);
        console.log("sum: " + (Number(sumprice) + Number(total)));
    };

    const getShelving = async () => {
        fetch(`${apiUrl}/shelv`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setShelving(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    let cartID = localStorage.getItem("cartID");

    useEffect(() => {
        getShelving();
    }, []);

    fetch(`${apiUrl}/ordersum/${cartID}`, requestOptions)
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

    async function reduce() {
        let quantity = amounts - num;
        let data = {
            Number: quantity,
            //Cost: shevprice,
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

        let res = await fetch(`${apiUrl}/UpdateQuantity/${shevID}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess2(true);
                    setErrorMessage("")
                    return { status: true, message: res.data };
                } else {
                    setError2(true);
                    setErrorMessage(res.error)
                    return { status: false, message: res.error };
                }
            });
        if (res.status) {
            setAlertMessage("เพิ่มสินค้าลงตะกร้าแล้ว");
            setSuccess2(true);
        } else {
            setAlertMessage(res.message);
            setError2(true);
        }
    }


    async function addproduct() {
        let data = {
            Quantity: typeof order.Quantity === "string" ? parseFloat(order.Quantity) : 0,
            Prices: typeof order.Prices === "string" ? parseFloat(order.Prices) : total,
            Shelving_ID: convertType(order.Shelving_ID),
            Shopping_Cart_ID: Number(cartID),
        };

        console.log(data)

        const requestOptions = {
            method: "Post",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        let res = await fetch(`${apiUrl}/orders`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    setSuccess2(true);
                    setErrorMessage("")
                    reduce()
                    sum()
                    return { status: true, message: res.data };
                } else {
                    setError2(true);
                    setErrorMessage(res.error)
                    return { status: false, message: res.error };
                }
            });
        if (res.status) {
            setAlertMessage("เพิ่มสินค้าลงตะกร้าแล้ว");
            setSuccess2(true);
        } else {
            setAlertMessage(res.message);
            setError2(true);
        }
    }

    async function sum() {
        let price = Number(sumprice) + Number(total)
        let data = {
            Total: price,
            Status_ID: 1,
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

        fetch(`${apiUrl}/cart/${cartID}`, requestOptions)
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
                open={success2}
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
            <Snackbar open={error2}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
                <Alert onClose={handleClose} severity="error">
                    <div className="good-font">
                        {message}
                    </div>
                </Alert>
            </Snackbar>

            <Paper>
                <Box display="flex" sx={{ marginTop: 2, paddingX: 2, paddingY: 1 }}>
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            เพิ่มรายการสินค้า
                        </Typography>
                    </Box>

                    <Box sx={{ paddingX: 1, paddingY: 0 }}>
                        <Button
                            component={RouterLink}
                            to="/Cart"
                            variant="contained"
                            color="primary"
                            startIcon={<ArrowBackIcon />}
                        >
                            กลับ
                        </Button>
                    </Box>

                    <Box sx={{ paddingX: 1, paddingY: 0 }}>
                        <Button
                            component={RouterLink}
                            to="/Pay"
                            variant="contained"
                            color="primary"
                            startIcon={<PaymentIcon />}
                        >
                            ชำระสินค้า
                        </Button>
                    </Box>

                </Box>
                <Divider />

                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รายการสินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Shelving_ID"
                                getOptionLabel={(item: ShelvingsInterface) => {
                                    if (item && item.Stock) {
                                      return `${item.Stock.Name} ราคา ${item.Cost}`;
                                    }
                                    return '';
                                  }}
                                  
                                options={shelving}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) => option.ID === value.ID}
                                onChange={(e, value) => {
                                    order.Shelving_ID = value?.ID;
                                    if (value) {
                                        setAmounts(value.Number)
                                        setShevID(value.ID)
                                        setShevprice(value.Cost)
                                    };
                                    console.log("shevID: " + shevID);
                                    console.log("Amount: " + amounts);
                                    console.log("shevprice: " + shevprice);
                                }}
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
                                InputProps={{ inputProps: { min: 0, max: amounts-1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={order.Quantity || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={3}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">รวมราคา</p>
                            <TextField
                                id="Prices"
                                variant="outlined"
                                type="number"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={order.Prices || ""}
                                onChange={handleInputPrice}
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

        </Container>
    );
}

export default OrderCreate;
