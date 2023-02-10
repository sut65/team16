import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import { MemberInterface } from "../../models/theerawat/IMember";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import GroupAddIcon from '@mui/icons-material/GroupAdd';
import { Dialog, DialogTitle } from "@mui/material";

function Member() {
 const [member, setMember] = React.useState<MemberInterface[]>([]);
 const [Member_ID, setMemberID] = React.useState(0); 
 const [openDelete, setOpendelete] = React.useState(false); 
 const [openUpdate, setOpenupdate] = React.useState(false);

 const getMember = async () => {
   const apiUrl = "http://localhost:8080/members";
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
         setMember (res.data);
       }
     });
 };

 const deleteMember = async () => {
  const apiUrl = `http://localhost:8080/members/${Member_ID}`;
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
              console.log("delete ID: " + Member_ID)
          }
          else { console.log("NO DATA") }
      });
  handleClose();
  getMember();
}

const handleRowClick: GridEventListener<'rowClick'> = (params) => {
  setMemberID(Number(params.row.ID)); 
  localStorage.setItem("Member_ID", params.row.ID); 
};

const handleClose = () => {
  setOpendelete(false)
  setOpenupdate(false)
};

 const columns: GridColDef[] = [
   { field: "ID", headerName: "ID", width: 50 },
   { field: "Mem_Name", headerName: "ชื่อ - นามสกุล", width: 250 },
   { field: "Mem_Age", headerName: "อายุ", width: 100 },
   { field: "Mem_Tel", headerName: "เบอร์โทรศัพท์", width: 200},
   { field: "Gender", headerName: "เพศ" , valueFormatter: (params)=>params.value.Gender_Name },
   { field: "Level", headerName: "ระดับสมาชิก", width: 150 , valueFormatter: (params)=>params.value.Level_Name},
   { field: "edit", headerName: "แก้ไข", width: 100,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpenupdate(true)}
                  startIcon={<EditIcon />}
              > </Button>
              );},},
    { field: "delete", headerName: "ลบ", width: 100,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpendelete(true)}
                  startIcon={<DeleteIcon />}
              > </Button>
              );},},
 ];

 useEffect(() => {
   getMember();
 }, []);

 return (

  <div>
  <Dialog open={openDelete} onClose={handleClose} >
           <DialogTitle><div className="good-font">ยืนยันการลบสมาชิกนี้</div></DialogTitle>
           <Button
                   variant="contained"
                   color="primary"
                   onClick={deleteMember}
               >
                   <div className="good-font">
                       ยืนยัน
                   </div>
               </Button>
       </Dialog>
      
       <Dialog open={openUpdate} onClose={handleClose} >
           <DialogTitle><div className="good-font">ยืนยันการแก้ไขสมาชิกนี้</div></DialogTitle>
           <Button
                   variant="contained"
                   color="primary"

                   component={RouterLink}
                   to="/MemberUpdate"
               >
                   <div className="good-font">
                       ยืนยัน
                   </div>
               </Button>
       </Dialog>
       
     <Container maxWidth="lg">
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
             <div className="title-big">รายชื่อสมาชิกฟาร์มมาร์ท</div>
           </Typography>
         </Box>

         <Box sx={{ paddingX: 1, paddingY: 0 }}>
           <Button
             component={RouterLink}
             to="/MemberCreate"
             variant="contained"
             color="primary"
             startIcon={<GroupAddIcon />}
           >
             <div className="good-font">สมัครสมาชิกฟาร์มมาร์ท</div>
           </Button>
         </Box>

       </Box>
       <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={member}
                        getRowId={(row) => row.ID}
                        columns={columns}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                        onRowClick={handleRowClick}
                    />
                </div>
     </Container>
   </div>
 );
}
export default Member;