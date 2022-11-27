/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "carauction.carauction";

export interface MsgMakeAuction {
  creator: string;
  name: string;
  initialPrice: number;
  minIncrement: number;
}

export interface MsgMakeAuctionResponse {
  id: number;
}

export interface MsgAddBid {
  creator: string;
  auctionID: number;
  bidPrice: number;
}

export interface MsgAddBidResponse {
  id: number;
}

export interface MsgEndAuction {
  creator: string;
  auctionID: number;
  bidID: number;
}

export interface MsgEndAuctionResponse {
}

function createBaseMsgMakeAuction(): MsgMakeAuction {
  return { creator: "", name: "", initialPrice: 0, minIncrement: 0 };
}

export const MsgMakeAuction = {
  encode(message: MsgMakeAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.initialPrice !== 0) {
      writer.uint32(24).uint64(message.initialPrice);
    }
    if (message.minIncrement !== 0) {
      writer.uint32(32).uint64(message.minIncrement);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMakeAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMakeAuction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.initialPrice = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.minIncrement = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMakeAuction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      initialPrice: isSet(object.initialPrice) ? Number(object.initialPrice) : 0,
      minIncrement: isSet(object.minIncrement) ? Number(object.minIncrement) : 0,
    };
  },

  toJSON(message: MsgMakeAuction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.initialPrice !== undefined && (obj.initialPrice = Math.round(message.initialPrice));
    message.minIncrement !== undefined && (obj.minIncrement = Math.round(message.minIncrement));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMakeAuction>, I>>(object: I): MsgMakeAuction {
    const message = createBaseMsgMakeAuction();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.initialPrice = object.initialPrice ?? 0;
    message.minIncrement = object.minIncrement ?? 0;
    return message;
  },
};

function createBaseMsgMakeAuctionResponse(): MsgMakeAuctionResponse {
  return { id: 0 };
}

export const MsgMakeAuctionResponse = {
  encode(message: MsgMakeAuctionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMakeAuctionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMakeAuctionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMakeAuctionResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgMakeAuctionResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMakeAuctionResponse>, I>>(object: I): MsgMakeAuctionResponse {
    const message = createBaseMsgMakeAuctionResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgAddBid(): MsgAddBid {
  return { creator: "", auctionID: 0, bidPrice: 0 };
}

export const MsgAddBid = {
  encode(message: MsgAddBid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.auctionID !== 0) {
      writer.uint32(16).uint64(message.auctionID);
    }
    if (message.bidPrice !== 0) {
      writer.uint32(24).uint64(message.bidPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBid();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.auctionID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.bidPrice = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddBid {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      auctionID: isSet(object.auctionID) ? Number(object.auctionID) : 0,
      bidPrice: isSet(object.bidPrice) ? Number(object.bidPrice) : 0,
    };
  },

  toJSON(message: MsgAddBid): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.auctionID !== undefined && (obj.auctionID = Math.round(message.auctionID));
    message.bidPrice !== undefined && (obj.bidPrice = Math.round(message.bidPrice));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBid>, I>>(object: I): MsgAddBid {
    const message = createBaseMsgAddBid();
    message.creator = object.creator ?? "";
    message.auctionID = object.auctionID ?? 0;
    message.bidPrice = object.bidPrice ?? 0;
    return message;
  },
};

function createBaseMsgAddBidResponse(): MsgAddBidResponse {
  return { id: 0 };
}

export const MsgAddBidResponse = {
  encode(message: MsgAddBidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddBidResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgAddBidResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBidResponse>, I>>(object: I): MsgAddBidResponse {
    const message = createBaseMsgAddBidResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgEndAuction(): MsgEndAuction {
  return { creator: "", auctionID: 0, bidID: 0 };
}

export const MsgEndAuction = {
  encode(message: MsgEndAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.auctionID !== 0) {
      writer.uint32(16).uint64(message.auctionID);
    }
    if (message.bidID !== 0) {
      writer.uint32(24).uint64(message.bidID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEndAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEndAuction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.auctionID = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.bidID = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEndAuction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      auctionID: isSet(object.auctionID) ? Number(object.auctionID) : 0,
      bidID: isSet(object.bidID) ? Number(object.bidID) : 0,
    };
  },

  toJSON(message: MsgEndAuction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.auctionID !== undefined && (obj.auctionID = Math.round(message.auctionID));
    message.bidID !== undefined && (obj.bidID = Math.round(message.bidID));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEndAuction>, I>>(object: I): MsgEndAuction {
    const message = createBaseMsgEndAuction();
    message.creator = object.creator ?? "";
    message.auctionID = object.auctionID ?? 0;
    message.bidID = object.bidID ?? 0;
    return message;
  },
};

function createBaseMsgEndAuctionResponse(): MsgEndAuctionResponse {
  return {};
}

export const MsgEndAuctionResponse = {
  encode(_: MsgEndAuctionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEndAuctionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEndAuctionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgEndAuctionResponse {
    return {};
  },

  toJSON(_: MsgEndAuctionResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEndAuctionResponse>, I>>(_: I): MsgEndAuctionResponse {
    const message = createBaseMsgEndAuctionResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  MakeAuction(request: MsgMakeAuction): Promise<MsgMakeAuctionResponse>;
  AddBid(request: MsgAddBid): Promise<MsgAddBidResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EndAuction(request: MsgEndAuction): Promise<MsgEndAuctionResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.MakeAuction = this.MakeAuction.bind(this);
    this.AddBid = this.AddBid.bind(this);
    this.EndAuction = this.EndAuction.bind(this);
  }
  MakeAuction(request: MsgMakeAuction): Promise<MsgMakeAuctionResponse> {
    const data = MsgMakeAuction.encode(request).finish();
    const promise = this.rpc.request("carauction.carauction.Msg", "MakeAuction", data);
    return promise.then((data) => MsgMakeAuctionResponse.decode(new _m0.Reader(data)));
  }

  AddBid(request: MsgAddBid): Promise<MsgAddBidResponse> {
    const data = MsgAddBid.encode(request).finish();
    const promise = this.rpc.request("carauction.carauction.Msg", "AddBid", data);
    return promise.then((data) => MsgAddBidResponse.decode(new _m0.Reader(data)));
  }

  EndAuction(request: MsgEndAuction): Promise<MsgEndAuctionResponse> {
    const data = MsgEndAuction.encode(request).finish();
    const promise = this.rpc.request("carauction.carauction.Msg", "EndAuction", data);
    return promise.then((data) => MsgEndAuctionResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
