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
                // Authorization: `Bearer ${localStorage.getItem("token")}`,
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
        { field: "id", headerName: "ลำดับ", width: 50 },
        { field: "discount_price", headerName: "ราคาที่ลด", width: 120 },
        {
            field: "discount_s", headerName: "วันที่เริ่มลดราคา", width: 210,
            renderCell: (params) => moment(params.row.date_rec).format('YY-MM-DD')
        },
        {
            field: "discount_e", headerName: "วันที่สิ้นสุด", width: 210,
            renderCell: (params) => moment(params.row.date_rec).format('YY-MM-DD')
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