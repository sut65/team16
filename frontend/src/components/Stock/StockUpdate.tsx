import { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { StocksInterface } from "../../models/methas/IStock";
import { KindsInterface } from "../../models/methas/IKind";
import { StoragesInterface } from "../../models/methas/IStorage";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import React from "react";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { TextField } from "@mui/material";
import { GetCurrentEmployee } from "../../services/HttpClientService";
import { EmployeeInterface } from "../../models/IEmployee";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
 ) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
 });



function StockUpdate() {
  const [date, setDate] = React.useState<Date | null>(null);
  const [stock, setStock] = React.useState<Partial<StocksInterface>>({});
  const [kind, setKind] = React.useState<KindsInterface[]>([]);
  const [errorMessage, setErrorMessage] = React.useState("");
  const [employee, setEmployee] = React.useState<EmployeeInterface>();
  const [storage, setStorage] = React.useState<StoragesInterface[]>([]);
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
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
    const id = event.target.id as keyof typeof StockUpdate;
    const { value } = event.target;
    setStock({ ...stock, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof stock;
    console.log(name, event.target.value);
    setStock({
      ...stock,
      [name]: event.target.value,
    });
  };

  const getKind = async () => {
    fetch(`${apiUrl}/kinds`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setKind(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const getStorage = async () => {
    fetch(`${apiUrl}/storages`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setStorage(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEmployee = async () => {
    let res = await GetCurrentEmployee();
    stock.Employee_ID = res.ID;
    if (res) {
        setEmployee(res);
        console.log(res)
    }
};


  useEffect(() => {
    getStorage();
    getKind();
    getEmployee();
  }, []);

  let stockID = localStorage.getItem("stockID"); // เรีกใช้ค่าจากlocal storage 

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      Name: stock.Name,
      Amount: convertType(stock.Amount),
      Price: convertType(stock.Price),
      Kind_ID: convertType(stock.Kind_ID),
      Storage_ID: convertType(stock.Storage_ID),
      Employee_ID: convertType(stock.Employee_ID),
      DateTime: date,
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

        fetch(`${apiUrl}/stocks/${stockID}`, requestOptions) // แนบIDไปด้วย
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
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
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
             อัปเดตสินค้า ID : {stockID}
            </div>
           </Typography>
         </Box>
       </Box>
       <Divider />
       <Grid container spacing={3} sx={{ padding: 2 }}>

<Grid item xs={6}>
  <p>Name</p>
  <FormControl fullWidth variant="outlined">
    <TextField
      id="Name"
      variant="outlined"
      type="string"
      size="medium"
      value={stock.Name || ""}
      onChange={handleInputChange}
    />
  </FormControl>
</Grid>
<Grid item xs={6}>
  <FormControl fullWidth variant="outlined">
    <p>Amount</p>
    <TextField
      id="Amount"
      variant="outlined"
      type="number"
      size="medium"
      value={stock.Amount || ""}
      onChange={handleInputChange}
    />
  </FormControl>
</Grid>
<Grid item xs={6}>
  <FormControl fullWidth variant="outlined">
    <p>Price</p>
    <TextField
      id="Price"
      variant="outlined"
      type="number"
      size="medium"
      value={stock.Price || ""}
      onChange={handleInputChange}
    />
  </FormControl>
</Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Kind</p>
              <Select
                native
                value={stock.Kind_ID+""}
                onChange={handleChange}
                inputProps={{
                  name: "Kind_ID",
                }}
              >
                <option aria-label="None" value="">
                </option>
                {kind.map((item: KindsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Storage</p>
              <Select
                native
                value={stock.Storage_ID+""}
                onChange={handleChange}
                inputProps={{
                  name: "Storage_ID",
                }}
              >
                <option aria-label="None" value="">
                </option>
                {storage.map((item: StoragesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>DateTime</p>
             <LocalizationProvider dateAdapter={AdapterDateFns}>
               <DatePicker
                 value={date}
                 onChange={(newValue) => {
                   setDate(newValue);
                 }}
                 renderInput={(params) => <TextField {...params} />}
               />
             </LocalizationProvider>
           </FormControl>
         </Grid>
         <Grid item xs={12}>
                        <FormControl fullWidth variant="outlined">
                            <p className="good-font">Employee</p>
                            <Select
                                native
                                value={stock.Employee_ID + ""}
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
              component={RouterLink}
              to="/Stock"
              variant="contained"
            >
              Back
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
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


export default StockUpdate;
