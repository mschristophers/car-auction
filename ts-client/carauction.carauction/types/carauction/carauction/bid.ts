/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "carauction.carauction";

export interface Bid {
  creator: string;
  id: number;
  auctionID: number;
  bidPrice: number;
}

function createBaseBid(): Bid {
  return { creator: "", id: 0, auctionID: 0, bidPrice: 0 };
}

export const Bid = {
  encode(message: Bid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.auctionID !== 0) {
      writer.uint32(24).uint64(message.auctionID);
    }
    if (message.bidPrice !== 0) {
      writer.uint32(32).uint64(message.bidPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Bid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBid();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.auctionID = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.bidPrice = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Bid {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      auctionID: isSet(object.auctionID) ? Number(object.auctionID) : 0,
      bidPrice: isSet(object.bidPrice) ? Number(object.bidPrice) : 0,
    };
  },

  toJSON(message: Bid): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.auctionID !== undefined && (obj.auctionID = Math.round(message.auctionID));
    message.bidPrice !== undefined && (obj.bidPrice = Math.round(message.bidPrice));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Bid>, I>>(object: I): Bid {
    const message = createBaseBid();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.auctionID = object.auctionID ?? 0;
    message.bidPrice = object.bidPrice ?? 0;
    return message;
  },
};

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