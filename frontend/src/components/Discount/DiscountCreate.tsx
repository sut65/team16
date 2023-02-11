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

import { DiscountInterface } from "../../models/thanadet/IDiscount"
import { Discount_Type_Interface } from "../../models/thanadet/IDiscount_Type"
import { EmployeeInterface } from "../../models/IEmployee"
import { StocksInterface } from "../../models/methas/IStock"
import { GetCurrentEmployee } from "../../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function DiscountCreate() {
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");
    const [message, setAlertMessage] = React.useState("");
    const [salePrice, setSalePrice] = React.useState(0);
    const [disPrice, setDisPrice] = React.useState(0);
    const [stockID, setStockID] = React.useState(0);
    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [stock, setStock] = React.useState<StocksInterface[]>([]);
    const [dt, setDt] = React.useState<Discount_Type_Interface[]>([]);
    const [discount, setDiscount] = React.useState<DiscountInterface>({
        Discount_s: new Date(),
        Discount_e: new Date(),
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
        const id = event.target.id as keyof typeof DiscountCreate;
        const { value } = event.target;
        setDiscount({ ...discount, [id]: value });
        // setSalePrice(value);
        // console.log("Price: " + disPrice);
        // console.log("SALE: " + salePrice);
        // console.log("result: " + (disPrice - salePrice));
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof discount;
        setDiscount({
            ...discount,
            [name]: event.target.value,
        });
    };

    const getDiscount_Type = async () => {
        fetch(`${apiUrl}/discount_types`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setDt(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getStock = async () => {
        fetch(`${apiUrl}/stocks`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setStock(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getEmployee = async () => {
        let res = await GetCurrentEmployee();
        discount.Employee_ID = res.ID;
        if (res) {
            setEmployee(res);
            console.log(res)
        }
    };

    useEffect(() => {
        getEmployee();
        getDiscount_Type();
        getStock();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        console.log("stock ID: " + stockID)
        let data = {
            Discount_Price: typeof discount.Discount_Price === "string" ? parseInt(discount.Discount_Price) : 0,
            Discount_s: discount.Discount_s,
            Discount_e: discount.Discount_e,
            Stock_ID: convertType(discount.Stock_ID),
            Discount_Type_ID: convertType(discount.Discount_Type_ID),
            Employee_ID: convertType(discount.Employee_ID),
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
        let res = await fetch(`${apiUrl}/discounts/${stockID}`, requestOptions)
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
            discounting();
        } else {
            setAlertMessage(res.message);
            setError(true);
        }
    }

    async function discounting() {
        let stockDisID = discount.Stock_ID;
        let data = {
            Price: (disPrice - salePrice),
        };
        console.log(stockDisID)
        console.log(data)

        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        fetch(`${apiUrl}/discounting/${stockDisID}`, requestOptions)
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
                id="success"
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
            <Snackbar
                id="error"
                open={error}
                autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    <div className="good-font">
                        {message}
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
                                เพิ่มส่วนลด
                            </div>
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">ประเภทของส่วนลด</p>
                            <Autocomplete
                                disablePortal
                                id="Discount_Type_ID"
                                getOptionLabel={(item: Discount_Type_Interface) => `${item.Type_Name}`}
                                options={dt}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { discount.Discount_Type_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกประเภทของส่วนลด" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">สินค้า</p>
                            <Autocomplete
                                disablePortal
                                id="Stock_ID"
                                getOptionLabel={(item: StocksInterface) => `${item.Name} ราคา ${item.Price}`}
                                options={stock}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => {
                                    discount.Stock_ID = value?.ID;
                                    if (value) {
                                        setDisPrice(value.Price)
                                        setStockID(value.ID)
                                    };
                                    console.log("Stock Price: " + disPrice);
                                }}
                                renderInput={(params) => <TextField {...params} label="เลือกสินค้า" />}
                            />
                        </FormControl>
                    </Grid>


                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">ราคาที่ลด หน่วยเป็นบาท</p>
                            <TextField
                                id="Discount_Price"
                                variant="outlined"
                                type="number"
                                size="medium"
                                InputProps={{ inputProps: { min: 1, max: 50 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={discount.Discount_Price || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">พนักงานที่บันทึก</p>
                            <Select
                                native
                                value={discount.Employee_ID + ""}
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
                            <p className="good-font">วันที่เริ่มลดราคา</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    // disabled
                                    value={discount.Discount_s}
                                    onChange={(newValue) => {
                                        setDiscount({
                                            ...discount,
                                            Discount_s: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">วันที่สิ้นสุดการลดราคา</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={discount.Discount_e}
                                    onChange={(newValue) => {
                                        setDiscount({
                                            ...discount,
                                            Discount_e: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/Discount" variant="contained" >
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
                                บันทึก
                            </div>
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default DiscountCreate;
