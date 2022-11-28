import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddBid } from "./types/carauction/carauction/tx";
import { MsgEndAuction } from "./types/carauction/carauction/tx";
import { MsgMakeAuction } from "./types/carauction/carauction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/carauction.carauction.MsgAddBid", MsgAddBid],
    ["/carauction.carauction.MsgEndAuction", MsgEndAuction],
    ["/carauction.carauction.MsgMakeAuction", MsgMakeAuction],
    
];

export { msgTypes }