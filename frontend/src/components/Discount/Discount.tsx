import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import moment from 'moment'

import { DiscountInterface } from "../../models/thanadet/IDiscount"

function Discount() {
    const [discount, setDiscount] = React.useState<DiscountInterface[]>([]);

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


    useEffect(() => {
        getDiscount();
    }, []);
    
    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 50 },
        { field: "Discount_Price", headerName: "ราคาที่ลด", width: 120 },
        {
            field: "Discount_s", headerName: "วันที่เริ่มลดราคา", width: 210,
            renderCell: (params) => moment(params.row.Discount_s).format('YY-MM-DD')
        },
        {
            field: "Discount_e", headerName: "วันที่สิ้นสุด", width: 210,
            renderCell: (params) => moment(params.row.Discount_e).format('YY-MM-DD')
        },
        {
            field: "Discount_Type", headerName: "ประเภท", width: 210,
            valueFormatter : (params) => params.value.Type_Name,
        },
        {
            field: "Stock", headerName: "สินค้า", width: 210,
            valueFormatter : (params) => params.value.Name,
        },
    ];


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
                    />
                </div>
            </Container>
        </div>
    );
}

export default Discount;