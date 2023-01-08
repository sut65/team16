import { EmployeeInterface } from "../IEmployee";
import { MemberInterface } from "../theerawat/IMember";

export interface Shopping_CartInterface {
    ID?: number,
    Total?: number,

    Mem_Tel?: string,
    Member?: MemberInterface
    Employee_ID?: number;
    Employee?: EmployeeInterface;
}