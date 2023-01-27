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

import { ReasonInterface } from "../../models/apisit/IReason";
import { EmployeeInterface } from "../../models/IEmployee";
import { IShelving } from "../../models/methas/IShelving";
import { SeparationInterface } from "../../models/apisit/ISeparation";
// import { GetCurrentAdmin } from "../services/HttpClientService";
import Autocomplete from "@mui/material/Autocomplete";
import { GetCurrentEmployee } from "../../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function SeparationCreate() {
    const [date, setDate] = React.useState<Date | null>(null);
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");

    const [emp, setEmployee] = React.useState<EmployeeInterface>();  
    const [reas, setReason] = React.useState<ReasonInterface[]>([]);
    const [shelf, setShelf] = React.useState<IShelving[]>([]);
    const [sep, setSeparation] = React.useState<SeparationInterface>({
        Date_Out: new Date(),
     });
    
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
        method: "GET",
        headers: { 
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json" },
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
        const id = event.target.id as keyof typeof SeparationCreate;
        const { value } = event.target;
        setSeparation({ ...sep, [id]: value });
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof sep;
        setSeparation({
            ...sep,
            [name]: event.target.value,
        });
    };

    const getReason = async () => {
        fetch(`${apiUrl}/reasons`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setReason(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getShelf = async () => {
        fetch(`${apiUrl}/Shelving`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setShelf(res.data);
                }
                else { console.log("NO DATA") }
            });
    };
    
    // const getEmployee = async () => {
    //     fetch(`${apiUrl}/emloyees`, requestOptions)
    //         .then((response) => response.json())
    //         .then((res) => {
    //             if (res.data) {
    //                 console.log(res.data)
    //                 setEmployee(res.data);
    //             }
    //             else { console.log("NO DATA") }
    //         });
    // };

    const getEmployee = async () => {
        let res = await GetCurrentEmployee();
        sep.Employee_ID = res.ID;
        if (res) {
            setEmployee(res);
            console.log(res)
        }
    };

    // const getAdmin = async () => {
    //     let res = await GetCurrentAdmin();
    //     activityHis.ADMIN_ID = res.ID;
    //     if (res) {
    //         setAdmin(res);
    //         console.log(res)
    //     }
    // };

    useEffect(() => {
        getReason();
        getShelf();
        getEmployee();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            Amount: typeof sep.Amount === "string" ? parseInt(sep.Amount) : 0,
            Date_Out: sep.Date_Out,
            Status: sep.Status ?? "",
            Employee_ID: convertType(sep.Employee_ID),
            Reason_ID: convertType(sep.Reason_ID),
            Shelving_ID: convertType(sep.Shelving_ID),
        };

        console.log(data)

        const requestOptions = {
            method: "POST",
            headers: { 
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json" },
            body: JSON.stringify(data),
        };

        fetch(`${apiUrl}/separations`, requestOptions)
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
                    การเพิ่มข้อเสร็จสิ้น
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    การเพิ่มข้อมูลไม่สำเร็จ
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
                            บันทึกรายการจำหน่ายสินค้า
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>

                <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>เหตุผล</p>
                            
                            <Autocomplete
                                disablePortal
                                id="Reason_ID"
                                getOptionLabel={(item: ReasonInterface) => `${item.Cuase}`}
                                options={reas}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { sep.Reason_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกเหตุผลที่จำหน่าย" />}
                            />
                        </FormControl>
                    </Grid>

                    {/* <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>พนักงาน</p>
                            <Select
                                native
                                value={sep.Employee_ID + ""}
                                onChange={handleChange}
                                disabled
                                inputProps={{
                                    name: "Employee_ID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    --เลือกพนักงาน--
                                </option>
                                <option value={emp?.ID} key={emp?.ID}>
                                    {emp?.Name}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid> */}

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">พนักงาน</p>
                            <Select
                                native
                                value={sep.Employee_ID + ""}
                                onChange={handleChange}
                                disabled
                                inputProps={{
                                    name: "Employee_ID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    เลือกพนักงาน
                                </option>
                                <option value={emp?.ID} key={emp?.ID}>
                                    {emp?.Name}
                                </option>
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>สินค้า</p>
                            
                            <Autocomplete
                                disablePortal
                                id="Shelving_ID"
                                getOptionLabel={(item: IShelving) => `${item.ID}`}
                                options={shelf}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { sep.Shelving_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกชั้นวาง" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">วันจำหน่าย</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={sep.Date_Out}
                                    onChange={(newValue) => {
                                        setSeparation({
                                            ...sep,
                                            Date_Out: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>จำนวน</p>
                            <TextField
                                id="Amount"
                                variant="outlined"
                                type="number"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={sep.Amount || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">สถานะ</p>
                            <TextField
                                id="Status"
                                variant="outlined"
                                type="string"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={sep.Status || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    
                    

                    <Grid item xs={12}>
                        <Button color="primary" component={RouterLink} to="/SeparationShow" variant="contained">
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

export default SeparationCreate;