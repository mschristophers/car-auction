/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "auction.auction";

export interface MsgCreateAuction {
  creator: string;
  name: string;
  initialPrice: string;
  duration: number;
}

export interface MsgCreateAuctionResponse {
  id: number;
}

export interface MsgPlaceBid {
  creator: string;
  auctionID: number;
  bidPrice: string;
}

export interface MsgPlaceBidResponse {
  id: number;
}

export interface MsgFinalizeAuction {
  creator: string;
  auctionID: number;
}

export interface MsgFinalizeAuctionResponse {
  id: number;
  finalPrice: string;
}

function createBaseMsgCreateAuction(): MsgCreateAuction {
  return { creator: "", name: "", initialPrice: "", duration: 0 };
}

export const MsgCreateAuction = {
  encode(message: MsgCreateAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.initialPrice !== "") {
      writer.uint32(26).string(message.initialPrice);
    }
    if (message.duration !== 0) {
      writer.uint32(32).uint64(message.duration);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAuction();
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
          message.initialPrice = reader.string();
          break;
        case 4:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAuction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      initialPrice: isSet(object.initialPrice) ? String(object.initialPrice) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
    };
  },

  toJSON(message: MsgCreateAuction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.initialPrice !== undefined && (obj.initialPrice = message.initialPrice);
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAuction>, I>>(object: I): MsgCreateAuction {
    const message = createBaseMsgCreateAuction();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.initialPrice = object.initialPrice ?? "";
    message.duration = object.duration ?? 0;
    return message;
  },
};

function createBaseMsgCreateAuctionResponse(): MsgCreateAuctionResponse {
  return { id: 0 };
}

export const MsgCreateAuctionResponse = {
  encode(message: MsgCreateAuctionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAuctionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAuctionResponse();
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

  fromJSON(object: any): MsgCreateAuctionResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateAuctionResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAuctionResponse>, I>>(object: I): MsgCreateAuctionResponse {
    const message = createBaseMsgCreateAuctionResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgPlaceBid(): MsgPlaceBid {
  return { creator: "", auctionID: 0, bidPrice: "" };
}

export const MsgPlaceBid = {
  encode(message: MsgPlaceBid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.auctionID !== 0) {
      writer.uint32(16).uint64(message.auctionID);
    }
    if (message.bidPrice !== "") {
      writer.uint32(26).string(message.bidPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceBid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceBid();
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
          message.bidPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlaceBid {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      auctionID: isSet(object.auctionID) ? Number(object.auctionID) : 0,
      bidPrice: isSet(object.bidPrice) ? String(object.bidPrice) : "",
    };
  },

  toJSON(message: MsgPlaceBid): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.auctionID !== undefined && (obj.auctionID = Math.round(message.auctionID));
    message.bidPrice !== undefined && (obj.bidPrice = message.bidPrice);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlaceBid>, I>>(object: I): MsgPlaceBid {
    const message = createBaseMsgPlaceBid();
    message.creator = object.creator ?? "";
    message.auctionID = object.auctionID ?? 0;
    message.bidPrice = object.bidPrice ?? "";
    return message;
  },
};

function createBaseMsgPlaceBidResponse(): MsgPlaceBidResponse {
  return { id: 0 };
}

export const MsgPlaceBidResponse = {
  encode(message: MsgPlaceBidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceBidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceBidResponse();
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

  fromJSON(object: any): MsgPlaceBidResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgPlaceBidResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlaceBidResponse>, I>>(object: I): MsgPlaceBidResponse {
    const message = createBaseMsgPlaceBidResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgFinalizeAuction(): MsgFinalizeAuction {
  return { creator: "", auctionID: 0 };
}

export const MsgFinalizeAuction = {
  encode(message: MsgFinalizeAuction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.auctionID !== 0) {
      writer.uint32(16).uint64(message.auctionID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeAuction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeAuction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.auctionID = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFinalizeAuction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      auctionID: isSet(object.auctionID) ? Number(object.auctionID) : 0,
    };
  },

  toJSON(message: MsgFinalizeAuction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.auctionID !== undefined && (obj.auctionID = Math.round(message.auctionID));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeAuction>, I>>(object: I): MsgFinalizeAuction {
    const message = createBaseMsgFinalizeAuction();
    message.creator = object.creator ?? "";
    message.auctionID = object.auctionID ?? 0;
    return message;
  },
};

function createBaseMsgFinalizeAuctionResponse(): MsgFinalizeAuctionResponse {
  return { id: 0, finalPrice: "" };
}

export const MsgFinalizeAuctionResponse = {
  encode(message: MsgFinalizeAuctionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.finalPrice !== "") {
      writer.uint32(18).string(message.finalPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFinalizeAuctionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFinalizeAuctionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.finalPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFinalizeAuctionResponse {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      finalPrice: isSet(object.finalPrice) ? String(object.finalPrice) : "",
    };
  },

  toJSON(message: MsgFinalizeAuctionResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.finalPrice !== undefined && (obj.finalPrice = message.finalPrice);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFinalizeAuctionResponse>, I>>(object: I): MsgFinalizeAuctionResponse {
    const message = createBaseMsgFinalizeAuctionResponse();
    message.id = object.id ?? 0;
    message.finalPrice = object.finalPrice ?? "";
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateAuction(request: MsgCreateAuction): Promise<MsgCreateAuctionResponse>;
  PlaceBid(request: MsgPlaceBid): Promise<MsgPlaceBidResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FinalizeAuction(request: MsgFinalizeAuction): Promise<MsgFinalizeAuctionResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAuction = this.CreateAuction.bind(this);
    this.PlaceBid = this.PlaceBid.bind(this);
    this.FinalizeAuction = this.FinalizeAuction.bind(this);
  }
  CreateAuction(request: MsgCreateAuction): Promise<MsgCreateAuctionResponse> {
    const data = MsgCreateAuction.encode(request).finish();
    const promise = this.rpc.request("auction.auction.Msg", "CreateAuction", data);
    return promise.then((data) => MsgCreateAuctionResponse.decode(new _m0.Reader(data)));
  }

  PlaceBid(request: MsgPlaceBid): Promise<MsgPlaceBidResponse> {
    const data = MsgPlaceBid.encode(request).finish();
    const promise = this.rpc.request("auction.auction.Msg", "PlaceBid", data);
    return promise.then((data) => MsgPlaceBidResponse.decode(new _m0.Reader(data)));
  }

  FinalizeAuction(request: MsgFinalizeAuction): Promise<MsgFinalizeAuctionResponse> {
    const data = MsgFinalizeAuction.encode(request).finish();
    const promise = this.rpc.request("auction.auction.Msg", "FinalizeAuction", data);
    return promise.then((data) => MsgFinalizeAuctionResponse.decode(new _m0.Reader(data)));
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
