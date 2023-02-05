import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import SignIn from "./components/SignIn";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from "@mui/material/AppBar";
import { styled, createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import MuiDrawer from "@mui/material/Drawer";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import IconButton from "@mui/material/IconButton";
import Container from "@mui/material/Container";
import MenuIcon from "@mui/icons-material/Menu";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import Button from "@mui/material/Button";

import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import HomeIcon from "@mui/icons-material/Home";
import PeopleIcon from "@mui/icons-material/People";
import LogoutIcon from '@mui/icons-material/Logout';
import InventoryIcon from '@mui/icons-material/Inventory';
import Homepage from "./components/Homepage";
import MonetizationOnIcon from '@mui/icons-material/MonetizationOn';
import ArticleIcon from '@mui/icons-material/Article';
import LocalShippingIcon from '@mui/icons-material/LocalShipping';
import GroupAddIcon from '@mui/icons-material/GroupAdd';
import GroupRemoveIcon from '@mui/icons-material/GroupRemove';

import "./styles.css"
import Stocks from "./components/Stock/Stock";
import StockCreate from "./components/Stock/StockCreate";
import Member from "./components/Member/Member";
import MemberCreate from "./components/Member/MemberCreate";
import Discount from "./components/Discount/Discount";
import DiscountCreate from "./components/Discount/DiscountCreate";
import DiscountUpdate from "./components/Discount/DiscountUpdate";
import { GetCurrentEmployee } from "./services/HttpClientService";
import Leave from "./components/Leave/Leave";
import LeaveCreate from "./components/Leave/LeaveCreate";
import SeparationShow from "./components/apisit/separationShow";
import SeparationCreate from "./components/apisit/separationCreate";
import OrderList from "./components/Shopping/OrderList";
import OrderCreate from "./components/Shopping/OrderCreate";
import OrderAdd from "./components/Shopping/OrderAdd";
import Cart from "./components/Shopping/Cart";
import Delivery from "./components/Delivery/Delivery";
import DeliveryCreate from "./components/Delivery/DeliveryCreate";
import DeliveryUpdate from "./components/Delivery/DeliveryUpdate";
import SeparationUpdate from "./components/apisit/separationUpdate";
import LeaveUpdate from "./components/Leave/LeaveUpdate";
import MemberUpdate from "./components/Member/MemberUpdate";
import Payment from "./components/Payment/payment";
import PaymentCreate from "./components/Payment/PaymentCreate";
import PaymentUpdate from "./components/Payment/PaymentUpdate";
import Pay from "./components/Payment/pay";
import EmployeeattemdanceIN from "./components/Panupol/Em_IN_list";
import EmployeeattemdanceINcreate from "./components/Panupol/Em_in";
import EmployeeattemdanceOUT from "./components/Panupol/Em_OUT_list";
import EmployeeattemdanceOUTcreate from "./components/Panupol/Em_out";
import EmployeeattemdanceINUpdate from "./components/Panupol/Em_IN_update";
import EmployeeattemdanceOutUpdate from "./components/Panupol/Em_Out_update";
import CommentShow from "./components/comments/commentShow";
import CommentCreate from "./components/comments/commentCreate";
import CommentUpdate from "./components/comments/commentUpdate";
import StockUpdate from "./components/Stock/StockUpdate";

var employeeName = "";

const getEmployee = async () => {
  let res = await GetCurrentEmployee();
  employeeName = res.Name;
  if (res) {
    console.log(res);
    console.log(employeeName);
  }
};

const drawerWidth = 240;

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== "open",
})<AppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(["width", "margin"], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== "open",
})(({ theme, open }) => ({
  "& .MuiDrawer-paper": {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
    boxSizing: "border-box",
    ...(!open && {
      overflowX: "hidden",
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9),
      },
    }),
  },
}));

const menu = [
  { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },
  { name: "ตะกร้าสินค้า", icon: <AddShoppingCartIcon />, path: "/Cart" },
  { name: "สต๊อกสินค้า", icon: <InventoryIcon />, path: "/Stock" },
  { name: "สมาชิก", icon: <PeopleIcon />, path: "/Member" },
  { name: "ส่วนลด", icon: <MonetizationOnIcon />, path: "/Discount" },
  { name: "แจ้งลา", icon: <ArticleIcon />, path: "/Leave" },
  { name: "จำหน่ายสินค้า", icon: <InventoryIcon />, path: "/SeparationShow" },
  { name: "รายการการส่งสินค้า", icon: <LocalShippingIcon />, path: "/Delivery" },
  { name: "รายการเข้างาน", icon: <GroupAddIcon />, path: "/EmployeeattemdanceIN" },
  { name: "รายการออกงาน", icon: <GroupRemoveIcon />, path: "/EmployeeattemdanceOUT" },
  { name: "ความคิดเห็น", icon: <LocalShippingIcon />, path: "/CommentShow" },
];


export default function App() {
  const [token, setToken] = useState<String>("");
  const [open, setOpen] = React.useState(true);
  const toggleDrawer = () => {
    setOpen(!open);
  };

  useEffect(() => {
    
    getEmployee();

    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    };

    setTimeout(() => {
      setOpen(false);
    }, 500);
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  const theme = createTheme({
    typography: {
      fontFamily: "Prompt, sans-serif",
      fontSize: 16,
    },
    palette: {
      primary: {
        main: '#98A8F8',
      },
      secondary: {
        main: '#FAF7F0',
      },
    },
  });

  return (
    <div className="App">
      <Router>
        <ThemeProvider theme={theme}>
          <Box sx={{ display: "flex" }}>
            <CssBaseline />
            <AppBar color="primary" position="absolute" open={open}>
              <Toolbar
                sx={{
                  pr: "24px", // keep right padding when drawer closed
                }}
              >
                <IconButton
                  edge="start"
                  color="secondary"
                  aria-label="open drawer"
                  onClick={toggleDrawer}
                  sx={{
                    marginRight: "36px",
                    ...(open && { display: "none" }),
                  }}
                >
                  <MenuIcon />
                </IconButton>
                <Typography
                  component="h1"
                  variant="h6"
                  color="secondary"
                  noWrap
                  sx={{ flexGrow: 1 }}
                >
                  <div className="good-font-big">
                    ระบบฟาร์มมาร์ท
                  </div>
                </Typography>
                <Typography
                  variant="inherit"
                  sx={{ flexGrow: 0.1 }}
                >
                  <span className="good-font-white">{employeeName}</span><span className="good-font-green"> : กำลังใช้งาน</span>
                </Typography>
                <Button color="secondary" onClick={signout} variant="outlined">
                  <LogoutIcon />
                  <Typography
                    color="#FF6464"
                    variant="button">
                    <div className="good-font">
                      ออกจากระบบ
                    </div>
                  </Typography>
                </Button>
              </Toolbar>
            </AppBar>
            <Drawer variant="permanent" open={open}>
              <Toolbar
                sx={{
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "flex-end",
                  px: [1],
                }}
              >
                <IconButton onClick={toggleDrawer}>
                  <ChevronLeftIcon />
                </IconButton>
              </Toolbar>
              <Divider />
              <List>
                {menu.map((item, index) => (
                  <Link
                    to={item.path}
                    key={item.name}
                    style={{ textDecoration: "none", color: "inherit" }}
                  >
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
            <Box
              component="main"
              sx={{
                backgroundColor: '#FAF7F0',
                flexGrow: 1,
                height: "100vh",
                overflow: "auto",
              }}
            >
              <Toolbar />
              <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
                <Routes>
                  <Route path="/" element={<Homepage />} />

                  <Route path="/Member" element={<Member />} />
                  <Route path="/MemberCreate" element={<MemberCreate />} />
                  <Route path="/MemberUpdate" element={<MemberUpdate />} />
                  <Route path="/Stock" element={<Stocks />} />
                  <Route path="/StockCreate" element={<StockCreate />} />
                  <Route path="/StockUpdate" element={<StockUpdate />} />
                  <Route path="/Discount" element={<Discount />} />
                  <Route path="/DiscountCreate" element={<DiscountCreate />} />
                  <Route path="/DiscountUpdate" element={<DiscountUpdate />} />
                  <Route path="/Delivery" element={<Delivery />} />
                  <Route path="/DeliveryCreate" element={<DeliveryCreate />} />
                  <Route path="/DeliveryUpdate" element={<DeliveryUpdate />} />
                  <Route path="/Leave" element={<Leave />} />
                  <Route path="/LeaveCreate" element={<LeaveCreate />} />
                  <Route path="/LeaveUpdate" element={<LeaveUpdate />} />
                  <Route path="/separationShow" element={<SeparationShow />} />
                  <Route path="/separationCreate" element={<SeparationCreate />} />
                  <Route path="/separationUpdate" element={<SeparationUpdate />} />
                  <Route path="/commentShow" element={<CommentShow />} />
                  <Route path="/commentCreate" element={<CommentCreate />} />
                  <Route path="/commentUpdate" element={<CommentUpdate />} />
                  <Route path="/OrderList" element={<OrderList />} />
                  <Route path="/OrderCreate" element={<OrderCreate />} />
                  <Route path="/OrderAdd" element={<OrderAdd />} />
                  <Route path="/Cart" element={<Cart />} />
                  <Route path="/Payment" element={<Payment />} />
                  <Route path="/PaymentCreate" element={<PaymentCreate />} />
                  <Route path="/EmployeeattemdanceIN" element={<EmployeeattemdanceIN />} />
                  <Route path="/EmployeeattemdanceINcreate" element={<EmployeeattemdanceINcreate />} />
                  <Route path="/EmployeeattemdanceOUT" element={<EmployeeattemdanceOUT />} />
                  <Route path="/EmployeeattemdanceOUTcreate" element={<EmployeeattemdanceOUTcreate />} />
                  <Route path="/PaymentUpdate" element={<PaymentUpdate />} />
                  <Route path="/Pay" element={<Pay />} />
                  <Route path="/EmployeeattemdanceINUpdate" element={<EmployeeattemdanceINUpdate />} />
                  <Route path="/EmployeeattemdanceOutUpdate" element={<EmployeeattemdanceOutUpdate />} />
                </Routes>
              </Container>
            </Box>
          </Box>
        </ThemeProvider>
      </Router>
    </div>
  );

}
