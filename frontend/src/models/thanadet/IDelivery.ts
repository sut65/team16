import { EmployeeInterface } from "../IEmployee";
import { CarInterface } from "./ICar";
import { PaymentInterface } from "../Natthapon/IPayment";

export interface DeliveryInterface {
    ID?: number,
    Location?: string
    Customer_name?: string,
    Delivery_date?: Date,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Car_ID?: number;
    Car?: CarInterface;
    Payment_ID?: number;
    Payment?: PaymentInterface;
}