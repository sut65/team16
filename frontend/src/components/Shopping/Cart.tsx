import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { CartInterface } from "../../models/Natthapon/ICart";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import GroupAddIcon from '@mui/icons-material/GroupAdd';

function Cart() {
 const [cart, setCart] = React.useState<CartInterface[]>([]);
 const getCart = async () => {
   const apiUrl = "http://localhost:8080/carts";
   const requestOptions = {
     method: "GET",
     headers: {Authorization: `Bearer ${localStorage.getItem("token")}`,
     "Content-Type": "application/json",},
   };

   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
         setCart (res.data);
       }
     });
 };

 const columns: GridColDef[] = [
   { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },
   { field: "Total", headerName: "Total", width: 80,  headerAlign:"center" },
   { field: "Status", headerName: "Status", width: 100, headerAlign:"center",valueFormatter: (params)=>params.value.Status},
   { field: "Member", headerName: "Member", width: 200, headerAlign:"center" ,valueFormatter: (params)=>params.value.Mem_Name},
   { field: "Employee", headerName: "Employee", width: 200, headerAlign:"center" ,valueFormatter: (params)=>params.value.Name},
 ];

 useEffect(() => {
   getCart();
 }, []);

 return (

    <div>
        <Container maxWidth="md">
        <Box
        display="flex"
        sx={{ marginTop: 2,}}
        >
        <Box flexGrow={1}>
        <Typography
            component="h2"
            variant="h6"
            color="primary"
            gutterBottom
        >
            Shopping Cart
        </Typography>
        </Box>

        <Box sx={{ paddingX: 1, paddingY: 0 }}> 
        <Button
            variant="contained"
            color="primary"
            startIcon={<DeleteIcon />}
        >
        Delete 
        </Button>
        </Box>

        <Box sx={{ paddingX: 1, paddingY: 0 }}> 
        <Button
            variant="contained"
            color="primary"
            startIcon={<EditIcon />}
        >
        Update 
        </Button>
        </Box>

        <Box sx={{ paddingX: 1, paddingY: 0 }}>
        <Button
            component={RouterLink}
            to="/OrderCreate"
            variant="contained"
            color="primary"
            startIcon={<GroupAddIcon />}
        >
            Create
        </Button>
        </Box>

        </Box>
        
          <div style={{ height: 400, width: "100%", marginTop: '20px'}}>
              <DataGrid
              rows={cart}
              getRowId={(row) => row.ID}
              columns={columns}
              pageSize={5}
              rowsPerPageOptions={[5]}
              />
          </div>
        </Container>
    </div>
 );
}
export default Cart;