// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "services/benchmark-service/messages.proto" (package "benchmark", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Status } from "../backend/resources";
/**
 * @generated from protobuf message benchmark.ExecuteRequest
 */
export interface ExecuteRequest {
    /**
     * @generated from protobuf field: repeated benchmark.Task tasks = 1;
     */
    tasks: Task[];
    /**
     * @generated from protobuf field: string group_id = 2;
     */
    groupId: string; // for logging
    /**
     * @generated from protobuf field: string contest_slug = 3;
     */
    contestSlug: string;
}
/**
 * @generated from protobuf message benchmark.Task
 */
export interface Task {
    /**
     * @generated from protobuf field: benchmark.HttpRequest request = 1;
     */
    request?: HttpRequest;
    /**
     * @generated from protobuf field: int32 thread_num = 6;
     */
    threadNum: number; // the number of threads for a task
    /**
     * @generated from protobuf field: int32 attempt_count = 7;
     */
    attemptCount: number; // the count of attempting for a task
}
/**
 * @generated from protobuf message benchmark.ExecuteResponse
 */
export interface ExecuteResponse {
    /**
     * @generated from protobuf field: bool ok = 1;
     */
    ok: boolean;
    /**
     * @generated from protobuf field: optional string error_message = 2;
     */
    errorMessage?: string; // if ok is false, this field is set
    /**
     * @generated from protobuf field: int64 time_elapsed = 3;
     */
    timeElapsed: bigint; // in milliseconds
    /**
     * @generated from protobuf field: int32 total_requests = 4;
     */
    totalRequests: number;
    /**
     * @generated from protobuf field: int32 requests_per_second = 5;
     */
    requestsPerSecond: number;
    /**
     * @generated from protobuf field: benchmark.Task task = 6;
     */
    task?: Task;
    /**
     * @generated from protobuf field: backend.Status status = 7;
     */
    status: Status;
}
/**
 * @generated from protobuf message benchmark.HttpRequest
 */
export interface HttpRequest {
    /**
     * @generated from protobuf field: string url = 1;
     */
    url: string; // e.g.) http://10.255.255.255/endpoint
    /**
     * @generated from protobuf field: benchmark.HttpMethod method = 2;
     */
    method: HttpMethod;
    /**
     * @generated from protobuf field: string content_type = 4;
     */
    contentType: string;
    /**
     * @generated from protobuf field: string body = 5;
     */
    body: string;
}
/**
 * @generated from protobuf enum benchmark.HttpMethod
 */
export enum HttpMethod {
    /**
     * @generated from protobuf enum value: GET = 0;
     */
    GET = 0,
    /**
     * @generated from protobuf enum value: POST = 1;
     */
    POST = 1,
    /**
     * @generated from protobuf enum value: PUT = 2;
     */
    PUT = 2,
    /**
     * @generated from protobuf enum value: DELETE = 3;
     */
    DELETE = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class ExecuteRequest$Type extends MessageType<ExecuteRequest> {
    constructor() {
        super("benchmark.ExecuteRequest", [
            { no: 1, name: "tasks", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Task },
            { no: 2, name: "group_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "contest_slug", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ExecuteRequest>): ExecuteRequest {
        const message = { tasks: [], groupId: "", contestSlug: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ExecuteRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExecuteRequest): ExecuteRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated benchmark.Task tasks */ 1:
                    message.tasks.push(Task.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string group_id */ 2:
                    message.groupId = reader.string();
                    break;
                case /* string contest_slug */ 3:
                    message.contestSlug = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ExecuteRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated benchmark.Task tasks = 1; */
        for (let i = 0; i < message.tasks.length; i++)
            Task.internalBinaryWrite(message.tasks[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string group_id = 2; */
        if (message.groupId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.groupId);
        /* string contest_slug = 3; */
        if (message.contestSlug !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.contestSlug);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message benchmark.ExecuteRequest
 */
export const ExecuteRequest = new ExecuteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Task$Type extends MessageType<Task> {
    constructor() {
        super("benchmark.Task", [
            { no: 1, name: "request", kind: "message", T: () => HttpRequest },
            { no: 6, name: "thread_num", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "attempt_count", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Task>): Task {
        const message = { threadNum: 0, attemptCount: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Task>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Task): Task {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* benchmark.HttpRequest request */ 1:
                    message.request = HttpRequest.internalBinaryRead(reader, reader.uint32(), options, message.request);
                    break;
                case /* int32 thread_num */ 6:
                    message.threadNum = reader.int32();
                    break;
                case /* int32 attempt_count */ 7:
                    message.attemptCount = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Task, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* benchmark.HttpRequest request = 1; */
        if (message.request)
            HttpRequest.internalBinaryWrite(message.request, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* int32 thread_num = 6; */
        if (message.threadNum !== 0)
            writer.tag(6, WireType.Varint).int32(message.threadNum);
        /* int32 attempt_count = 7; */
        if (message.attemptCount !== 0)
            writer.tag(7, WireType.Varint).int32(message.attemptCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message benchmark.Task
 */
export const Task = new Task$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ExecuteResponse$Type extends MessageType<ExecuteResponse> {
    constructor() {
        super("benchmark.ExecuteResponse", [
            { no: 1, name: "ok", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "error_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "time_elapsed", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "total_requests", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "requests_per_second", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "task", kind: "message", T: () => Task },
            { no: 7, name: "status", kind: "enum", T: () => ["backend.Status", Status] }
        ]);
    }
    create(value?: PartialMessage<ExecuteResponse>): ExecuteResponse {
        const message = { ok: false, timeElapsed: 0n, totalRequests: 0, requestsPerSecond: 0, status: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ExecuteResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExecuteResponse): ExecuteResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool ok */ 1:
                    message.ok = reader.bool();
                    break;
                case /* optional string error_message */ 2:
                    message.errorMessage = reader.string();
                    break;
                case /* int64 time_elapsed */ 3:
                    message.timeElapsed = reader.int64().toBigInt();
                    break;
                case /* int32 total_requests */ 4:
                    message.totalRequests = reader.int32();
                    break;
                case /* int32 requests_per_second */ 5:
                    message.requestsPerSecond = reader.int32();
                    break;
                case /* benchmark.Task task */ 6:
                    message.task = Task.internalBinaryRead(reader, reader.uint32(), options, message.task);
                    break;
                case /* backend.Status status */ 7:
                    message.status = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ExecuteResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool ok = 1; */
        if (message.ok !== false)
            writer.tag(1, WireType.Varint).bool(message.ok);
        /* optional string error_message = 2; */
        if (message.errorMessage !== undefined)
            writer.tag(2, WireType.LengthDelimited).string(message.errorMessage);
        /* int64 time_elapsed = 3; */
        if (message.timeElapsed !== 0n)
            writer.tag(3, WireType.Varint).int64(message.timeElapsed);
        /* int32 total_requests = 4; */
        if (message.totalRequests !== 0)
            writer.tag(4, WireType.Varint).int32(message.totalRequests);
        /* int32 requests_per_second = 5; */
        if (message.requestsPerSecond !== 0)
            writer.tag(5, WireType.Varint).int32(message.requestsPerSecond);
        /* benchmark.Task task = 6; */
        if (message.task)
            Task.internalBinaryWrite(message.task, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* backend.Status status = 7; */
        if (message.status !== 0)
            writer.tag(7, WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message benchmark.ExecuteResponse
 */
export const ExecuteResponse = new ExecuteResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class HttpRequest$Type extends MessageType<HttpRequest> {
    constructor() {
        super("benchmark.HttpRequest", [
            { no: 1, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "method", kind: "enum", T: () => ["benchmark.HttpMethod", HttpMethod] },
            { no: 4, name: "content_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "body", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<HttpRequest>): HttpRequest {
        const message = { url: "", method: 0, contentType: "", body: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<HttpRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: HttpRequest): HttpRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string url */ 1:
                    message.url = reader.string();
                    break;
                case /* benchmark.HttpMethod method */ 2:
                    message.method = reader.int32();
                    break;
                case /* string content_type */ 4:
                    message.contentType = reader.string();
                    break;
                case /* string body */ 5:
                    message.body = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: HttpRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string url = 1; */
        if (message.url !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.url);
        /* benchmark.HttpMethod method = 2; */
        if (message.method !== 0)
            writer.tag(2, WireType.Varint).int32(message.method);
        /* string content_type = 4; */
        if (message.contentType !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.contentType);
        /* string body = 5; */
        if (message.body !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.body);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message benchmark.HttpRequest
 */
export const HttpRequest = new HttpRequest$Type();
