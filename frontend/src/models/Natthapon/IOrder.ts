import { EmployeeInterface } from "../IEmployee";
import { Shopping_CartInterface } from "./IShopping_Cart";
import { IShelving } from "../methas/IShelving";


export interface OrderInterface {
    ID?: number,
    Quantity?: number,

    Shopping_Cart_ID?: number;
    Shopping_Cart?: Shopping_CartInterface;
    Shelving_ID?: number;
    Shelving?: IShelving;
}