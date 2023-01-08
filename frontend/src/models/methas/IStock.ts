import { EmployeeInterface } from "../IEmployee";
import { KindsInterface } from "./IKind";
import { StoragesInterface } from "./IStorage";

export interface StocksInterface{
    ID: number;
    Name: string;
    Quantity: number;
    Price: number;
    Employee_ID: number;
    Employee: EmployeeInterface;
    Kind_ID: number;
    Kind: KindsInterface;
    Storage_ID: number;
    Storage: StoragesInterface;
    DateTime: Date
}