import { Review_pointInterface } from "./IReview_point";
import { Type_CommentInterface } from "./Type_comment";


export interface SeparationInterface {
    ID?: number,
    
    Review_point_ID?: number;
    Review_point?: Review_pointInterface;
    
    Type_Com_ID?: number;
    Type_Comment?: Type_CommentInterface;
    
    Payment_ID?: number;
    // Payment?: PaymentInterface;

    Date_Now?: Date,
    Bought_now?: number,
}