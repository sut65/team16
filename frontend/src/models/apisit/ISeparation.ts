import { EmployeeInterface } from "../IEmployee";
import { ReasonInterface } from "./IReason";
import { IShelving } from "./../methas/IShelving";


export interface SeparationInterface {
    ID?: number,
    
    Employee_ID?: number;
    Employee?: EmployeeInterface;
    
    Reason_ID?: number;
    Reason?: ReasonInterface;
    
    Shelving_ID?: number;
    Shelving?: IShelving;

    Date_Out?: Date | null,
    Amount?: number,
    Status?: string,
}