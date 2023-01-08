import { EmployeeInterface } from "../IEmployee";
import { LabelsInterface } from "./ILabel";
import { StocksInterface } from "./IStock";

export interface IShelving{
    ID: string;
    Employee_ID: number;
    Employee: EmployeeInterface;
    Label_ID: number;
    Label: LabelsInterface;
    Stock_ID: number;
    Stock: StocksInterface;


}