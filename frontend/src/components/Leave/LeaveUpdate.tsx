import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import TextField, { TextFieldProps } from "@mui/material/TextField";
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
import { EmployeeInterface } from "../../models/IEmployee";
import { Autocomplete, Select, SelectChangeEvent } from "@mui/material";
import { GetCurrentEmployee } from "../../services/HttpClientService";
import { SectionInterface } from "../../models/theerawat/ISection";
import { L_TypeInterface } from "../../models/theerawat/IL_Type";
import { LeaveInterface } from "../../models/theerawat/ILeave";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { DatePicker } from "@mui/x-date-pickers";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
 props,
 ref
) {
 return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function LeaveUpdate() {
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const [errorMessage, setErrorMessage] = React.useState("");

 const [leave, setLeave] = React.useState<Partial<LeaveInterface>>({
    Doc_DateS: new Date(),
    Doc_DateE: new Date(),
 });
 const [employee, setEmployee] = React.useState<EmployeeInterface>();
 const [l_type, setL_Type] = React.useState<L_TypeInterface[]>([]);
 const [section, setSection] = React.useState<SectionInterface[]>([]);

 const [message, setAlertMessage] = React.useState("");

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
   const id = event.target.id as keyof typeof LeaveUpdate;
   const { value } = event.target;
   setLeave ({ ...leave, [id]: value });
 };

 const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof leave;
    setLeave({
        ...leave,
        [name]: event.target.value,
    });
};

 const apiUrl = "http://localhost:8080";

 const getL_Type = async () => {
    fetch(`${apiUrl}/l_types`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setL_Type (res.data);
            }
            else { console.log("NO DATA") }
        });
};

const getSection = async () => {
    fetch(`${apiUrl}/sections`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setSection(res.data);
            }
            else { console.log("NO DATA") }
        });
};

const getEmployee = async () => {
    let res = await GetCurrentEmployee();
    leave.Employee_ID = res.ID;
    if (res) {
        setEmployee(res);
        console.log(res)
    }
};

useEffect(() => {
    getL_Type();
    getSection();
    getEmployee();
}, []);

const requestOptions = {
    method: "GET",
    headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json"
    },
};

const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
};

let Leave_ID = localStorage.getItem("Leave_ID");

async function submit() {
   let data = {
    Doc_Reason: leave.Doc_Reason ?? "",
    Doc_DateS:  leave.Doc_DateS,
    Doc_DateE:  leave.Doc_DateE,
    Doc_Cont: leave.Doc_Cont ?? "",

     L_Type_ID: convertType(leave.L_Type_ID),
     Section_ID: convertType(leave.Section_ID),
     Employee_ID: convertType(leave.Employee_ID),
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

   let res = await fetch(`${apiUrl}/leaves/${Leave_ID}`, requestOptions)
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
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
       <div className="good-font">
          {message}
       </div>
       </Alert>
     </Snackbar>
     <Paper>

       <Box
         display="flex"
         sx={{ marginTop: 2, }}
       >
         <Box sx={{ paddingX: 2, paddingY: 1 }}>
           <Typography
             component="h2"
             variant="h6"
             color="primary"
             gutterBottom
           >
             <div className="good-font">แก้ไขการแจ้งลา ID : {Leave_ID} </div>
           </Typography>
         </Box>
       </Box>
       
       <Divider />
       <Grid container spacing={3} sx={{ padding: 2 }}>

       <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">ประเภทการลา</p>
                <Autocomplete
                disablePortal
                id="L_Type_ID"
                getOptionLabel={(item: L_TypeInterface) => `${item.Type_Name}`}
                options={l_type}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { leave.L_Type_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Type -" />}
                />
            </FormControl>
        </Grid>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">แผนก</p>
                <Autocomplete
                disablePortal
                id="Section_ID"
                getOptionLabel={(item: SectionInterface) => `${item.Sec_Name}`}
                options={section}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { leave.Section_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Section -" />}
                />
            </FormControl>
        </Grid>

        <Grid item xs={12}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">เหตุผลการลา / รายละเอียด</p>
             <TextField
               id="Doc_Reason"
               variant="outlined"
               type="string"
               size="medium"
               value={leave.Doc_Reason || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

         <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">วันเริ่มลา</p>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                    <DatePicker
                         value={leave.Doc_DateS}
                        onChange={(newValue) => {
                            setLeave({
                                ...leave,
                                Doc_DateS: newValue,
                            });
                        }}
                        renderInput={(params: JSX.IntrinsicAttributes & TextFieldProps) => <TextField {...params} />}
                    />
                </LocalizationProvider>
            </FormControl>
        </Grid>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">วันสิ้นสุดลา</p>
                <LocalizationProvider dateAdapter={AdapterDateFns}>
                    <DatePicker
                         value={leave.Doc_DateE}
                        onChange={(newValue) => {
                            setLeave({
                                ...leave,
                                Doc_DateE: newValue,
                            });
                        }}
                        renderInput={(params: JSX.IntrinsicAttributes & TextFieldProps) => <TextField {...params} />}
                    />
                </LocalizationProvider>
            </FormControl>
        </Grid>

        <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">ช่องทางการติดต่อ</p>
             <TextField
               id="Doc_Cont"
               variant="outlined"
               type="string"
               size="medium"
               value={leave.Doc_Cont || ""}
               onChange={handleInputChange}
             />
           </FormControl>
        </Grid>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">พนักงานที่บันทึก</p>
                <Select
                    native
                    value={leave.Employee_ID + ""}
                    onChange={handleChange}
                    disabled
                    inputProps={{name: "Employee_ID",}}
                 >
                <option aria-label="None" value="">
                     Disable
                </option>
                <option value={employee?.ID} key={employee?.ID}>
                    {employee?.Name}
                </option>
                </Select>
            </FormControl>
        </Grid>

         <Grid item xs={12}>
           <Button component={RouterLink} to="/Leave" variant="contained">
           <div className="good-font-white"> ย้อนกลับ </div>
           </Button>
           <Button
             style={{ float: "right" }}
             onClick={submit}
             variant="contained"
             color="primary"
           >
            <div className="good-font-white"> ยืนยันการแก้ไขข้อมูล </div>
           </Button>
         </Grid>
         </Grid>
     </Paper>
   </Container>
 );
}
export default LeaveUpdate;