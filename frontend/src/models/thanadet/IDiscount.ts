import { Discount_TypeInterface } from "./IDiscount_Type";
import { EmployeeInterface } from "../IEmployee";
import { StocksInterface } from "../methas/IStock";

export interface DiscountInterface {
    ID?: number,
    Discount_Price?: number
    Discount_s: Date | null,
    Discount_e: Date | null,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Discount_Type_ID?: number;
    Discount_Type?: Discount_TypeInterface;
    Stock_ID?: number;
    Stock?: StocksInterface;
}