import { EmployeeInterface } from "../IEmployee";
import { MemberInterface } from "../theerawat/IMember";
import { StatusInterface } from "./IStatus";


export interface CartInterface {
    ID?: number,
    Total?: number,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Member_ID?: number;
    Member?: MemberInterface;
    Status_ID?: number;
    Status?: StatusInterface ;
}