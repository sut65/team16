import { OvertimeInterface } from "./IOvertime"
import { Working_timeInterface } from "./IWorking_time"
import { DutyInterface } from "./IDuty"
import { EmployeeInterface } from "../IEmployee"

export interface Record_employee_leave {
    ID?: number;
    Employee?: EmployeeInterface;
    Employee_ID?: number;     // foreignkey.ID?
    Duty?: DutyInterface; 
    Duty_ID?: number;     // foreignkey.ID?
    Working_time?: Working_timeInterface;
    Working_time_ID?: number; // foreignkey.ID?
    Overtime?: OvertimeInterface;
    Overtime_ID?: number;  // foreignkey.ID?
    Time_Out?: Date | null;
	Number_Em?: string; 
}