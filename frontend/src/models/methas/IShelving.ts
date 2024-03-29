import { EmployeeInterface } from "../IEmployee";
import { LabelsInterface } from "./ILabel";
import { StocksInterface } from "./IStock";
//ระบบชั้นวางสินค้า
export interface ShelvingsInterface{
    ID: number;
    Number: number,
    Date_Time: Date | null
    Cost: number,
    Employee_ID: number;
    Employee: EmployeeInterface;
    Label_ID: number;
    Label: LabelsInterface;
    Stock_ID: number;
    Stock: StocksInterface;

}
