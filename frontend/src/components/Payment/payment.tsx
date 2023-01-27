import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'

import { PaymentInterface } from "../../models/Natthapon/IPayment"
import { Dialog, DialogTitle } from "@mui/material";


function Payment() {
    const [payment, setPayment] = React.useState<PaymentInterface[]>([]);
    const [paymentID, setPaymentID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getPayment = async () => {
        const apiUrl = "http://localhost:8080/payments";
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
                    setPayment(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deletePayment = async () => {
        const apiUrl = `http://localhost:8080/payment/${paymentID}`;
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
                    console.log("delete ID: " + paymentID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getPayment();
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือPaymentID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setPaymentID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("paymentID", params.row.ID); //setเพื่อการแก้ไข
    };

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    useEffect(() => {
        getPayment();
    }, []);

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 50 },
        { field: "Price", headerName: "ยอดรวม", width: 80 },
        {
            field: "Time", headerName: "วันที่ชำระ", width: 200,
            renderCell: (params) => moment(params.row.Time).format('YY-MM-DD HH:mm:ss')
        },
        {
            field: "Shopping_Cart", headerName: "ตะกร้า", width: 100,
            valueFormatter: (params) => params.value.ID,
        },
        {
            field: "Payment_method", headerName: "ช่องทางการขำระ", width: 180,
            valueFormatter: (params) => params.value.Method,
        },
        {
            field: "Employee", headerName: "พนักงาน", width: 250,
            valueFormatter: (params) => params.value.Name,
        },
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
                <DialogTitle><div className="good-font">ยืนยันการลบการชำระ</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deletePayment}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            {/* ยืนยันการแก้ไข */}
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขการชำระ</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" ไปที่หน้าแก้ไข
                        component={RouterLink}
                        to="/PaymentUpdate"
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
                                รายการการชำระสินค้า
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/PaymentCreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ชำระสินค้า
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={payment}
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

export default Payment;