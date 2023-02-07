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
import { ShelvingsInterface } from "../../models/methas/IShelving";
import { EmployeeInterface } from "../../models/IEmployee";
import { StocksInterface } from "../../models/methas/IStock";
import { LabelsInterface } from "../../models/methas/ILabel";
import { Autocomplete, Select, SelectChangeEvent } from "@mui/material";
import { GetCurrentEmployee } from "../../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
 props,
 ref
) {
 return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function ShelvingCreate() {
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const [errorMessage, setErrorMessage] = React.useState("");
 const [shelving, setShelving] = React.useState<Partial<ShelvingsInterface>>({});
 const [employee, setEmployee] = React.useState<EmployeeInterface>();
 const [stock, setStock] = React.useState<StocksInterface[]>([]);
 const [label, setLabel] = React.useState<LabelsInterface[]>([]);

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
   const id = event.target.id as keyof typeof ShelvingCreate;
   const { value } = event.target;
   setShelving({ ...shelving, [id]: value });
 };

 const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof shelving;
    console.log(name, event.target.value);
    setShelving({
        ...shelving,
        [name]: event.target.value,
    });
};

 const apiUrl = "http://localhost:8080";

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

const getLabel = async () => {
    fetch(`${apiUrl}/labels`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setLabel(res.data);
            }
            else { console.log("NO DATA") }
        });
};

const getEmployee = async () => {
    let res = await GetCurrentEmployee();
    shelving.Employee_ID = res.ID;
    if (res) {
        setEmployee(res);
        console.log(res)
    }
};

useEffect(() => {
    getStock();
    getLabel();
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

async function submit() {
   let data = {
     Amount: convertType(shelving.Amount),
     Price: convertType(shelving.Price),
     Stock_ID: convertType(shelving.Stock_ID),
     Label_ID: convertType(shelving.Label_ID),
     Employee_ID: convertType(shelving.Employee_ID),
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

   fetch(`${apiUrl}/shelves`, requestOptions)
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
         บันทึกข้อมูลสำเร็จ
       </Alert>
     </Snackbar>
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
         บันทึกข้อมูลไม่สำเร็จ {errorMessage}
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
             <div className="good-font">Add The Shelf</div>
           </Typography>
         </Box>
       </Box>

       <Divider />
       <Grid container spacing={3} sx={{ padding: 2 }}>
       <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">Name</p>
                <Autocomplete
                disablePortal
                id="Stock_ID"
                getOptionLabel={(item: StocksInterface) => `${item.Name}`}
                options={stock}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { shelving.Stock_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Name -" />}
                />
            </FormControl>
        </Grid>

         <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">Label</p>
                <Autocomplete
                disablePortal
                id="Label_ID"
                getOptionLabel={(item: LabelsInterface) => `${item.Name}`}
                options={label}
                sx={{ width: 'auto' }}
                isOptionEqualToValue={(option, value) =>
                    option.ID === value.ID}
                onChange={(e, value) => { shelving.Label_ID = value?.ID }}
                renderInput={(params) => <TextField {...params} label="- Select Label -" />}
                />
            </FormControl>
        </Grid>

        <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">Amount</p>
             <TextField
               id="Amount"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 1 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={shelving.Amount || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>


        <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p className="good-font">Price</p>
             <TextField
               id="Price"
               variant="outlined"
               type="number"
               size="medium"
               InputProps={{ inputProps: { min: 1 } }}
               InputLabelProps={{
                 shrink: true,
               }}
               value={shelving.Price || ""}
               onChange={handleInputChange}
             />
           </FormControl>
         </Grid>

        <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
                <p className="good-font">Employee</p>
                <Select
                    native
                    value={shelving.Employee_ID + ""}
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
           <Button component={RouterLink} to="/Shelving" variant="contained">
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
 );
}
export default ShelvingCreate;
