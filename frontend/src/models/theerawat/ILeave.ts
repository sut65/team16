import { EmployeeInterface } from "../IEmployee";
import { L_TypeInterface } from "./IL_Type";
import { SectionInterface } from "./ISection";

export interface MemberInterface {
    ID?: number,
    Doc_Reason?: string,
    Doc_DateS?: Date | null,
    Doc_DateE?: Date | null,
    Doc_Cont?: string,

    Employee_ID?: number;
    Employee?: EmployeeInterface;
    Section_ID?: number;
    Section?: SectionInterface;
    L_Type_ID?: number;
    L_Type?: L_TypeInterface;
}