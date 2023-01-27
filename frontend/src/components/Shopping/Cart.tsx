  import React, { useEffect } from "react";
  import { Link as RouterLink } from "react-router-dom";
  import Typography from "@mui/material/Typography";
  import Button from "@mui/material/Button";
  import Container from "@mui/material/Container";
  import Box from "@mui/material/Box";
  import { DataGrid, GridColDef, GridEventListener } from "@mui/x-data-grid";
  import { CartInterface } from "../../models/Natthapon/ICart";
  import EditIcon from '@mui/icons-material/Edit';
  import ShoppingCartCheckoutIcon from '@mui/icons-material/ShoppingCartCheckout';
  import PaymentIcon from '@mui/icons-material/Payment';
  import { Dialog, DialogTitle } from "@mui/material";

  function Cart() {
  const [cart, setCart] = React.useState<CartInterface[]>([]);
  const [cartID, setCartID] = React.useState(0); // เก็บค่าIDของข้อมูลที่ต้องการจ่าย/ลบ
  const [openDelete, setOpendelete] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การลบ
  const [openPament, setOpenPament] = React.useState(false); // มีเพ่ือsetการเปิดปิดหน้าต่าง"ยืนยัน"การจ่าย

  const getCart = async () => {
  const apiUrl = "http://localhost:8080/unpaids";
    const requestOptions = {
      method: "GET",
      headers: {Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",},
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          console.log(res.data)
          setCart (res.data);
        }else { console.log("NO DATA") }
      });
  };

    // function ลบข้อมูล
    const deleteCart = async () => {
      const apiUrl = `http://localhost:8080/cart/${cartID}`;
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
                  console.log("delete ID: " + cartID)
              }
              else { console.log("NO DATA") }
          });
      handleClose();
      getCart();
  }

  // เมื่อมีการคลิ๊กที่แถวใดแถวหนึ่งในDataGrid functionนี้จะsetค่าIDของข้อมูลที่ต้องการ(ในกรณีนี้คือCartID)เพื่อรอสำหรับการจ่าย/ลบ
  const handleRowClick: GridEventListener<'rowClick'> = (params) => {
      setCartID(Number(params.row.ID)); //setเพื่อรอการลบ
      localStorage.setItem("cartID", params.row.ID); //setเพื่อการจ่าย
  };

   // function มีเพื่อปิดหน้าต่าง "ยืนยัน" การจ่าย/ลบ
  const handleClose = () => {
      setOpendelete(false)
      setOpenPament(false)
  };

  useEffect(() => {
      getCart();
  }, []);


  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 30,  headerAlign:"center", align:"center" },
    { field: "Total", headerName: "รวมยอด", width: 80,  headerAlign:"center", align:"center" },
    { field: "Status", headerName: "สถานะการชำระ", width: 120, headerAlign:"center", align:"center",valueFormatter: (params)=>params.value.Status},
    { field: "Member", headerName: "สมากชิก", width: 200, headerAlign:"center" , align:"center",valueFormatter: (params)=>params.value.Mem_Name},
    { field: "Employee", headerName: "พนักงาน", width: 200, headerAlign:"center" , align:"center",valueFormatter: (params)=>params.value.Name},
    //ปุ่ม delete กับ edit เรียกหน้าต่างย่อย(Dialog) เพื่อให้ยืนยันการจ่าย/ลบ
    {
      field: "pay", headerName: "ชำระ", width: 100,
      renderCell: () => {
          return (
              <Button
                  variant="contained"
                  color="primary"
                  onClick={() => setOpenPament(true)}
              >
                  pay
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

  useEffect(() => {
    getCart();
  }, []);

  return (

    <div>
        {/* ยืนยันการลบ */}
        <Dialog open={openDelete} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการลบตะกร้า</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" เพื่อเรียก function ลบข้อมูล
                        onClick={deleteCart}
                    >
                        <div className="good-font">
                            ยืนยัน
                        </div>
                    </Button>
            </Dialog>
            {/* ยืนยันการจ่าย */}
            <Dialog open={openPament} onClose={handleClose} >
                <DialogTitle><div className="good-font">ยืนยันการชำระ</div></DialogTitle>
                <Button
                        variant="contained"
                        color="primary"
                        //กด "ยืนยัน" ไปที่หน้าจ่าย
                        component={RouterLink}
                        to="/Pay"
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
                รายการตะกร้าสินค้า
            </Typography>
          </Box>

        <Box sx={{ paddingX: 1, paddingY: 0 }}> 
          <Button
              component={RouterLink}
              to="/Payment"
              variant="contained"
              color="primary"
              startIcon={<PaymentIcon />}
          >
          ประวัติการชำระ
          </Button>
        </Box>
        
        <Box sx={{ paddingX: 1, paddingY: 0 }}> 
          <Button
              component={RouterLink}
              to="/Order"
              variant="contained"
              color="primary"
              startIcon={<EditIcon />}
          >
          รายการสินค้า
        </Button>
        </Box>


        <Box sx={{ paddingX: 1, paddingY: 0 }}>
        <Button
            component={RouterLink}
            to="/OrderCreate"
            variant="contained"
            color="primary"
            startIcon={<ShoppingCartCheckoutIcon />}
        >
            สร้างรายการสินค้า
        </Button>
        </Box>

        </Box>
        
          <div style={{ height: 550, width: "100%", marginTop: '20px'}}>
              <DataGrid
              rows={cart}
              getRowId={(row) => row.ID}
              columns={columns}
              pageSize={10}
              rowsPerPageOptions={[5]}
              onRowClick={handleRowClick}
              />
          </div>
        </Container>
    </div>
  );
  }
  export default Cart;