// @generated by protobuf-ts 2.8.3
// @generated from protobuf file "backend/services.proto" (package "backend", syntax proto3)
// tslint:disable
import { HealthcheckService } from "./services";
import type { PingServerSideStreamingResponse } from "./messages";
import type { PingServerSideStreamingRequest } from "./messages";
import type { PingUnaryResponse } from "./messages";
import type { PingUnaryRequest } from "./messages";
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { BackendService } from "./services";
import type { VerifyTokenResponse } from "./messages";
import type { VerifyTokenRequest } from "./messages";
import type { ListContestsResponse } from "./messages";
import type { ListContestsRequest } from "./messages";
import type { PostLoginResponse } from "./messages";
import type { PostLoginRequest } from "./messages";
import type { ListSubmitsResponse } from "./messages";
import type { ListSubmitsRequest } from "./messages";
import type { GetSubmitResponse } from "./messages";
import type { GetSubmitRequest } from "./messages";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { PostSubmitResponse } from "./messages";
import type { PostSubmitRequest } from "./messages";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetRankingResponse } from "./messages";
import type { GetRankingRequest } from "./messages";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service backend.BackendService
 */
export interface IBackendServiceClient {
    /**
     * @generated from protobuf rpc: GetRanking(backend.GetRankingRequest) returns (backend.GetRankingResponse);
     */
    getRanking(input: GetRankingRequest, options?: RpcOptions): UnaryCall<GetRankingRequest, GetRankingResponse>;
    /**
     * @generated from protobuf rpc: PostSubmit(backend.PostSubmitRequest) returns (backend.PostSubmitResponse);
     */
    postSubmit(input: PostSubmitRequest, options?: RpcOptions): UnaryCall<PostSubmitRequest, PostSubmitResponse>;
    /**
     * @generated from protobuf rpc: GetSubmit(backend.GetSubmitRequest) returns (stream backend.GetSubmitResponse);
     */
    getSubmit(input: GetSubmitRequest, options?: RpcOptions): ServerStreamingCall<GetSubmitRequest, GetSubmitResponse>;
    /**
     * @generated from protobuf rpc: ListSubmits(backend.ListSubmitsRequest) returns (backend.ListSubmitsResponse);
     */
    listSubmits(input: ListSubmitsRequest, options?: RpcOptions): UnaryCall<ListSubmitsRequest, ListSubmitsResponse>;
    /**
     * @generated from protobuf rpc: PostLogin(backend.PostLoginRequest) returns (backend.PostLoginResponse);
     */
    postLogin(input: PostLoginRequest, options?: RpcOptions): UnaryCall<PostLoginRequest, PostLoginResponse>;
    /**
     * @generated from protobuf rpc: ListContests(backend.ListContestsRequest) returns (backend.ListContestsResponse);
     */
    listContests(input: ListContestsRequest, options?: RpcOptions): UnaryCall<ListContestsRequest, ListContestsResponse>;
    /**
     * @generated from protobuf rpc: VerifyToken(backend.VerifyTokenRequest) returns (backend.VerifyTokenResponse);
     */
    verifyToken(input: VerifyTokenRequest, options?: RpcOptions): UnaryCall<VerifyTokenRequest, VerifyTokenResponse>;
}
/**
 * @generated from protobuf service backend.BackendService
 */
export class BackendServiceClient implements IBackendServiceClient, ServiceInfo {
    typeName = BackendService.typeName;
    methods = BackendService.methods;
    options = BackendService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetRanking(backend.GetRankingRequest) returns (backend.GetRankingResponse);
     */
    getRanking(input: GetRankingRequest, options?: RpcOptions): UnaryCall<GetRankingRequest, GetRankingResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetRankingRequest, GetRankingResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: PostSubmit(backend.PostSubmitRequest) returns (backend.PostSubmitResponse);
     */
    postSubmit(input: PostSubmitRequest, options?: RpcOptions): UnaryCall<PostSubmitRequest, PostSubmitResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<PostSubmitRequest, PostSubmitResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetSubmit(backend.GetSubmitRequest) returns (stream backend.GetSubmitResponse);
     */
    getSubmit(input: GetSubmitRequest, options?: RpcOptions): ServerStreamingCall<GetSubmitRequest, GetSubmitResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetSubmitRequest, GetSubmitResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListSubmits(backend.ListSubmitsRequest) returns (backend.ListSubmitsResponse);
     */
    listSubmits(input: ListSubmitsRequest, options?: RpcOptions): UnaryCall<ListSubmitsRequest, ListSubmitsResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListSubmitsRequest, ListSubmitsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: PostLogin(backend.PostLoginRequest) returns (backend.PostLoginResponse);
     */
    postLogin(input: PostLoginRequest, options?: RpcOptions): UnaryCall<PostLoginRequest, PostLoginResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<PostLoginRequest, PostLoginResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListContests(backend.ListContestsRequest) returns (backend.ListContestsResponse);
     */
    listContests(input: ListContestsRequest, options?: RpcOptions): UnaryCall<ListContestsRequest, ListContestsResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListContestsRequest, ListContestsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: VerifyToken(backend.VerifyTokenRequest) returns (backend.VerifyTokenResponse);
     */
    verifyToken(input: VerifyTokenRequest, options?: RpcOptions): UnaryCall<VerifyTokenRequest, VerifyTokenResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<VerifyTokenRequest, VerifyTokenResponse>("unary", this._transport, method, opt, input);
    }
}
/**
 * @generated from protobuf service backend.HealthcheckService
 */
export interface IHealthcheckServiceClient {
    /**
     * @generated from protobuf rpc: PingUnary(backend.PingUnaryRequest) returns (backend.PingUnaryResponse);
     */
    pingUnary(input: PingUnaryRequest, options?: RpcOptions): UnaryCall<PingUnaryRequest, PingUnaryResponse>;
    /**
     * @generated from protobuf rpc: PingServerSideStreaming(backend.PingServerSideStreamingRequest) returns (stream backend.PingServerSideStreamingResponse);
     */
    pingServerSideStreaming(input: PingServerSideStreamingRequest, options?: RpcOptions): ServerStreamingCall<PingServerSideStreamingRequest, PingServerSideStreamingResponse>;
}
/**
 * @generated from protobuf service backend.HealthcheckService
 */
export class HealthcheckServiceClient implements IHealthcheckServiceClient, ServiceInfo {
    typeName = HealthcheckService.typeName;
    methods = HealthcheckService.methods;
    options = HealthcheckService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: PingUnary(backend.PingUnaryRequest) returns (backend.PingUnaryResponse);
     */
    pingUnary(input: PingUnaryRequest, options?: RpcOptions): UnaryCall<PingUnaryRequest, PingUnaryResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<PingUnaryRequest, PingUnaryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: PingServerSideStreaming(backend.PingServerSideStreamingRequest) returns (stream backend.PingServerSideStreamingResponse);
     */
    pingServerSideStreaming(input: PingServerSideStreamingRequest, options?: RpcOptions): ServerStreamingCall<PingServerSideStreamingRequest, PingServerSideStreamingResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<PingServerSideStreamingRequest, PingServerSideStreamingResponse>("serverStreaming", this._transport, method, opt, input);
    }
}
