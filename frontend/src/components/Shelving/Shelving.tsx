import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import { ShelvingsInterface } from "../../models/methas/IShelving";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from '@mui/icons-material/Edit';
import AddIcon from '@mui/icons-material/Add';
import { Dialog, DialogTitle } from "@mui/material";
import moment from "moment";

function Shelving() {
  const [meat, setMeat] = React.useState<ShelvingsInterface[]>([]);
  const [fresh, setFresh] = React.useState<ShelvingsInterface[]>([]);
  const [dairy, setDairy] = React.useState<ShelvingsInterface[]>([]);
  const [shelvingID, setShelvingID] = React.useState(0);
  const [openDelete, setOpenDelete] = React.useState(false);
  const [openUpdate, setOpenUpdate] = React.useState(false);
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };
  const getMeat = async () => {
    const apiUrl = "http://localhost:8080/shelvings/1";
    await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data)
          setMeat(res.data);
        }
        else { console.log("NO DATA") }
      });
  };

  const getFresh = async () => {
    const apiUrl = "http://localhost:8080/shelvings/2";
    await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data)
          setFresh(res.data);
        }
        else { console.log("NO DATA") }
      });
  };

  const getDairy = async () => {
    const apiUrl = "http://localhost:8080/shelvings/3";
    await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data)
          setDairy(res.data);
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
    getMeat();
    getFresh();
    getDairy();
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
    { field: "ID", headerName: "ID", width: 50, headerAlign: "center", align: "center" },
    { field: "Stock", headerName: "ชื่อ", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "Number", headerName: "จำนวน", width: 100, headerAlign: "center", align: "center" },
    { field: "Cost", headerName: "ราคา", width: 150, headerAlign: "center", align: "center" },
    { field: "Label", headerName: "ชั้นวาง", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "DateTime", headerName: "วันที่-เวลา", width: 200, headerAlign: "center",
        renderCell: (params) => moment(params.row.DateTime).format('DD-MM-YYYY HH:mm:ss')
    },
    { field: "Edit", headerName: "แก้ไข", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenUpdate(true)}
            startIcon={<EditIcon />}
          >
            แก้ไข
          </Button>
        );
      },
    },
    {
      field: "Delete", headerName: "ลบ", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenDelete(true)}
            startIcon={<DeleteIcon />}
            >
            ลบ
          </Button>
        );
      },
    },
  ];

  //2
  const columns2: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 50, headerAlign: "center", align: "center" },
    { field: "Stock", headerName: "ชื่อ", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "Number", headerName: "จำนวน", width: 100, headerAlign: "center", align: "center" },
    { field: "Cost", headerName: "ราคา", width: 150, headerAlign: "center", align: "center" },
    { field: "Label", headerName: "ชั้นวาง", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "DateTime", headerName: "วันที่-เวลา", width: 200, headerAlign: "center",
        renderCell: (params) => moment(params.row.DateTime).format('DD-MM-YYYY HH:mm:ss')
    },
    { field: "Edit", headerName: "แก้ไข", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenUpdate(true)}
            startIcon={<EditIcon />}
          >
            แก้ไข
          </Button>
        );
      },
    },
    {
      field: "Delete", headerName: "ลบ", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenDelete(true)}
            startIcon={<DeleteIcon />}
            >
            ลบ
          </Button>
        );
      },
    },
  ];

  //3
  const columns3: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 50, headerAlign: "center", align: "center" },
    { field: "Stock", headerName: "ชื่อ", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "Number", headerName: "จำนวน", width: 100, headerAlign: "center", align: "center" },
    { field: "Cost", headerName: "ราคา", width: 150, headerAlign: "center", align: "center" },
    { field: "Label", headerName: "ชั้นวาง", width: 120, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Name },
    { field: "DateTime", headerName: "วันที่-เวลา", width: 200, headerAlign: "center",
        renderCell: (params) => moment(params.row.DateTime).format('DD-MM-YYYY HH:mm:ss')
    },
    { field: "Edit", headerName: "แก้ไข", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenUpdate(true)}
            startIcon={<EditIcon />}
          >
            แก้ไข
          </Button>
        );
      },
    },
    {
      field: "Delete", headerName: "ลบ", width: 120, headerAlign: "center", align: "center",
      renderCell: () => {
        return (
          <Button
            variant="contained"
            color="primary"
            onClick={() => setOpenDelete(true)}
            startIcon={<DeleteIcon />}
            >
            ลบ
          </Button>
        );
      },
    },
  ];

  useEffect(() => {
    getMeat();
    getFresh();
    getDairy();
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
          sx={{ marginTop: 2, }}
        >

          <Box flexGrow={1}>
            <Button
              variant="contained"
              color="primary"
            >
              <div className="good-font">เนื้อ</div>
            </Button>
          </Box>

          <Box sx={{ paddingX: 1, paddingY: 0 }}>
            <Button
              component={RouterLink}
              to="/ShelvingCreate"
              variant="contained"
              color="primary"
              startIcon={<AddIcon />}
            >
              <div className="good-font">เพิ่มเนื้อบนชั้นวาง</div>
            </Button>
          </Box>

        </Box>
        <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
          <DataGrid
            rows={meat}
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
          sx={{ marginTop: 2, }}
        >
          <Box flexGrow={1}>
            <Button
              variant="contained"
              color="primary"
            >
              <div className="good-font">ผักและผลไม้</div>
            </Button>
          </Box>
          <Box sx={{ paddingX: 1, paddingY: 0 }}>
            <Button
              component={RouterLink}
              to="/ShelvingCreate"
              variant="contained"
              color="primary"
              startIcon={<AddIcon />}
            >
              <div className="good-font">เพิ่มผักและผลไม้บนชั้นวาง</div>
            </Button>
          </Box>

        </Box>
        <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
          <DataGrid
            rows={fresh}
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
          sx={{ marginTop: 2, }}
        >
          <Box flexGrow={1}>
            <Button
              variant="contained"
              color="primary"
            >
              <div className="good-font">ผลิตภัณฑ์ประเภทนม</div>
            </Button>
          </Box>

          <Box sx={{ paddingX: 1, paddingY: 0 }}>
            <Button
              component={RouterLink}
              to="/ShelvingCreate"
              variant="contained"
              color="primary"
              startIcon={<AddIcon />}
            >
              <div className="good-font">เพิ่มผลิตภัณฑ์นมบนชั้นวาง</div>
            </Button>
          </Box>

        </Box>
        <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
          <DataGrid
            rows={dairy}
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
