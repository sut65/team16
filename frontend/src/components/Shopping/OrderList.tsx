import React, { useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import EditIcon from '@mui/icons-material/Edit';
import { OrderInterface } from "../../models/Natthapon/IOrder"
import { Dialog, DialogTitle } from "@mui/material";


function Order() {
    const [order, setOrder] = React.useState<OrderInterface[]>([]);
    const [orderID, setOrderID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [price, setPrice] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการแก้ไข/ลบ
    const [num, setNum] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการจ่าย/ลบ
    const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ

    let cartID = localStorage.getItem("cartID"); // เรีกใช้ค่าจากlocal storage 
    let Total = localStorage.getItem("Total"); // เรีกใช้ค่าจากlocal storage 

    // โหลดข้อมูลทั้งหมดใส่ datagrid
    const getOrder = async () => {
        const apiUrl = `http://localhost:8080/ordercart/${cartID}`;
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
                    setOrder(res.data);
                }
                else { console.log("NO DATA") }
            });
    };
    let minus = Number(Total) - Number(price)
    // console.log("total " + Total)
    // console.log("price " + price)
    // console.log("minus " + minus)


    async function sum() {   
        let data = {
            Total: minus,
            Status_ID: 1,
        };
        console.log(minus)

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

    // async function reduce() {
    //     let quantity =  order;
    //     let data = {
    //         Amount: quantity,
    //     };

    //     console.log(quantity)
    //     console.log(data)

    //     const requestOptions = {
    //         method: "PATCH",
    //         headers: {
    //             Authorization: `Bearer ${localStorage.getItem("token")}`,
    //             "Content-Type": "application/json"
    //         },
    //         body: JSON.stringify(data),
    //     };

    //     fetch(`${apiUrl}/UpdateQuantity/${shevID}`, requestOptions)
    //         .then((response) => response.json())
    //         .then((res) => {
    //             if (res.data) {
    //                 setErrorMessage("")
    //             } else {
    //                 setErrorMessage(res.error)
    //             }
    //         });

    // }

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
                    console.log("delete ID: " + orderID)
                }
                else { console.log("NO DATA") }
            });
        handleClose();
        sum();
        getOrder();
    }

    // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือOrderID)เพื่อรอสำหรับการแก้ไข/ลบ
    const handleRowClick: GridEventListener<'rowClick'> = (params) => {
        setOrderID(Number(params.row.ID)); //setเพื่อรอการลบ
        setPrice(Number(params.row.Prices)); //setเพื่อรอการลบ
        setNum(Number(params.row.Quantity)); //setเพื่อรอการลบ
        localStorage.setItem("orderID", params.row.ID); //setเพื่อการแก้ไข
    };
    console.log("Quantity " + num)

     // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การแก้ไข/ลบ
    const handleClose = () => {
        setOpendelete(false)
    };

    useEffect(() => {
        getOrder();
    }, []);

    const columns: GridColDef[] = [
      { field: "ID", headerName: "ID", width: 100,  headerAlign:"center", align:"center" },
      { field: "Shelving", headerName: "สินค้า", width: 100, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.Stock.Name},
      { field: "Quantity", headerName: "จำนวน", width: 150, headerAlign:"center", align:"center" },
      { field: "Prices", headerName: "รวมราคา", width: 150, headerAlign:"center", align:"center" },
      { field: "Shopping_Cart", headerName: "ตะกร้า", width: 150, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.ID},
        //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการแก้ไข/ลบ
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
            <Container maxWidth="md">
                <Box display="flex" sx={{ marginTop: 2,}}>
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