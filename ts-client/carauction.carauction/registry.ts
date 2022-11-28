import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgEndAuction } from "./types/carauction/carauction/tx";
import { MsgMakeAuction } from "./types/carauction/carauction/tx";
import { MsgAddBid } from "./types/carauction/carauction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/carauction.carauction.MsgEndAuction", MsgEndAuction],
    ["/carauction.carauction.MsgMakeAuction", MsgMakeAuction],
    ["/carauction.carauction.MsgAddBid", MsgAddBid],
    
];

export { msgTypes }