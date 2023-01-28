import React, { useEffect } from "react";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'
import { Dialog, DialogTitle } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import FormControl from "@mui/material/FormControl";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { MemberInterface } from "../../models/theerawat/IMember";
import { EmployeeInterface } from "../../models/IEmployee";
import { GenderInterface } from "../../models/theerawat/IGender";
import { LevelInterface } from "../../models/theerawat/ILevel";
import { Autocomplete, Select, SelectChangeEvent } from "@mui/material";
import { GetCurrentEmployee } from "../../services/HttpClientService";
import Typography from "@mui/material/Typography";
import dayjs, { Dayjs } from 'dayjs';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { ClockPicker } from '@mui/x-date-pickers';
import { Component } from 'react';
import { render } from 'react-dom';


import { Record_employee_leave } from "../../models/panupol/IEm_out"
import { DutyInterface } from "../../models/panupol/IDuty"
import { Working_timeInterface } from "../../models/panupol/IWorking_time"
import { OvertimeInterface } from "../../models/panupol/IOvertime"

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
   ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
   });

function Record_employee_leaves_update() {
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);
    const [errorMessage, setErrorMessage] = React.useState("")

    const [employee, setEmployee] = React.useState<EmployeeInterface>();
    const [duty, setduty] = React.useState<DutyInterface[]>([]);
    const [Working_time, setWorking_time] = React.useState<Working_timeInterface[]>([]);
    const [Overtime, setOvertime] = React.useState<OvertimeInterface[]>([]);
    const [Em_Out, setEm_Out] = React.useState<Partial<Record_employee_leave>>({
      Time_Out: new Date(), Status_ID: false, 
    });

  

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
        const id = event.target.id as keyof typeof Record_employee_leaves_update;
        const { value } = event.target;
        setEm_Out({ ...Em_Out, [id]: value });
      };
     
      const handleChange = (event: SelectChangeEvent) => {
         const name = event.target.name as keyof typeof Em_Out;
         setEm_Out({
             ...Em_Out,
             [name]: event.target.value,
         });
     };

     const requestOptions = {
      method: "GET",
      headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json"
      },
  };

     const apiUrl = "http://localhost:8080"

     const getduty = async () => {
        fetch(`${apiUrl}/listduty`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setduty(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

     
      
      const getworking_time = async () => {
        fetch(`${apiUrl}/listworking_time`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setWorking_time(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getovertime = async () => {
        fetch(`${apiUrl}/listovertime`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setOvertime(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    const getEmployee = async () => {
      let res = await GetCurrentEmployee();
      Em_Out.Employee_ID = res.ID;
      if (res) {
          setEmployee(res);
          console.log(res)
      }
  };

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
};


    let employee_leavesID = localStorage.getItem("Employee_leaveID"); // เรีกใช้ค่าจากlocal storage 

    useEffect(() => {
      let date = new Date();
      console.log(date.toLocaleString());
      getduty();
      getovertime();
      getworking_time();
      getEmployee();
  }, []);

  async function submit() {
    let date = new Date();
    let data = {
      Employee_ID: convertType(Em_Out.Employee_ID),
      Duty_ID: convertType(Em_Out.Duty_ID),
      Working_time_ID: convertType(Em_Out.Working_time_ID),
      Overtime_ID: convertType(Em_Out.Overtime_ID),
      Time_Out: new Date() ,
      Number_Em: Em_Out.Number_Em ?? "",
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

  fetch(`${apiUrl}/record_employee_leaves/${employee_leavesID}`, requestOptions)  // แนบIDไปด้วย
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
      <div>
        <Container maxWidth="md">
          <Snackbar
            open={success}
            autoHideDuration={6000}
            onClose={handleClose}
            anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
          >
            <Alert onClose={handleClose} severity="success">
              บันทึกข้อมูลสำเร็จ
            </Alert>
          </Snackbar>
          <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              บันทึกข้อมูลไม่สำเร็จ
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
                  แก้ไขรายการออกงานออกงาน
                </Typography>
              </Box>
            </Box>
     
            <Divider />
            <Grid container spacing={5} sx={{ padding: 2 }}>  
            <Grid item xs={12}>
                 <FormControl fullWidth variant="outlined">
                     <p className="good-font">พนักงานที่บันทึก</p>
                     <Select
                         native
                         value={Em_Out.Employee_ID + ""}
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

              <Grid item xs={5}>
                 <FormControl fullWidth variant="outlined">
                     <p className="good-font">หน้าที่ในการทำงาน</p>
                     <Autocomplete
                     disablePortal
                     id="Duty_ID"
                     getOptionLabel={(item: DutyInterface) => `${item.Name}`}
                     options={duty}
                     sx={{ width: 'auto' }}
                     isOptionEqualToValue={(option, value) =>
                         option.ID === value.ID}
                     onChange={(e, value) => { Em_Out.Duty_ID = value?.ID }}
                     renderInput={(params) => <TextField {...params} label="- Select duty -" />}
                     />
                 </FormControl>
             </Grid>
             <Grid item xs={7}>
                 <FormControl fullWidth variant="outlined">
                     <p className="good-font">เวลาในการทำงาน</p>
                     <Autocomplete
                     disablePortal
                     id="Working_time_ID"
                     getOptionLabel={(item: Working_timeInterface) => `${item.Name}`}
                     options={Working_time}
                     sx={{ width: 'auto' }}
                     isOptionEqualToValue={(option, value) =>
                         option.ID === value.ID}
                     onChange={(e, value) => { Em_Out.Working_time_ID = value?.ID }}
                     renderInput={(params) => <TextField {...params} label="- Select working time -" />}
                     />
                 </FormControl>
             </Grid>

             <Grid item xs={5}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">ทำงานนอกเวลา</p>
                <Autocomplete
                disablePortal
                id="Overtime_ID"
                getOptionLabel={(item: OvertimeInterface) => `${item.Name}`}
                options={Overtime}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { Em_Out.Overtime_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select overtime -" />}
                />
            </FormControl>
        </Grid>
        
              <Grid item xs={7}>
                <FormControl fullWidth variant="outlined">
                  <p className="good-font">เบอร์มือถือ</p>
                  <TextField
                    id="Number_Em"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={Em_Out.Number_Em || ""}
                    onChange={handleInputChange}
                  />
                </FormControl>
              </Grid>
     
              <Grid item xs={12}>
                <Button component={RouterLink} to="/EmployeeattemdanceOUT" variant="contained">
                  Back
                </Button>
                <Button
                  style={{ float: "right" }}
                  onClick={submit}
                  variant="contained"
                  color="primary"
                >
                  Submit
                </Button>
              </Grid>
            </Grid>
     
          </Paper>
        </Container>
        </div>
      );
    
}
export default Record_employee_leaves_update;


