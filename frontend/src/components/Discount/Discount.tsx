import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'

import { DiscountInterface } from "../../models/thanadet/IDiscount"
import { Dialog, DialogTitle } from "@mui/material";


function Discount() {
    const [discount, setDiscount] = React.useState<DiscountInterface[]>([]);
    const [discountID, setDiscountID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [cost, setCost] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [shelvingID, setShelvingID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getDiscount = async () => {
        const apiUrl = "http://localhost:8080/discounts";
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
                    setDiscount(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deleteDiscount = async () => {
        const apiUrl = `http://localhost:8080/discount/${discountID}`;
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
                    console.log("delete ID: " + discountID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        resetCost();
        setTimeout(() => {
            getDiscount();
        }, 0);
    }

    async function resetCost() {
        let data = {
            Cost: cost,
        };
        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        };

        fetch(`http://localhost:8080/reset_cost/${shelvingID}/${cost}`, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    if (res.data) {
                        console.log("reset cost ID: " + discountID)
                    }
                    else { console.log("NO DATA") }
                }
            });
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือdiscountID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setDiscountID(Number(params.row.ID)); //setเพื่อรอการลบ
        setCost(Number(params.row.Discount_Price)); //setเพื่อรอการลบ
        setShelvingID(Number(params.row.Shelving_ID)); //setเพื่อรอการลบ
        localStorage.setItem("discountID", params.row.ID); //setเพื่อการแก้ไข
        console.log(cost);
        console.log(shelvingID);
    };

    // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    useEffect(() => {
        getDiscount();
    }, []);

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 50 },
        { field: "Discount_Price", headerName: "ราคาที่ลด", width: 80 },
        {
            field: "Discount_s", headerName: "วันที่เริ่มลดราคา", width: 150,
            renderCell: (params) => moment(params.row.Discount_s).format('YY-MM-DD')
        },
        {
            field: "Discount_e", headerName: "วันที่สิ้นสุด", width: 150,
            renderCell: (params) => moment(params.row.Discount_e).format('YY-MM-DD')
        },
        {
            field: "Discount_Type", headerName: "ประเภท", width: 180,
            valueFormatter: (params) => params.value.Type_Name,
        },
        {
            field: "Shelving", headerName: "สินค้า(ราคาปัจจุบัน)", width: 150,
            valueFormatter: (params) => params.value.Stock.Name + " ราคา " +params.value.Cost,
        },
        // ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
        {
            field: "edit", headerName: "แก้ไข", width: 100,
            renderCell: () => {
                return (
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={() => setOpenupdate(true)}
                    >
                        แก้ไข
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
                        ลบ
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
                    onClick={deleteDiscount}
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
                    to="/DiscountUpdate"
                >
                    <div className="good-font">
                        ยืนยัน
                    </div>
                </Button>
            </Dialog>
            <Container maxWidth='xl'>
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
                            <div className="title-big">
                                ส่วนลด
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/DiscountCreate"
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
                        rows={discount}
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

export default Discount;