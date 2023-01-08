import { EmployeeInterface } from "../IEmployee";
import { GenderInterface } from "./IGender";
import { LevelInterface } from "./ILevel";

export interface MemberInterface {
    ID?: number,
    Mem_Name?: string,
    Mem_Age?: number,
    Mem_Tel?: string,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Gender_ID?: number;
    Gender?: GenderInterface;
    Level_ID?: number;
    Level?: LevelInterface;
}