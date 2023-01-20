import { Discount_TypeInterface } from "./IDiscount_Type";
import { EmployeeInterface } from "../IEmployee";
import { StocksInterface } from "../methas/IStock";

export interface Shopping_CartInterface {
    ID?: number,
    Discount_Price?: number
    Discount_s?: number,
    Discount_e?: number,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Discount_Type_ID?: Discount_TypeInterface;
    Discount_Type?: number;
    Stock_ID?: number;
    Stock?: StocksInterface;
}