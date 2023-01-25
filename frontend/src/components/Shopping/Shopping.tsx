import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { OrderInterface } from "../../models/Natthapon/IOrder";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import GroupAddIcon from '@mui/icons-material/GroupAdd';
import ShoppingBasketIcon from '@mui/icons-material/ShoppingBasket';

function Order() {
 const [order, setOrder] = React.useState<OrderInterface[]>([]);
 const getOrder = async () => {
   const apiUrl = "http://localhost:8080/Orders";
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
         setOrder (res.data);
       }
     });
 };

 const columns: GridColDef[] = [
   { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },
   { field: "Shelving?", headerName: "Shelving?", width: 100, headerAlign:"center",valueFormatter: (params)=>params.value.ID},
   { field: "Quantity", headerName: "Quantity", width: 250, headerAlign:"center" },
   { field: "Member", headerName: "Member", width: 200, headerAlign:"center" ,valueFormatter: (params)=>params.value.Mem_Name},
 ];

 useEffect(() => {
   getOrder();
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
            Order
        </Typography>
        </Box>

        <Box sx={{ paddingX: 1, paddingY: 0 }}> 
        <Button
            component={RouterLink}
            to="/Cart"
            variant="contained"
            color="primary"
            startIcon={<ShoppingBasketIcon />}
        >
        Shopping Cart 
        </Button>
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
            rows={order}
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
export default Order;