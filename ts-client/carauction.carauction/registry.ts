import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddBid } from "./types/carauction/carauction/tx";
import { MsgMakeAuction } from "./types/carauction/carauction/tx";
import { MsgEndAuction } from "./types/carauction/carauction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/carauction.carauction.MsgAddBid", MsgAddBid],
    ["/carauction.carauction.MsgMakeAuction", MsgMakeAuction],
    ["/carauction.carauction.MsgEndAuction", MsgEndAuction],
    
];

export { msgTypes }