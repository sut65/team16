import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import Typography from "@mui/material/Typography";

import Button from "@mui/material/Button";

import Container from "@mui/material/Container";

import Box from "@mui/material/Box";

import { StocksInterface } from "../../models/methas/IStock";

import { DataGrid, GridColDef } from "@mui/x-data-grid";

//Stock Increase
function StockIncrease() {

 const [stocks, setStocks] = React.useState<StocksInterface[]>([]);


 const getStocks = async () => {

   const apiUrl = "http://localhost:8080/stocks";

   const requestOptions = {

     method: "GET",

     headers: { "Content-Type": "application/json" },

   };


   fetch(apiUrl, requestOptions)

     .then((response) => response.json())

     .then((res) => {

       console.log(res.data);

       if (res.data) {

         setStocks(res.data);

       }

     });

 };


 const columns: GridColDef[] = [

   { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },

   { field: "Name", headerName: "Name", width: 150, headerAlign:"center" },

   { field: "Amount", headerName: "Amount", width: 150, headerAlign:"center" },

   { field: "Price", headerName: "Price", width: 150, headerAlign:"center" },

   { field: "Kind", headerName: "Kind", valueFormatter:(params) => params.value.Name, width: 150, headerAlign:"center" },

   { field: "Storage", headerName: "Storage", valueFormatter:(params) => params.value.Name, width: 150, headerAlign:"center" },

   { field: "DateTime", headerName: "DateTime", width: 200, headerAlign:"center" },

 ];


 useEffect(() => {

   getStocks();

 }, []);


 return (

   <div>

     <Container maxWidth="md">

       <Box

         display="flex"

         sx={{

           marginTop: 2,

         }}

       >

         <Box flexGrow={1}>

           <Typography

             component="h2"

             variant="h6"

             color="primary"

             gutterBottom

           >

             Stocks

           </Typography>

         </Box>

         <Box>

           <Button

             component={RouterLink}

             to="/StockCreate"

             variant="contained"

             color="primary"

           >

             Create Stock

           </Button>

         </Box>

       </Box>

       <div style={{ height: 400, width: "100%", marginTop: '20px'}}>

         <DataGrid
           rows={stocks}

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


export default StockIncrease;