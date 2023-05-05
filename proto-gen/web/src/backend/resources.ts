// @generated by protobuf-ts 2.8.3
// @generated from protobuf file "backend/resources.proto" (syntax proto3)
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
import { Timestamp } from "../google/protobuf/timestamp";
/**
 * @generated from protobuf message Contest
 */
export interface Contest {
    /**
     * @generated from protobuf field: int32 year = 3;
     */
    year: number; // primary key of contest
    /**
     * @generated from protobuf field: google.protobuf.Timestamp qualifier_start_at = 4;
     */
    qualifierStartAt?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp qualifier_end_at = 5;
     */
    qualifierEndAt?: Timestamp;
    /**
     * @generated from protobuf field: int32 qualifier_submit_limit = 6;
     */
    qualifierSubmitLimit: number;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp final_start_at = 7;
     */
    finalStartAt?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp final_end_at = 8;
     */
    finalEndAt?: Timestamp;
    /**
     * @generated from protobuf field: int32 final_submit_limit = 9;
     */
    finalSubmitLimit: number;
}
/**
 * @generated from protobuf message Group
 */
export interface Group {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * @generated from protobuf field: int32 score = 2;
     */
    score: number;
    /**
     * @generated from protobuf field: int32 year = 3;
     */
    year: number;
    /**
     * @generated from protobuf field: Role role = 4;
     */
    role: Role;
}
/**
 * @generated from protobuf message Submit
 */
export interface Submit {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: int32 group_id = 2;
     */
    groupId: number;
    /**
     * @generated from protobuf field: int32 year = 3;
     */
    year: number;
    /**
     * @generated from protobuf field: int32 score = 4;
     */
    score: number;
    /**
     * @generated from protobuf field: Language language = 5;
     */
    language: Language;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp submited_at = 6;
     */
    submitedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional google.protobuf.Timestamp completed_at = 7;
     */
    completedAt?: Timestamp; // it this field is not null, this submit is completed
    /**
     * @generated from protobuf field: repeated TaskResult task_results = 8;
     */
    taskResults: TaskResult[];
}
/**
 * @generated from protobuf message TaskResult
 */
export interface TaskResult {
    /**
     * @generated from protobuf field: int32 id = 1;
     */
    id: number;
    /**
     * @generated from protobuf field: int32 request_per_sec = 2;
     */
    requestPerSec: number;
    /**
     * @generated from protobuf field: string url = 3;
     */
    url: string;
    /**
     * @generated from protobuf field: string method = 4;
     */
    method: string;
    /**
     * @generated from protobuf field: string request_content_type = 5;
     */
    requestContentType: string;
    /**
     * @generated from protobuf field: optional string request_body = 6;
     */
    requestBody?: string;
    /**
     * @generated from protobuf field: string response_code = 7;
     */
    responseCode: string;
    /**
     * @generated from protobuf field: string response_content_type = 8;
     */
    responseContentType: string;
    /**
     * @generated from protobuf field: string response_body = 9;
     */
    responseBody: string;
    /**
     * @generated from protobuf field: int32 thread_num = 10;
     */
    threadNum: number;
    /**
     * @generated from protobuf field: int32 attempt_count = 11;
     */
    attemptCount: number;
    /**
     * @generated from protobuf field: int32 attempt_time = 12;
     */
    attemptTime: number;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp created_at = 13;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: optional google.protobuf.Timestamp deleted_at = 14;
     */
    deletedAt?: Timestamp;
    /**
     * @generated from protobuf field: optional string error_message = 15;
     */
    errorMessage?: string;
}
/**
 * @generated from protobuf enum Language
 */
export enum Language {
    /**
     * @generated from protobuf enum value: PHP = 0;
     */
    PHP = 0,
    /**
     * @generated from protobuf enum value: GO = 1;
     */
    GO = 1,
    /**
     * @generated from protobuf enum value: RUST = 2;
     */
    RUST = 2,
    /**
     * @generated from protobuf enum value: JAVASCRIPT = 3;
     */
    JAVASCRIPT = 3,
    /**
     * @generated from protobuf enum value: CSHARP = 4;
     */
    CSHARP = 4,
    /**
     * @generated from protobuf enum value: CPP = 5;
     */
    CPP = 5,
    /**
     * @generated from protobuf enum value: RUBY = 6;
     */
    RUBY = 6,
    /**
     * @generated from protobuf enum value: PYTHON = 7;
     */
    PYTHON = 7
}
/**
 * @generated from protobuf enum Role
 */
export enum Role {
    /**
     * @generated from protobuf enum value: CONTESTANT = 0;
     */
    CONTESTANT = 0,
    /**
     * @generated from protobuf enum value: GUEST = 1;
     */
    GUEST = 1
}
// @generated message type with reflection information, may provide speed optimized methods
class Contest$Type extends MessageType<Contest> {
    constructor() {
        super("Contest", [
            { no: 3, name: "year", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "qualifier_start_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "qualifier_end_at", kind: "message", T: () => Timestamp },
            { no: 6, name: "qualifier_submit_limit", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 7, name: "final_start_at", kind: "message", T: () => Timestamp },
            { no: 8, name: "final_end_at", kind: "message", T: () => Timestamp },
            { no: 9, name: "final_submit_limit", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Contest>): Contest {
        const message = { year: 0, qualifierSubmitLimit: 0, finalSubmitLimit: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Contest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Contest): Contest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 year */ 3:
                    message.year = reader.int32();
                    break;
                case /* google.protobuf.Timestamp qualifier_start_at */ 4:
                    message.qualifierStartAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.qualifierStartAt);
                    break;
                case /* google.protobuf.Timestamp qualifier_end_at */ 5:
                    message.qualifierEndAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.qualifierEndAt);
                    break;
                case /* int32 qualifier_submit_limit */ 6:
                    message.qualifierSubmitLimit = reader.int32();
                    break;
                case /* google.protobuf.Timestamp final_start_at */ 7:
                    message.finalStartAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.finalStartAt);
                    break;
                case /* google.protobuf.Timestamp final_end_at */ 8:
                    message.finalEndAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.finalEndAt);
                    break;
                case /* int32 final_submit_limit */ 9:
                    message.finalSubmitLimit = reader.int32();
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
    internalBinaryWrite(message: Contest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 year = 3; */
        if (message.year !== 0)
            writer.tag(3, WireType.Varint).int32(message.year);
        /* google.protobuf.Timestamp qualifier_start_at = 4; */
        if (message.qualifierStartAt)
            Timestamp.internalBinaryWrite(message.qualifierStartAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp qualifier_end_at = 5; */
        if (message.qualifierEndAt)
            Timestamp.internalBinaryWrite(message.qualifierEndAt, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* int32 qualifier_submit_limit = 6; */
        if (message.qualifierSubmitLimit !== 0)
            writer.tag(6, WireType.Varint).int32(message.qualifierSubmitLimit);
        /* google.protobuf.Timestamp final_start_at = 7; */
        if (message.finalStartAt)
            Timestamp.internalBinaryWrite(message.finalStartAt, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp final_end_at = 8; */
        if (message.finalEndAt)
            Timestamp.internalBinaryWrite(message.finalEndAt, writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        /* int32 final_submit_limit = 9; */
        if (message.finalSubmitLimit !== 0)
            writer.tag(9, WireType.Varint).int32(message.finalSubmitLimit);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message Contest
 */
export const Contest = new Contest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Group$Type extends MessageType<Group> {
    constructor() {
        super("Group", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "score", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "year", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "role", kind: "enum", T: () => ["Role", Role] }
        ]);
    }
    create(value?: PartialMessage<Group>): Group {
        const message = { id: "", score: 0, year: 0, role: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Group>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Group): Group {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string id */ 1:
                    message.id = reader.string();
                    break;
                case /* int32 score */ 2:
                    message.score = reader.int32();
                    break;
                case /* int32 year */ 3:
                    message.year = reader.int32();
                    break;
                case /* Role role */ 4:
                    message.role = reader.int32();
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
    internalBinaryWrite(message: Group, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string id = 1; */
        if (message.id !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.id);
        /* int32 score = 2; */
        if (message.score !== 0)
            writer.tag(2, WireType.Varint).int32(message.score);
        /* int32 year = 3; */
        if (message.year !== 0)
            writer.tag(3, WireType.Varint).int32(message.year);
        /* Role role = 4; */
        if (message.role !== 0)
            writer.tag(4, WireType.Varint).int32(message.role);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message Group
 */
export const Group = new Group$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Submit$Type extends MessageType<Submit> {
    constructor() {
        super("Submit", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "group_id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "year", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "score", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "language", kind: "enum", T: () => ["Language", Language] },
            { no: 6, name: "submited_at", kind: "message", T: () => Timestamp },
            { no: 7, name: "completed_at", kind: "message", T: () => Timestamp },
            { no: 8, name: "task_results", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => TaskResult }
        ]);
    }
    create(value?: PartialMessage<Submit>): Submit {
        const message = { id: 0, groupId: 0, year: 0, score: 0, language: 0, taskResults: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Submit>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Submit): Submit {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* int32 group_id */ 2:
                    message.groupId = reader.int32();
                    break;
                case /* int32 year */ 3:
                    message.year = reader.int32();
                    break;
                case /* int32 score */ 4:
                    message.score = reader.int32();
                    break;
                case /* Language language */ 5:
                    message.language = reader.int32();
                    break;
                case /* google.protobuf.Timestamp submited_at */ 6:
                    message.submitedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.submitedAt);
                    break;
                case /* optional google.protobuf.Timestamp completed_at */ 7:
                    message.completedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.completedAt);
                    break;
                case /* repeated TaskResult task_results */ 8:
                    message.taskResults.push(TaskResult.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: Submit, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* int32 group_id = 2; */
        if (message.groupId !== 0)
            writer.tag(2, WireType.Varint).int32(message.groupId);
        /* int32 year = 3; */
        if (message.year !== 0)
            writer.tag(3, WireType.Varint).int32(message.year);
        /* int32 score = 4; */
        if (message.score !== 0)
            writer.tag(4, WireType.Varint).int32(message.score);
        /* Language language = 5; */
        if (message.language !== 0)
            writer.tag(5, WireType.Varint).int32(message.language);
        /* google.protobuf.Timestamp submited_at = 6; */
        if (message.submitedAt)
            Timestamp.internalBinaryWrite(message.submitedAt, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* optional google.protobuf.Timestamp completed_at = 7; */
        if (message.completedAt)
            Timestamp.internalBinaryWrite(message.completedAt, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* repeated TaskResult task_results = 8; */
        for (let i = 0; i < message.taskResults.length; i++)
            TaskResult.internalBinaryWrite(message.taskResults[i], writer.tag(8, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message Submit
 */
export const Submit = new Submit$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TaskResult$Type extends MessageType<TaskResult> {
    constructor() {
        super("TaskResult", [
            { no: 1, name: "id", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "request_per_sec", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "method", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "request_content_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "request_body", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "response_code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "response_content_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "response_body", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "thread_num", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 11, name: "attempt_count", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 12, name: "attempt_time", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 13, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 14, name: "deleted_at", kind: "message", T: () => Timestamp },
            { no: 15, name: "error_message", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<TaskResult>): TaskResult {
        const message = { id: 0, requestPerSec: 0, url: "", method: "", requestContentType: "", responseCode: "", responseContentType: "", responseBody: "", threadNum: 0, attemptCount: 0, attemptTime: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<TaskResult>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TaskResult): TaskResult {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 id */ 1:
                    message.id = reader.int32();
                    break;
                case /* int32 request_per_sec */ 2:
                    message.requestPerSec = reader.int32();
                    break;
                case /* string url */ 3:
                    message.url = reader.string();
                    break;
                case /* string method */ 4:
                    message.method = reader.string();
                    break;
                case /* string request_content_type */ 5:
                    message.requestContentType = reader.string();
                    break;
                case /* optional string request_body */ 6:
                    message.requestBody = reader.string();
                    break;
                case /* string response_code */ 7:
                    message.responseCode = reader.string();
                    break;
                case /* string response_content_type */ 8:
                    message.responseContentType = reader.string();
                    break;
                case /* string response_body */ 9:
                    message.responseBody = reader.string();
                    break;
                case /* int32 thread_num */ 10:
                    message.threadNum = reader.int32();
                    break;
                case /* int32 attempt_count */ 11:
                    message.attemptCount = reader.int32();
                    break;
                case /* int32 attempt_time */ 12:
                    message.attemptTime = reader.int32();
                    break;
                case /* google.protobuf.Timestamp created_at */ 13:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* optional google.protobuf.Timestamp deleted_at */ 14:
                    message.deletedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.deletedAt);
                    break;
                case /* optional string error_message */ 15:
                    message.errorMessage = reader.string();
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
    internalBinaryWrite(message: TaskResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 id = 1; */
        if (message.id !== 0)
            writer.tag(1, WireType.Varint).int32(message.id);
        /* int32 request_per_sec = 2; */
        if (message.requestPerSec !== 0)
            writer.tag(2, WireType.Varint).int32(message.requestPerSec);
        /* string url = 3; */
        if (message.url !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.url);
        /* string method = 4; */
        if (message.method !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.method);
        /* string request_content_type = 5; */
        if (message.requestContentType !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.requestContentType);
        /* optional string request_body = 6; */
        if (message.requestBody !== undefined)
            writer.tag(6, WireType.LengthDelimited).string(message.requestBody);
        /* string response_code = 7; */
        if (message.responseCode !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.responseCode);
        /* string response_content_type = 8; */
        if (message.responseContentType !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.responseContentType);
        /* string response_body = 9; */
        if (message.responseBody !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.responseBody);
        /* int32 thread_num = 10; */
        if (message.threadNum !== 0)
            writer.tag(10, WireType.Varint).int32(message.threadNum);
        /* int32 attempt_count = 11; */
        if (message.attemptCount !== 0)
            writer.tag(11, WireType.Varint).int32(message.attemptCount);
        /* int32 attempt_time = 12; */
        if (message.attemptTime !== 0)
            writer.tag(12, WireType.Varint).int32(message.attemptTime);
        /* google.protobuf.Timestamp created_at = 13; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(13, WireType.LengthDelimited).fork(), options).join();
        /* optional google.protobuf.Timestamp deleted_at = 14; */
        if (message.deletedAt)
            Timestamp.internalBinaryWrite(message.deletedAt, writer.tag(14, WireType.LengthDelimited).fork(), options).join();
        /* optional string error_message = 15; */
        if (message.errorMessage !== undefined)
            writer.tag(15, WireType.LengthDelimited).string(message.errorMessage);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message TaskResult
 */
export const TaskResult = new TaskResult$Type();
