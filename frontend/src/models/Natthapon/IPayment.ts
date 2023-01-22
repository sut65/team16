import { EmployeeInterface } from "../IEmployee";
import { OrderInterface } from "./IOrder";
import { Payment_methodInterface } from "./IPayment_method";

export interface PaymentInterface {
    ID?: number,
    Time?: Date,
    Price?: number,

    Order_ID?: number,
    Order?: OrderInterface
    Payment_method_ID?: number
    Payment_method?: Payment_methodInterface
    Employee_ID?: number;
    Employee?: EmployeeInterface;
}