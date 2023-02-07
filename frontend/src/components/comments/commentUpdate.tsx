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

import { Type_CommentInterface } from "../../models/apisit/Type_comment";
import { Review_pointInterface } from "../../models/apisit/IReview_point";
import { PaymentInterface } from "../../models/Natthapon/IPayment";
import { CommentInterface } from "../../models/apisit/IComment";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function CommentUpdate() {

    const [date, setDate] = React.useState<Date | null>(null);
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("");
    const [message, setAlertMessage] = React.useState("");


    const [type_com, setType_Com] = React.useState<Type_CommentInterface[]>([]);
    const [review_point, setReview_Point] = React.useState<Review_pointInterface[]>([]);
    const [pay, setPayment] = React.useState<PaymentInterface[]>([]);
    const [com, setComment] = React.useState<CommentInterface>({
        Date_Now: new Date(),
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
        const id = event.target.id as keyof typeof CommentUpdate;
        const { value } = event.target;
        setComment({ ...com, [id]: value });
    };

    const handleChange = (event: SelectChangeEvent) => {
        const name = event.target.name as keyof typeof com;
        setComment({
            ...com,
            [name]: event.target.value,
        });
    };

    const getReview_Point = async () => {
        fetch(`${apiUrl}/review_points`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setReview_Point(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getType_Comment = async () => {
        fetch(`${apiUrl}/type_comments`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setType_Com(res.data);
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

    let commentID = localStorage.getItem("commentID"); // เรีกใช้ค่าจากlocal storage 

    useEffect(() => {
        getReview_Point();
        getType_Comment();
        getPayment();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            Bought_now: typeof com.Bought_now === "string" ? parseInt(com.Bought_now) : 0,
            Date_Now: com.Date_Now,
            Comments: com.Comments ?? "",
            Review_point_ID: convertType(com.Review_point_ID),
            Type_Com_ID: convertType(com.Type_Com_ID),
            Payment_ID: convertType(com.Payment_ID),
        };

        console.log(data)

        const requestOptions = {
            method: "PATCH", // ใช้ PATCH
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        let res = await fetch(`${apiUrl}/comment/${commentID}`, requestOptions) // แนบIDไปด้วย
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
                setAlertMessage("เพิ่มข้อมูลเสร็จสิ้น");
                setSuccess(true);
            } else {
                setAlertMessage(res.message);
                setError(true);
            }
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
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
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
                                แก้ไขรายการความคิดเห็น ID : {commentID}
                            </div>
                        </Typography>
                    </Box>
                </Box>
                <Divider />

                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>ประเภทความคิดเห็น</p>

                            <Autocomplete
                                disablePortal
                                id="Type_Com_ID"
                                getOptionLabel={(item: Type_CommentInterface) => `${item.Type_Com_Name}`}
                                options={type_com}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { com.Type_Com_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="เลือกประเภทความคิดเห็น" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">ความคิดเห็น</p>
                            <TextField
                                id="Comments"
                                variant="outlined"
                                type="string"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={com.Comments || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>ให้คะแนน</p>

                            <Autocomplete
                                disablePortal
                                id="Review_point_ID"
                                getOptionLabel={(item: Review_pointInterface) => `${item.Point_Name}`}
                                options={review_point}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { com.Review_point_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="ให้คะแนน" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>ใบเสร็จ</p>

                            <Autocomplete
                                disablePortal
                                id="Payment_ID"
                                getOptionLabel={(item: PaymentInterface) => `${item.ID}`}
                                options={pay}
                                sx={{ width: 'auto' }}
                                isOptionEqualToValue={(option, value) =>
                                    option.ID === value.ID}
                                onChange={(e, value) => { com.Payment_ID = value?.ID }}
                                renderInput={(params) => <TextField {...params} label="แนบใบเสร็จ" />}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">วันที่</p>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <DatePicker
                                    value={com.Date_Now}
                                    onChange={(newValue) => {
                                        setComment({
                                            ...com,
                                            Date_Now: newValue,
                                        });
                                    }}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </LocalizationProvider>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p>จำนวนที่ซื้อ</p>
                            <TextField
                                id="Bought_now"
                                variant="outlined"
                                type="number"
                                size="medium"
                                InputProps={{ inputProps: { min: 1 } }}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                value={com.Bought_now || ""}
                                onChange={handleInputChange}
                            />
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/CommentShow" variant="contained">
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
        </Container >
    );
}

export default CommentUpdate;
