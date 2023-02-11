import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'
import PaymentIcon from '@mui/icons-material/Payment';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

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
        localStorage.setItem("cartID", params.row.Shopping_Cart_ID); //setเพื่อการแก้ไข
        localStorage.setItem("total", params.row.Paytotal); //setเพื่อการแก้ไข
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
        { field: "ID", headerName: "ID", width: 50 ,headerAlign:"center", align:"center", headerClassName: 'green1'},
        {
            field: "Shopping_Cart", headerName: "ตะกร้า", width: 60,headerAlign:"center", align:"center", headerClassName: 'green1',
            valueFormatter: (params) => params.value.ID,
        },
        { field: "Paytotal", headerName: "ยอดรวม", width: 80 ,headerAlign:"center", align:"center", headerClassName: 'green1'},
        {
            field: "Payment_method", headerName: "ช่องทางการขำระ", width: 150,headerAlign:"center", align:"center", headerClassName: 'green1',
            valueFormatter: (params) => params.value.Method,
        },
        {
            field: "Time", headerName: "วันที่ชำระสินค้า", width: 180, headerAlign:"center", align:"center", headerClassName: 'green1',
            renderCell: (params) => moment(params.row.Time).format('YY-MM-DD HH:mm:ss')
        },
        {
            field: "Note", headerName: "หมายเหตุ", width: 180,headerAlign:"center", align:"center", headerClassName: 'green1',
            valueFormatter: (params) => params.value.Method,
        },
        {
            field: "Employee", headerName: "พนักงาน", width: 200,headerAlign:"center", align:"center", headerClassName: 'green1',
            valueFormatter: (params) => params.value.Name,
        },
        //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
        {
            field: "edit", headerName: "แก้ไข", width: 100,headerAlign:"center", align:"center", headerClassName: 'green1',
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
            field: "delete", headerName: "ลบ", width: 100,headerAlign:"center", align:"center", headerClassName: 'green1',
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
                <DialogTitle><div className="good-font">ยืนยันการลบรายการ</div></DialogTitle>
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
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขรายการ</div></DialogTitle>
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
            <Container maxWidth="lg">
                <Box display="flex" sx={{ marginTop: 2,}}>
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            ประวัติการชำระสินค้า
                        </Typography>
                    </Box>

                    <Box sx={{ paddingX: 1, paddingY: 0 }}> 
                        <Button
                            component={RouterLink}
                            to="/Cart"
                            variant="contained"
                            color="primary"
                            startIcon={<ArrowBackIcon />}
                        >
                        กลับ
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