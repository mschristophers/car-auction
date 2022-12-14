/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "auction.auction";

export interface Auction {
  creator: string;
  id: number;
  name: string;
  initialPrice: string;
  duration: number;
  createdAt: number;
  currentHighestBidID: number;
  highestBidIsPresent: boolean;
  ended: boolean;
}

function createBaseAuction(): Auction {
  return {
    creator: "",
    id: 0,
    name: "",
    initialPrice: "",
    duration: 0,
    createdAt: 0,
    currentHighestBidID: 0,
    highestBidIsPresent: false,
    ended: false,
  };
}

export const Auction = {
  encode(message: Auction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.initialPrice !== "") {
      writer.uint32(34).string(message.initialPrice);
    }
    if (message.duration !== 0) {
      writer.uint32(40).uint64(message.duration);
    }
    if (message.createdAt !== 0) {
      writer.uint32(48).int64(message.createdAt);
    }
    if (message.currentHighestBidID !== 0) {
      writer.uint32(56).uint64(message.currentHighestBidID);
    }
    if (message.highestBidIsPresent === true) {
      writer.uint32(64).bool(message.highestBidIsPresent);
    }
    if (message.ended === true) {
      writer.uint32(72).bool(message.ended);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Auction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuction();
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
          message.name = reader.string();
          break;
        case 4:
          message.initialPrice = reader.string();
          break;
        case 5:
          message.duration = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.createdAt = longToNumber(reader.int64() as Long);
          break;
        case 7:
          message.currentHighestBidID = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.highestBidIsPresent = reader.bool();
          break;
        case 9:
          message.ended = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Auction {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      initialPrice: isSet(object.initialPrice) ? String(object.initialPrice) : "",
      duration: isSet(object.duration) ? Number(object.duration) : 0,
      createdAt: isSet(object.createdAt) ? Number(object.createdAt) : 0,
      currentHighestBidID: isSet(object.currentHighestBidID) ? Number(object.currentHighestBidID) : 0,
      highestBidIsPresent: isSet(object.highestBidIsPresent) ? Boolean(object.highestBidIsPresent) : false,
      ended: isSet(object.ended) ? Boolean(object.ended) : false,
    };
  },

  toJSON(message: Auction): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.name !== undefined && (obj.name = message.name);
    message.initialPrice !== undefined && (obj.initialPrice = message.initialPrice);
    message.duration !== undefined && (obj.duration = Math.round(message.duration));
    message.createdAt !== undefined && (obj.createdAt = Math.round(message.createdAt));
    message.currentHighestBidID !== undefined && (obj.currentHighestBidID = Math.round(message.currentHighestBidID));
    message.highestBidIsPresent !== undefined && (obj.highestBidIsPresent = message.highestBidIsPresent);
    message.ended !== undefined && (obj.ended = message.ended);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Auction>, I>>(object: I): Auction {
    const message = createBaseAuction();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    message.initialPrice = object.initialPrice ?? "";
    message.duration = object.duration ?? 0;
    message.createdAt = object.createdAt ?? 0;
    message.currentHighestBidID = object.currentHighestBidID ?? 0;
    message.highestBidIsPresent = object.highestBidIsPresent ?? false;
    message.ended = object.ended ?? false;
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
