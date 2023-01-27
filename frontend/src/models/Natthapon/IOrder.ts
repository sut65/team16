import { CartInterface } from "./ICart";
import { IShelving } from "../methas/IShelving";


export interface OrderInterface {
    ID?: number,
    Quantity?: number,
    Price?: number,
    
    Shelving_ID?: number;
    Shelving?: IShelving;
    Shopping_Cart_ID?: number;
    Shopping_Cart?: CartInterface;
}