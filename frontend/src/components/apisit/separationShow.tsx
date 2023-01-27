import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { SeparationInterface } from "../../models/apisit/ISeparation";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from "moment";
import { Dialog, DialogTitle } from "@mui/material";


function SeparationShow() {
    const [separation, setSeparation] = React.useState<SeparationInterface[]>([]);
    const [separationID, setSeparationID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข

    const getSeparation = async () => {
        const apiUrl = "http://localhost:8080/separations";
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json"
            },
        };

        await fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setSeparation(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // function ลบข้อมูล
    const deleteSeparation = async () => {
        const apiUrl = `http://localhost:8080/separation/${separationID}`;
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
                    console.log("delete ID: " + separationID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getSeparation();
    }

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลับดับ", width: 100 },
        // { field: "Cuase", headerName: "เหตุผล", width: 120 },
        {
            field: "Reason", headerName: "เหตุผล", width: 150,
            valueFormatter: (params) => params.value.Cuase,
        },
        // { field: "name", headerName: "พนักงาน", width: 120 },
        {
            field: "Employee", headerName: "พนักงาน", width: 180,
            valueFormatter: (params) => params.value.Name,
        },
        // { field: "shelving_id", headerName: "ชั้นวาง", width: 120 },
        {
            field: "Shelving", headerName: "สินค้า", width: 100,
            valueFormatter: (params) => params.value.ID,
        },
        // { field: "date_out", headerName: "เวลาจำหน่าย", width: 150 },
        {
            field: "Date_Out", headerName: "เวลาจำหน่าย", width: 150,
            renderCell: (params) => moment(params.row.Date_Out).format('YY-MM-DD')
        },
        { field: "Amount", headerName: "จำนวน", width: 100 },
        { field: "Status", headerName: "สถานะ", width: 100 },

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


    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setSeparationID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("separationID", params.row.ID); //setเพื่อการแก้ไข
    };

    // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };


    useEffect(() => {
        getSeparation();
    }, []);


    return (
        <div>
            {/* ยืนยันการลบ */}
            <Dialog open={openDelete} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการลบรายการจำหน่าย</div></DialogTitle>
                <Button
                    variant="contained"
                    color="primary"
                    //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                    onClick={deleteSeparation}
                >
                    <div className="good-font">
                        ยืนยัน
                    </div>
                </Button>
            </Dialog>
            {/* ยืนยันการแก้ไข */}
            <Dialog open={openUpdate} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการแก้ไขรายการจำหน่าย</div></DialogTitle>
                <Button
                    variant="contained"
                    color="primary"
                    //กด "ยืนยัน" ไปที่หน้าแก้ไข
                    component={RouterLink}
                    to="/SeparationUpdate"
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
                                การจำหน่ายสินค้า
                            </div>
                        </Typography>
                    </Box>
                    <Box>
                        <Button
                            component={RouterLink}
                            to="/SeparationCreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                บันทึกรายการจำหน่ายสินค้า
                            </div>
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={separation}
                        // getRowId={(row: any) =>  row.id}
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

export default SeparationShow;