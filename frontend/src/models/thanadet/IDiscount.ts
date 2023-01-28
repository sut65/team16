import { Discount_Type_Interface } from "./IDiscount_Type";
import { EmployeeInterface } from "../IEmployee";
import { StocksInterface } from "../methas/IStock";

export interface DiscountInterface {
    ID?: number;
    Discount_Price?: number;
    Discount_s: Date | null;
    Discount_e: Date | null;

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Discount_Type_ID?: number;
<<<<<<< HEAD
    Discount_Type?: Discount_Type_Interface;
=======
    Discount_Type?: Discount_TypeInterface;
>>>>>>> issue-126
    Stock_ID?: number;
    Stock?: StocksInterface;
}