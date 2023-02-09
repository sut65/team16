import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import { ShelvingsInterface } from "../../models/methas/IShelving";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from '@mui/icons-material/Edit';
import GroupAddIcon from '@mui/icons-material/GroupAdd';
import { Dialog, DialogTitle } from "@mui/material";

function Shelving() {
 const [shelving, setShelving] = React.useState<ShelvingsInterface[]>([]);
 const [shelvingID, setShelvingID] = React.useState(0);
 const [openDelete, setOpenDelete] = React.useState(false);
 const [openUpdate, setOpenUpdate] = React.useState(false);

 const getShelving = async () => {
  const apiUrl = "http://localhost:8080/shelvings";
  const requestOptions = {
      method: "GET",
      headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
      },
  };

  await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
          if (res.data) {
              console.log(res.data)
              setShelving(res.data);
          }
          else { console.log("NO DATA") }
      });
};

 const deleteShelf = async () => {
  const apiUrl = `http://localhost:8080/shelvings/${shelvingID}`;
  const requestOptions = {
      method: "DELETE",
      headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
      },
  };
  await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
          if (res.data) {
              console.log("delete ID: " + shelvingID)
          }
          else { console.log("NO DATA") }
      });
  handleClose();
  getShelving();
}

const handleRowClick: GridEventListener<'rowClick'> = (params) => {
  setShelvingID(Number(params.row.ID));
  localStorage.setItem("shelvingID", params.row.ID);
};

const handleClose = () => {
  setOpenDelete(false)
  setOpenUpdate(false)
};

 const columns1: GridColDef[] = [
   { field: "ID", headerName: "ID", width: 50,  headerAlign:"center", align:"center" },
   { field: "Stock", headerName: "Name", width: 150, headerAlign:"center", align:"center" ,valueFormatter: (params)=>params.value.Name },
   { field: "Number", headerName: "Amount", width: 150, headerAlign:"center", align:"center" },
   { field: "Cost", headerName: "Price", width: 150, headerAlign:"center", align:"center" },
   { field: "Label", headerName: "Label", width: 150, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.Name },
   { field: "Edit", headerName: "Edit", width: 120,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpenUpdate(true)}
                  startIcon={<EditIcon />}
              > </Button>
              );},},
    { field: "Delete", headerName: "Delete", width: 120,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpenDelete(true)}
                  startIcon={<DeleteIcon />}
              > </Button>
              );},},
 ];

 //2
 const columns2: GridColDef[] = [
  { field: "ID", headerName: "ID", width: 50,  headerAlign:"center", align:"center" },
  { field: "Stock", headerName: "Name", width: 150, headerAlign:"center", align:"center" ,valueFormatter: (params)=>params.value.Name },
  { field: "Number", headerName: "Amount", width: 150, headerAlign:"center", align:"center" },
  { field: "Cost", headerName: "Price", width: 150, headerAlign:"center", align:"center" },
  { field: "Label", headerName: "Label", width: 150, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.Name },
  { field: "Edit", headerName: "Edit", width: 120,
     renderCell: () => {
         return (
             <Button
                 variant="contained"
                 color="primary"
                 onClick={() => setOpenUpdate(true)}
                 startIcon={<EditIcon />}
             > </Button>
             );},},
   { field: "Delete", headerName: "Delete", width: 120,
     renderCell: () => {
         return (
             <Button
                 variant="contained"
                 color="primary"
                 onClick={() => setOpenDelete(true)}
                 startIcon={<DeleteIcon />}
             > </Button>
             );},},
];

//3
const columns3: GridColDef[] = [
  { field: "ID", headerName: "ID", width: 50,  headerAlign:"center", align:"center" },
  { field: "Stock", headerName: "Name", width: 150, headerAlign:"center", align:"center" ,valueFormatter: (params)=>params.value.Name },
  { field: "Number", headerName: "Amount", width: 150, headerAlign:"center", align:"center" },
  { field: "Cost", headerName: "Price", width: 150, headerAlign:"center", align:"center" },
  { field: "Label", headerName: "Label", width: 150, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.Name },
  { field: "Edit", headerName: "Edit", width: 120,
     renderCell: () => {
         return (
             <Button
                 variant="contained"
                 color="primary"
                 onClick={() => setOpenUpdate(true)}
                 startIcon={<EditIcon />}
             > </Button>
             );},},
   { field: "Delete", headerName: "Delete", width: 120,
     renderCell: () => {
         return (
             <Button
                 variant="contained"
                 color="primary"
                 onClick={() => setOpenDelete(true)}
                 startIcon={<DeleteIcon />}
             > </Button>
             );},},
];

 useEffect(() => {
   getShelving();
 }, []);

 return (

  <div>
  <Dialog open={openDelete} onClose={handleClose} >
           <DialogTitle><div className="good-font">ยืนยันการลบชั้นวางนี้</div></DialogTitle>
           <Button
                   variant="contained"
                   color="primary"
                   onClick={deleteShelf}
               >
                   <div className="good-font">
                       ยืนยัน
                   </div>
               </Button>
       </Dialog>

       <Dialog open={openUpdate} onClose={handleClose} >
           <DialogTitle><div className="good-font">ยืนยันการแก้ไขชั้นวางนี้</div></DialogTitle>
           <Button
                   variant="contained"
                   color="primary"

                   component={RouterLink}
                   to="/ShelvingUpdate"
               >
                   <div className="good-font">
                       ยืนยัน
                   </div>
               </Button>
       </Dialog>
     <Container maxWidth="md">
       <Box
         display="flex"
         sx={{ marginTop: 2,}}
       >

         <Box flexGrow={1}>
           <Button
             variant="contained"
             color="primary"
           >
             <div className="good-font">Meat</div>
           </Button>
         </Box>

         <Box sx={{ paddingX: 1, paddingY: 0 }}>
           <Button
             component={RouterLink}
             to="/ShelvingCreate"
             variant="contained"
             color="primary"
             startIcon={<GroupAddIcon />}
           >
             <div className="good-font">Add Meat to the shelf</div>
           </Button>
         </Box>

       </Box>
       <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={shelving}
                        getRowId={(row) => row.ID}
                        columns={columns1}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                        onRowClick={handleRowClick}
                    />
                </div>
     </Container>


     <Container maxWidth="md">
       <Box
         display="flex"
         sx={{ marginTop: 2,}}
       >
         <Box flexGrow={1}>
           <Button
             variant="contained"
             color="primary"
           >
             <div className="good-font">Fresh</div>
           </Button>
         </Box>

         <Box sx={{ paddingX: 1, paddingY: 0 }}>
           <Button
             component={RouterLink}
             to="/ShelvingCreate"
             variant="contained"
             color="primary"
             startIcon={<GroupAddIcon />}
           >
             <div className="good-font">Add Fresh to the shelf</div>
           </Button>
         </Box>

       </Box>
       <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={shelving}
                        getRowId={(row) => row.ID}
                        columns={columns2}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                        onRowClick={handleRowClick}
                    />
                </div>
     </Container>

     <Container maxWidth="md">
       <Box
         display="flex"
         sx={{ marginTop: 2,}}
       >
         <Box flexGrow={1}>
           <Button
             variant="contained"
             color="primary"
           >
             <div className="good-font">Dairy</div>
           </Button>
         </Box>

         <Box sx={{ paddingX: 1, paddingY: 0 }}>
           <Button
             component={RouterLink}
             to="/ShelvingCreate"
             variant="contained"
             color="primary"
             startIcon={<GroupAddIcon />}
           >
             <div className="good-font">Add Dairy to the shelf</div>
           </Button>
         </Box>

       </Box>
       <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={shelving}
                        getRowId={(row) => row.ID}
                        columns={columns3}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                        onRowClick={handleRowClick}
                    />
                </div>
     </Container>
   </div>
 );
}
export default Shelving;
