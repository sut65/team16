import { EmployeeInterface } from "../IEmployee";
import { ReasonInterface } from "./IReason";
// import { ShelvingInterface } from "./IShelving";


export interface SeparationInterface {
    ID?: number,
    
    Employee_ID?: number;
    Employee?: EmployeeInterface;
    
    Reason_ID?: number;
    Reason?: ReasonInterface;
    
    Shelving_ID?: number;
    // Shelving?: ShelvingInterface;

    Date_Out?: Date,
    Amount?: number,
    Status?: string,
}