import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { LeaveInterface } from "../../models/theerawat/ILeave";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import ArticleIcon from '@mui/icons-material/Article';
import moment from "moment";

function Leave() {
  const [leave, setLeave] = React.useState<LeaveInterface[]>([]);
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

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 50, headerAlign: "center" },
    { field: "L_Type", headerName: "Type", width: 150, headerAlign: "center", valueFormatter: (params) => params.value.Type_Name, },
    { field: "Section", headerName: "Section", width: 120, headerAlign: "center", valueFormatter: (params) => params.value.Sec_Name, },
    { field: "Doc_Reason", headerName: "Reason", width: 200, headerAlign: "center" },
    {
      field: "Doc_DateS", headerName: "Start Date", width: 115,
      renderCell: (params) => moment(params.row.date_rec).format('YY-MM-DD')
    },
    {
      field: "Doc_DateE", headerName: "End Date", width: 115,
      renderCell: (params) => moment(params.row.date_rec).format('YY-MM-DD')
    },
    { field: "Doc_Cont", headerName: "Contact", width: 100, headerAlign: "center" },
  ];

  useEffect(() => {
    getLeave();
  }, []);

  return (

    <div>
      <Container maxWidth="md">
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
              Leave
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
              to="/LeaveCreate"
              variant="contained"
              color="primary"
              startIcon={<ArticleIcon />}
            >
              Create
            </Button>
          </Box>

        </Box>
        <div style={{ height: 400, width: "100%", marginTop: '20px' }}>
          <DataGrid
            rows={leave}
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
export default Leave;