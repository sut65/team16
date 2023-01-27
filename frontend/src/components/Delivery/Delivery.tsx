import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'

import { DeliveryInterface } from "../../models/thanadet/IDelivery"
import { Dialog, DialogTitle } from "@mui/material";


function Delivery() {
    const [delivery, setDelivery] = React.useState<DeliveryInterface[]>([]);
    const [deliveryID, setDeliveryID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getDelivery = async () => {
        const apiUrl = "http://localhost:8080/deliveries";
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
                    setDelivery(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deleteDelivery = async () => {
        const apiUrl = `http://localhost:8080/delivery/${deliveryID}`;
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
                    console.log("delete ID: " + deliveryID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getDelivery();
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือdeliveryID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setDeliveryID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("deliveryID", params.row.ID); //setเพื่อการแก้ไข
    };

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    useEffect(() => {
        getDelivery();
    }, []);

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 50 },
        { field: "Location", headerName: "สถานที่", width: 80 },
        { field: "Customer_name", headerName: "ชื่อลูกค้า", width: 80 },
        {
            field: "Delivery_date", headerName: "วันที่สิ้นสุด", width: 150,
            renderCell: (params) => moment(params.row.Delivery_date).format('YY-MM-DD')
        },
        {
            field: "Employee", headerName: "ชื่อพนักงานที่ส่งสินค้า", width: 150,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Car", headerName: "รถยนต์", width: 150,
            valueFormatter: (params) => params.value.Registation_Number,
        },
        {
            field: "Payment", headerName: "รายการสินค้า", width: 150,
            valueFormatter: (params) => params.value.ID,
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
                <DialogTitle><div className="good-font">ยืนยันการลบรายการการส่งสินค้านี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deleteDelivery}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            {/* ยืนยันการแก้ไข */}
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขรายการการส่งสินค้านี้นี้</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" ไปที่หน้าแก้ไข
                        component={RouterLink}
                        to="/DeliveryUpdate"
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
                                ส่วนลด
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/DeliveryCreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                เพิ่มรายการการส่งสินค้า
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={delivery}
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

export default Delivery;