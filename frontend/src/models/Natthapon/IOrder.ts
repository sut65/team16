import { EmployeeInterface } from "../IEmployee";
import { MemberInterface } from "../theerawat/IMember";
import { IShelving } from "../methas/IShelving";


export interface OrderInterface {
    ID?: number,
    Quantity?: number,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Member_ID?: number
    Member?: MemberInterface
    Shelving_ID?: number;
    Shelving?: IShelving;
}