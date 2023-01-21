import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { MemberInterface } from "../../models/theerawat/ILeave";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import GroupAddIcon from '@mui/icons-material/GroupAdd';

function Member() {
 const [member, setMember] = React.useState<MemberInterface[]>([]);
 const getMember = async () => {
   const apiUrl = "http://localhost:8080/member";
   const requestOptions = {
     method: "GET",
     headers: { "Content-Type": "application/json" },
   };

   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       console.log(res.data);
       if (res.data) {
         setMember (res.data);
       }
     });
 };

 const columns: GridColDef[] = [
   { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },
   { field: "Name", headerName: "Name", width: 250, headerAlign:"center" },
   { field: "Age", headerName: "Age", width: 100, headerAlign:"center" },
   { field: "Tel", headerName: "Phone number", width: 150, headerAlign:"center" },
   { field: "Gender", headerName: "Gender", width: 100, headerAlign:"center" },
   { field: "Level", headerName: "Level", width: 200, headerAlign:"center" },
 ];

 useEffect(() => {
   getMember();
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
             Member
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
             to="/MemberCreate"
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
           rows={member}
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
export default Member;