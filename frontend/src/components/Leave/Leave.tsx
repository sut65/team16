import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import { LeaveInterface } from "../../models/theerawat/ILeave";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import ArticleIcon from '@mui/icons-material/Article';
import moment from "moment";
import { Dialog, DialogTitle } from "@mui/material";

function Leave() {
  const [leave, setLeave] = React.useState<LeaveInterface[]>([]);
  const [Leave_ID, setLeaveID] = React.useState(0); 
  const [openDelete, setOpendelete] = React.useState(false); 
  const [openUpdate, setOpenupdate] = React.useState(false);

  const getLeave = async () => {
    const apiUrl = "http://localhost:8080/leaves";
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
          setLeave(res.data);
        }
        else { console.log("NO DATA") }
      });
  };

  const deleteLeave = async () => {
    const apiUrl = `http://localhost:8080/leaves/${Leave_ID}`;
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
                console.log("delete ID: " + Leave_ID)
            }
            else { console.log("NO DATA") }
        });
    handleClose();
    getLeave();
  }

  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
    setLeaveID(Number(params.row.ID)); 
    localStorage.setItem("Leave_ID", params.row.ID); 
  };

  const handleClose = () => {
    setOpendelete(false)
    setOpenupdate(false)
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 50 },
    { field: "L_Type", headerName: "ประเภทการลา", width: 150, valueFormatter: (params) => params.value.Type_Name, },
    { field: "Section", headerName: "แผนก", width: 160, valueFormatter: (params) => params.value.Sec_Name, },
    { field: "Doc_Reason", headerName: "รายละเอียด / เหตผลการลา", width: 220 },
    {
      field: "Doc_DateS", headerName: "วันเริ่มลา", width: 100,
      renderCell: (params) => moment(params.row.Doc_DateS).format('YY-MM-DD')
    },
    {
      field: "Doc_DateE", headerName: "วันสิ้นสุดลา", width: 100,
      renderCell: (params) => moment(params.row.Doc_DateE).format('YY-MM-DD')
    },
    { field: "Doc_Cont", headerName: "เบอร์ติดต่อ", width: 140 },
    { field: "edit", headerName: "แก้ไข", width: 90,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpenupdate(true)}
                  startIcon={<EditIcon />}
              > </Button>
              );},},
    { field: "delete", headerName: "ลบ", width: 90,
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
    getLeave();
  }, []);

  return (

    <div>
       <Dialog open={openDelete} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการลบการแจ้งลานี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        onClick={deleteLeave}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
           
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขการแจ้งลานี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"

                        component={RouterLink}
                        to="/LeaveUpdate"
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>

      <Container maxWidth="lg">
        <Box
          display="flex"
          sx={{ marginTop: 2, }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              <div className="good-font">การแจ้งลาของพนักงานฟาร์มมาร์ท</div>
            </Typography>
          </Box>

          <Box sx={{ paddingX: 1, paddingY: 0 }}>
            <Button
              component={RouterLink}
              to="/LeaveCreate"
              variant="contained"
              color="primary"
              startIcon={<ArticleIcon />}
            >
              <div className="good-font">เขียนเอกสารแจ้งลา</div>
            </Button>
          </Box>

        </Box>
        <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={leave}
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
export default Leave;