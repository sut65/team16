import { EmployeeInterface } from "../IEmployee";
import { Shopping_CartInterface } from "./IShopping_Cart";
import { Payment_methodInterface } from "./IPayment_method";

export interface PaymentInterface {
    ID?: number,
    Time?: Date,
    Price?: number,

    Shopping_Cart_ID?: number,
    Shopping_Cart?: Shopping_CartInterface
    Payment_method_ID?: number
    Payment_method?: Payment_methodInterface
    Employee_ID?: number;
    Employee?: EmployeeInterface;
}