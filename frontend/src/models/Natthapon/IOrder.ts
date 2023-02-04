import { CartInterface } from "./ICart";
import { ShelvingsInterface } from "../methas/IShelving";


export interface OrderInterface {
    ID?: number,
    Quantity?: number,
    Prices?: number,
    
    Shelving_ID?: number;
    Shelving?: ShelvingsInterface;
    Shopping_Cart_ID?: number;
    Shopping_Cart?: CartInterface;
}