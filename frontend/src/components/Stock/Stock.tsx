import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { Dialog, DialogTitle } from "@mui/material";
import { StocksInterface } from "../../models/methas/IStock";



function Stock() {
    const [stock, setStock] = React.useState<StocksInterface[]>([]);
    const [stockID, setStockID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpenDelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenUpdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getStocks = async () => {
        const apiUrl = "http://localhost:8080/stocks";
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
                    setStock(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deleteStock = async () => {
        const apiUrl = `http://localhost:8080/stocks/${stockID}`;
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
                    console.log("delete ID: " + stockID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getStocks();
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือstockID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setStockID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("stockID", params.row.ID); //setเพื่อการแก้ไข
    };

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpenDelete(false)
        setOpenUpdate(false)
    };

    useEffect(() => {
        getStocks();
    }, []);

    const columns: GridColDef[] = [

      { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },
      { field: "Name", headerName: "Name", width: 150, headerAlign:"center" },
      { field: "Amount", headerName: "Amount", width: 150, headerAlign:"center" },
      { field: "Price", headerName: "Price", width: 150, headerAlign:"center" },
      { field: "Kind", headerName: "Kind", valueFormatter:(params) => params.value.Name, width: 150, headerAlign:"center" },
      { field: "Storage", headerName: "Storage", valueFormatter:(params) => params.value.Name, width: 150, headerAlign:"center" },
      { field: "DateTime", headerName: "DateTime", width: 200, headerAlign:"center" },
        //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
        {
            field: "Edit", headerName: "Edit", width: 120,
            renderCell: () => {
                return (
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => setOpenUpdate(true)}
                        startIcon={<EditIcon />}
                    >
                        Edit
                    </Button>
                );
            },
        },
        {
            field: "Delete", headerName: "Delete", width: 120,
            renderCell: () => {
                return (
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => setOpenDelete(true)}
                        startIcon={<DeleteIcon />}
                    >
                        Delete
                    </Button>
                );
            },
        },
    ];

    return (
        <div>
            {/* ยืนยันการลบ */}
            <Dialog open={openDelete} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการลบสต๊อก</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deleteStock}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            {/* ยืนยันการแก้ไข */}
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการอัปเดตสต๊อก</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" ไปที่หน้าแก้ไข
                        component={RouterLink}
                        to="/StockUpdate"
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            <Container maxWidth="lg">
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
                            <div className="good-font">
                                สต๊อกสินค้า
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/StockCreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                เพิ่มสต๊อก
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={stock}
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

export default Stock;
