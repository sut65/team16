import { EmployeeInterface } from "../IEmployee";
import { KindsInterface } from "./IKind";
import { StoragesInterface } from "./IStorage";

//ระบบสต๊อกสินค้า
export interface StocksInterface{
    ID: number,
    Name?: string,
    Amount: number,
    Price: number,
    Employee_ID: number,
    Employee: EmployeeInterface,
    Kind_ID: number,
    Kind: KindsInterface,
    Storage_ID: number,
    Storage: StoragesInterface,
    DateTime: Date | null
}