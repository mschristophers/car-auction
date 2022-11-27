import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgMakeAuction } from "./types/carauction/carauction/tx";
import { MsgAddBid } from "./types/carauction/carauction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/carauction.carauction.MsgMakeAuction", MsgMakeAuction],
    ["/carauction.carauction.MsgAddBid", MsgAddBid],
    
];

export { msgTypes }