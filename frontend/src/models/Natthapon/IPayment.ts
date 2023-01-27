import { EmployeeInterface } from "../IEmployee";
import { CartInterface } from "./ICart";
import { Payment_methodInterface } from "./IPayment_method";

export interface PaymentInterface {
    ID?: number,
    Time?: Date| null,
    Price?: number,

    Shopping_Cart_ID?: number,
    Shopping_Cart?: CartInterface
    Payment_method_ID?: number
    Payment_method?: Payment_methodInterface
    Employee_ID?: number;
    Employee?: EmployeeInterface;
}