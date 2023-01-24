import React, {useEffect, useRef, useState} from 'react'
import MainLayout from './Layout'
import { toast } from 'react-toastify';
// import { ComponentToPrint } from './payment';
import { useReactToPrint } from 'react-to-print';
import { EmployeeInterface } from "../../models/IEmployee";
import { MemberInterface } from "../../models/theerawat/IMember";
import { OrderInterface } from "../../models/Natthapon/IOrder";
import { StocksInterface } from "../../models/methas/IStock";
import { GetCurrentEmployee } from "../../services/HttpClientService";


function POSPage() {
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);
  const [errorMessage, setErrorMessage] = React.useState("");

  const [employee, setEmployee] = React.useState<EmployeeInterface>();
  const [member, setMember] = React.useState<MemberInterface[]>([]);
  const [shelving, setShelving] = React.useState<StocksInterface[]>([]);
  const [order, setOder] = React.useState<OrderInterface>({});

  const [cart, setCart] = useState([]);
  const [totalAmount, setTotalAmount] = useState(0);

  const toastOptions = {
    autoClose: 400,
    pauseOnHover: true,
  }

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json"
    },
  };

  const getShelving = async () => {
    fetch(`${apiUrl}/stock`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data)
          setShelving(res.data);
        }
        else { console.log("NO DATA") }
      });
  };

  const getMember = async () => {
    fetch(`${apiUrl}/member`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data)
          setMember(res.data);
        }
        else { console.log("NO DATA") }
      });
  };

  // const getEmployee = async () => {
  //   let res = await GetCurrentEmployee();
  //   order.Employee_ID = res.ID;
  //   if (res) {
  //       setEmployee(res);
  //       console.log(res)
  //   }
  // };

  // const addProductToCart = async(shelving: StocksInterface) =>{
  //   // check if the adding product exist
  //   let findProductInCart = await order.find(i=>{
  //     return i.ID === shelving.ID
  //   });
  // }
  //   if(findProductInCart){
  //     let newCart: OrderInterface[] = [];
  //     let newItem;

  //     cart.forEach(cartItem => {
  //       if(cartItem.id === shelving.ID){
  //         newItem = {
  //           ...cartItem,
  //           quantity: cartItem.quantity + 1,
  //           totalAmount: cartItem.price * (cartItem.quantity + 1)
  //         }
  //         newCart.push(newItem);
  //       }else{
  //         newCart.push(cartItem);
  //       }
  //     });

  //     setCart(newCart);
  //     toast(`Added ${newItem.name} to cart`,toastOptions)

  //   }else{
  //     let addingProduct = {
  //       ...product,
  //       'quantity': 1,
  //       'totalAmount': product.price,
  //     }
  //     setCart([...cart, addingProduct]);
  //     toast(`Added ${product.name} to cart`, toastOptions)
  //   }

  // }

  // const removeProduct = async(product) =>{
  //   const newCart =cart.filter(cartItem => cartItem.id !== product.id);
  //   setCart(newCart);
  // }

  // const componentRef = useRef();

  // const handleReactToPrint = useReactToPrint({
  //   content: () => componentRef.current,
  // });

  // const handlePrint = () => {
  //   handleReactToPrint();
  // }

  // useEffect(() => {
  //   getShelving();
  //   getMember();
  //   getEmployee();
  // },[]);

  // useEffect(() => {
  //   let newTotalAmount = 0;
  //   cart.forEach(icart => {
  //     newTotalAmount = newTotalAmount + parseInt(icart.totalAmount);
  //   })
  //   setTotalAmount(newTotalAmount);
  // },[cart])

  // onClick={() => addProductToCart(shelving)}

  return (
    <MainLayout>
      <div className='row'>
        <div className='col-lg-8'>
          {<div className='row'>
              {shelving.map((shelving, key) =>
                <div key={key} className='col-lg-4 mb-4'>
                  <div className='pos-item px-3 text-center border' >
                      <p>{shelving.Name}</p>
                      <img src={shelving.Name} className="img-fluid" alt={shelving.Name} />
                      <p>${shelving.Price}</p>
                  </div>

                </div>
              )}
            </div>}
        
        </div>
        {/* <div className='col-lg-4'>
              <div style={{display: "none"}}>
                <ComponentToPrint cart={cart} totalAmount={totalAmount} ref={componentRef}/>
              </div>
              <div className='table-responsive bg-dark'>
                <table className='table table-responsive table-dark table-hover'>
                  <thead>
                    <tr>
                      <td>#</td>
                      <td>Name</td>
                      <td>Price</td>
                      <td>Qty</td>
                      <td>Total</td>
                      <td>Action</td>
                    </tr>
                  </thead>
                  <tbody>
                    { cart ? cart.map((cartProduct, key) => <tr key={key}>
                      <td>{cartProduct.id}</td>
                      <td>{cartProduct.name}</td>
                      <td>{cartProduct.price}</td>
                      <td>{cartProduct.quantity}</td>
                      <td>{cartProduct.totalAmount}</td>
                      <td>
                        <button className='btn btn-danger btn-sm' onClick={() => removeProduct(cartProduct)}>Remove</button>
                      </td>

                    </tr>)

                    : 'No Item in Cart'}
                  </tbody>
                </table>
                <h2 className='px-2 text-white'>Total Amount: ${totalAmount}</h2>
              </div>

              <div className='mt-3'>
                { totalAmount !== 0 ? <div>
                  <button className='btn btn-primary' onClick={handlePrint}>
                    Pay Now
                  </button>

                </div> : 'Please add a product to the cart'

                }
              </div>


        </div> */}
      </div>
    </MainLayout>
  )

}

export default POSPage