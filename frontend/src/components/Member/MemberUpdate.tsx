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
import { MemberInterface } from "../../models/theerawat/IMember";
import { EmployeeInterface } from "../../models/IEmployee";
import { GenderInterface } from "../../models/theerawat/IGender";
import { LevelInterface } from "../../models/theerawat/ILevel";
import { Autocomplete, Select, SelectChangeEvent } from "@mui/material";
import { GetCurrentEmployee } from "../../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
 props,
 ref
) {
 return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function MemberUpdate() {
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const [errorMessage, setErrorMessage] = React.useState("");

 const [member, setMember] = React.useState<Partial<MemberInterface>>({});
 const [employee, setEmployee] = React.useState<EmployeeInterface>();
 const [gender, setGender] = React.useState<GenderInterface[]>([]);
 const [level, setLevel] = React.useState<LevelInterface[]>([]);

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
   const id = event.target.id as keyof typeof MemberUpdate;
   const { value } = event.target;
   setMember({ ...member, [id]: value });
 };

 const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof member;
    setMember({
        ...member,
        [name]: event.target.value,
    });
};

 const apiUrl = "http://localhost:8080";

 const getGender = async () => {
    fetch(`${apiUrl}/genders`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setGender(res.data);
            }
            else { console.log("NO DATA") }
        });
};

const getLevel = async () => {
    fetch(`${apiUrl}/levels`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setLevel(res.data);
            }
            else { console.log("NO DATA") }
        });
};

const getEmployee = async () => {
    let res = await GetCurrentEmployee();
    member.Employee_ID = res.ID;
    if (res) {
        setEmployee(res);
        console.log(res)
    }
};

useEffect(() => {
    getGender();
    getLevel();
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

let Member_ID = localStorage.getItem("Member_ID");

async function submit() {
   let data = {
     Mem_Name: member.Mem_Name ?? "",
     Mem_Age: typeof member.Mem_Age === "string" ? parseInt(member.Mem_Age) : 0,
     Mem_Tel: member.Mem_Tel ?? "",

     Gender_ID: convertType(member.Gender_ID),
     Level_ID: convertType(member.Level_ID),
     Employee_ID: convertType(member.Employee_ID),
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

   let res = await fetch(`${apiUrl}/members/${Member_ID}`, requestOptions)
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
        setAlertMessage("แก้ไขข้อมูลสำเร็จ");
        setSuccess(true);
      } else {
        setAlertMessage(res.message);
        setError(true);
      }
}

 return (

   <Container maxWidth="md">
      <Snackbar
       id="success" open={success} autoHideDuration={6000} onClose={handleClose}
       anchorOrigin={{ vertical: "bottom", horizontal: "center" }} >
       <Alert onClose={handleClose} severity="success">
       <div className="good-font">
                    {message}
       </div>
       </Alert>
     </Snackbar>
     <Snackbar 
        id="error" open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
       <div className="good-font">
                    {message}
       </div>
       </Alert>
     </Snackbar>
     <Paper>

       <Box
         display="flex"
         sx={{ marginTop: 2,backgroundColor: 'lavender', }}
       >
         <Box sx={{ paddingX: 2, paddingY: 1 }}>
           <Typography
             component="h2"
             variant="h6"
             color="primary"
             gutterBottom
           >
             <div className="good-font-big">แก้ไขรายชื่อสมาชิก ID : {Member_ID} </div>
           </Typography>
         </Box>
       </Box>

       <Divider />
       <Grid container spacing={3} sx={{ padding: 2 }}>
         <Grid item xs={8}>
           <p className="good-font">ชื่อ - นามสกุล</p>
           <FormControl fullWidth variant="outlined">
             <TextField
               id="Mem_Name"
               variant="outlined"
               type="string"
               size="medium"
               value={member.Mem_Name || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

         <Grid item xs={4}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">อายุ</p>
             <TextField
               id="Mem_Age"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 1 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={member.Mem_Age || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

         <Grid item xs={7}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">เบอร์มือถือ</p>
             <TextField
               id="Mem_Tel"
               variant="outlined"
               type="string"
               size="medium"
               value={member.Mem_Tel || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

         <Grid item xs={5}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">เพศ</p>
                <Autocomplete
                disablePortal
                id="Gender_ID"
                getOptionLabel={(item: GenderInterface) => `${item.Gender_Name}`}
                options={gender}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { member.Gender_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Gender -" />}
                />
            </FormControl>
        </Grid>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">ระดับสมาชิก</p>
                <Autocomplete
                disablePortal
                id="Level_ID"
                getOptionLabel={(item: LevelInterface) => `${item.Level_Name}`}
                options={level}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { member.Level_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Level -" />}
                />
            </FormControl>
        </Grid>

        <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">พนักงานที่บันทึก</p>
                <Select
                    native
                    value={member.Employee_ID + ""}
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
           <Button component={RouterLink} to="/Member" variant="contained">
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
export default MemberUpdate;