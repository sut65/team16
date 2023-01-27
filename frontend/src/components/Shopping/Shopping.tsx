import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";

import { OrderInterface } from "../../models/Natthapon/IOrder"
import { Dialog, DialogTitle } from "@mui/material";


function Order() {
    const [order, setOrder] = React.useState<OrderInterface[]>([]);
    const [orderID, setOrderID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getOrder = async () => {
        const apiUrl = "http://localhost:8080/orders";
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
                    setOrder(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deleteOrder = async () => {
        const apiUrl = `http://localhost:8080/order/${orderID}`;
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
                    console.log("delete ID: " + orderID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getOrder();
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือOrderID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setOrderID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("orderID", params.row.ID); //setเพื่อการแก้ไข
    };

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    useEffect(() => {
        getOrder();
    }, []);

    const columns: GridColDef[] = [
      { field: "ID", headerName: "ID", width: 50,  headerAlign:"center" },
      { field: "Shelving", headerName: "สินค้า", width: 250, headerAlign:"center",valueFormatter: (params)=>params.value.ID},
      { field: "Quantity", headerName: "จำนวน", width: 100, headerAlign:"center" },
        //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
      {
          field: "edit", headerName: "แก้ไข", width: 100,
          renderCell: () => {
              return (
                  <Button
                      variant="contained"
                      color="primary"
                      onClick={() => setOpenupdate(true)}
                  >
                      Edit
                  </Button>
              );
          },
      },
      {
          field: "delete", headerName: "ลบ", width: 100,
          renderCell: () => {
              return (
                  <Button
                      variant="contained"
                      color="primary"
                      onClick={() => setOpendelete(true)}
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
                <DialogTitle><div className="good-font">ยืนยันการลบส่วนลดนี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deleteOrder}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            {/* ยืนยันการแก้ไข */}
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขส่วนลดนี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" ไปที่หน้าแก้ไข
                        component={RouterLink}
                        to="/OrderUpdate"
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
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
                            <div className="good-font">
                                ส่วนลด
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/OrderCreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                เพิ่มส่วนลด
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={order}
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

export default Order;