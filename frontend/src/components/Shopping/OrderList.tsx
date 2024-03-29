import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { OrderInterface } from "../../models/Natthapon/IOrder"
import { Dialog, DialogTitle } from "@mui/material";
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';


function Order() {
    const [order, setOrder] = React.useState<OrderInterface[]>([]);
    const [orderID, setOrderID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [price, setPrice] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false);
    const [num, setNum] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการจ่าย/ลบ
    const [shelvnum, setShelvNum] = React.useState(0);
    const [shevID, setShevID] = React.useState(0);

    let cartID = localStorage.getItem("cartID"); // เรีกใช้ค่าจากlocal storage 
    let Total = localStorage.getItem("Total"); // เรีกใช้ค่าจากlocal storage 


    const requestOptions = {
        method: "GET",
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
        },
    };
    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getOrder = async () => {
        const apiUrl = `http://localhost:8080/ordercart/${cartID}`;
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

    useEffect(() => {
        getOrder();
    }, []);

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือOrderID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setOrderID(Number(params.row.ID)); //setเพื่อรอการลบ
        setPrice(Number(params.row.Prices)); //setเพื่อรอการลบ
        setNum(Number(params.row.Quantity)); //setเพื่อรอการลบ
        localStorage.setItem("orderID", params.row.ID); //setเพื่อการแก้ไข
        localStorage.setItem("shelvID", params.row.Shelving.ID); //setเพื่อการแก้ไข
        localStorage.setItem("Quantity", params.row.Quantity); //setเพื่อการแก้ไข
        localStorage.setItem("Prices", params.row.Prices); //setเพื่อการแก้ไข

        fetch(`http://localhost:8080/order/${orderID}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.Shelving_ID && res.Shelving_Number) {
                    // console.log("shevID " + res.Shelving_ID)
                    console.log("shevNumber " + res.Shelving_Number)
                    setShevID(res.Shelving_ID);
                    setShelvNum(res.Shelving_Number);
                }
                else { console.log("NO DATA") }
            });
        console.log("Quantity " + num)
        console.log("Price " + price)
    };

    // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    // let minus = Number(Total) - Number(price)
    // console.log("total " + Total)
    // console.log("price " + price)
    // console.log("minus " + minus)


    async function sum() {
        let minus = Number(Total) - Number(price);
        let data = {
            Total: minus,
            Status_ID: 1,
        };
        console.log(data)

        // เพิ่มเงื่อนไขอัพเดท Total ในกรณีที่ minus เป็น 0
        if (minus === 0) {
            data.Total = 0;
        }

        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        fetch(`http://localhost:8080/cart/${cartID}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("Total: " + data.Total)
                }
                else { console.log("NO DATA") }

            });

    }


    let Num = Number(shelvnum) + Number(num);
    console.log("Num " + Num)
    async function restore() {
        let Num = Number(shelvnum) + Number(num);
        let data = {
            Number: Num,
        };
        console.log(data)

        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        fetch(`http://localhost:8080/UpdateQuantity/${shevID}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                }
                else { console.log("NO DATA") }
            });

    }

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
                    restore()
                    sum();
                    console.log("delete ID: " + orderID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();

        getOrder();
    }



    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 100, headerAlign: "center", align: "center" },
        { field: "Shelving", headerName: "สินค้า", width: 100, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.Stock.Name },
        { field: "Quantity", headerName: "จำนวน", width: 150, headerAlign: "center", align: "center" },
        { field: "Prices", headerName: "รวมราคา", width: 150, headerAlign: "center", align: "center" },
        { field: "Shopping_Cart", headerName: "ตะกร้า", width: 150, headerAlign: "center", align: "center", valueFormatter: (params) => params.value.ID },
        //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
        {
            field: "edit", headerName: "แก้ไข", width: 100, headerAlign: "center", align: "center", headerClassName: 'green1',
            renderCell: () => {
                return (
                    <Button
                        variant="contained"
                        color="success"
                        onClick={() => setOpenupdate(true)}
                        startIcon={<EditIcon />}
                    >
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
                        color="error"
                        onClick={() => setOpendelete(true)}
                        startIcon={<DeleteIcon />}
                    >
                    </Button>
                );
            },
        },
    ];

    return (
        <div>
            {/* ยืนยันการลบ */}
            <Dialog open={openDelete} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการลบรายการนี้</div></DialogTitle>
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
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขรายการ</div></DialogTitle>
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
                <Box display="flex" sx={{ marginTop: 2, }}>
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            รายการสินค้า
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