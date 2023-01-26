import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { SeparationInterface } from "../../models/apisit/ISeparation";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import moment from "moment";


function SeparationShow() {
    const [separation, setSeparation] = React.useState<SeparationInterface[]>([]);

    const getSeparation = async () => {
        const apiUrl = "http://localhost:8080/separations";
        const requestOptions = {
            method: "GET",
            headers: { 
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json" },
        };

        await fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log(res.data)
                    setSeparation(res.data);
                }
                else {console.log("NO DATA")}
            });
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลับดับ", width: 100 },
        // { field: "Cuase", headerName: "เหตุผล", width: 120 },
        {
            field: "Reason", headerName: "เหตุผล", width: 180,
            valueFormatter: (params) => params.value.Cuase,
        },
        // { field: "name", headerName: "พนักงาน", width: 120 },
        {
            field: "Employee", headerName: "เหตุผล", width: 180,
            valueFormatter: (params) => params.value.Name,
        },
        // { field: "shelving_id", headerName: "ชั้นวาง", width: 120 },
        {
            field: "Shelving", headerName: "เหตุผล", width: 180,
            valueFormatter: (params) => params.value.ID,
        },
        // { field: "date_out", headerName: "เวลาจำหน่าย", width: 150 },
        {
            field: "Date_Out", headerName: "เวลาจำหน่าย", width: 150,
            renderCell: (params) => moment(params.row.Date_Out).format('YY-MM-DD')
        },
        { field: "amount", headerName: "จำนวน", width: 100 },
        { field: "status", headerName: "สถานะ", width: 200 },
        
    ];

    useEffect(() => {
        getSeparation();
    }, []);

    return (
        <div>
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
                    />
                </div>
                {/* <Box>
                        <Button
                            component={RouterLink}
                            to="/Ac_his_sum"
                            variant="contained"
                            color="primary"
                        >
                            <div className="good-font-white">
                                ดูชั่วโมงรวม
                            </div>
                        </Button>
                    </Box> */}
            </Container>
        </div>
    );
}

export default SeparationShow;