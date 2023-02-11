import { Discount_Type_Interface } from "./IDiscount_Type";
import { EmployeeInterface } from "../IEmployee";
import { ShelvingsInterface } from "../methas/IShelving";

export interface DiscountInterface {
    ID?: number;
    Discount_Price?: number;
    Discount_s: Date | null;
    Discount_e: Date | null;

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Discount_Type_ID?: number;
    Discount_Type?: Discount_Type_Interface;
    Shelving_ID?: number;
    Shelving?: ShelvingsInterface;
}