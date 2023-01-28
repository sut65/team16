import { EmployeeInterface } from "../IEmployee";
import { LabelsInterface } from "./ILabel";
import { StocksInterface } from "./IStock";
//ระบบชั้นวางสินค้า
export interface IShelving{
<<<<<<< HEAD
    ID: number;
    Employee_ID: number;
    Employee: EmployeeInterface;
    Label_ID: number;
    Label: LabelsInterface;
    Stock_ID: number;
    Stock: StocksInterface;
=======
    ID: string,
    Employee_ID: number,
    Employee: EmployeeInterface,
    Label_ID: number,
    Label: LabelsInterface,
    Stock_ID: number,
    Stock: StocksInterface
>>>>>>> issue-126

}