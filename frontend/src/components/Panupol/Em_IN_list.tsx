import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import moment from 'moment'
import { Dialog, DialogTitle } from "@mui/material";
import Typography from "@mui/material/Typography";

import { Employee_attendanceInterface } from "../../models/panupol/IEm_in"


function Employeeattemdance() {
    const [Employeeattemdance, setEmployeeattemdance] = React.useState<Employee_attendanceInterface[]>([]);
    const [EmployeeattemdanceID, setEmployeeattemdanceID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
    const [openUpdate, setOpenupdate] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การแก้ไข


    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getEmployeeattemdance = async () => {
        const apiUrl = "http://localhost:8080/employee_attendances";
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
                    setEmployeeattemdance(res.data);
                }
                else { console.log("NO DATA") }
            });
    };

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือdiscountID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setEmployeeattemdanceID(Number(params.row.ID)); //setเพื่อรอการลบ
        localStorage.setItem("Employee_attendanceID", params.row.ID); //setเพื่อการแก้ไข
    };

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
     const handleClose = () => {
        setOpendelete(false)
        setOpenupdate(false)
    };

    // function ลบข้อมูล
    const deleteEmployeeattemdance = async () => {
        const apiUrl = `http://localhost:8080/employee_attendances/${EmployeeattemdanceID}`;
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
                    console.log("delete ID: " + EmployeeattemdanceID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        getEmployeeattemdance();
    }

    useEffect(() => {
        getEmployeeattemdance();
    }, []);

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 50 },
        { field: "Employee", headerName: "พนักงาน", width: 220 ,headerAlign:"center" ,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Duty", headerName: "หน้าที่", width: 130,headerAlign:"center" ,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Working_time", headerName: "ระยะเวลาในการทำงาน", width: 200,headerAlign:"center" ,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Overtime", headerName: "ทำงานนอกเวลา", width: 180,headerAlign:"center" ,
            valueFormatter: (params) => params.value.Name,
        },
        {
            field: "Time_IN", headerName: "เวลาเข้างาน", width: 170,headerAlign:"center" ,
            renderCell: (params) => moment(params.row.Time_IN).format('YYYY-MM-DD HH:mm:ss')
            
        },
        { field: "Number_Em", headerName: "เบอร์โทร", width: 150 ,headerAlign:"center" ,},
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
                <DialogTitle><div className="good-font">ยืนยันการลบรายการ</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deleteEmployeeattemdance}
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
                        to="/EmployeeattemdanceINUpdate"
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
                                การเข้างาน
                            </div>
                        </Typography>
                    </Box>
                    
                    <Box sx={{ paddingX: 1, paddingY: 0 }}>
                        <Button
                            component={RouterLink}
                            to="/EmployeeattemdanceINcreate"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ลงชื่อเข้างาน
                            </div>
                        </Button>
                    </Box>
                
                </Box>
                <div style={{ height: 400, width: '100%', marginTop: '20px' }}>
                    <DataGrid
                        rows={Employeeattemdance}
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
export default Employeeattemdance;


