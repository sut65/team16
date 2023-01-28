import { PaymentInterface } from "../Natthapon/IPayment";
import { Review_pointInterface } from "./IReview_point";
import { Type_CommentInterface } from "./Type_comment";


export interface CommentInterface {
    ID?: number,
    
    Comments?: string,
    Review_point_ID?: number;
    Review_point?: Review_pointInterface;
    
    Type_Com_ID?: number;
    Type_Com?: Type_CommentInterface;
    
    Payment_ID?: number;
    Payment?: PaymentInterface;

    Date_Now?: Date | null,
    Bought_now?: number,
}