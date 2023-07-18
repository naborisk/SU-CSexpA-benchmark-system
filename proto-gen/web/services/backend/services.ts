// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "services/backend/services.proto" (package "backend", syntax proto3)
// tslint:disable
import { PingServerSideStreamingResponse } from "./messages";
import { PingServerSideStreamingRequest } from "./messages";
import { PingUnaryResponse } from "./messages";
import { PingUnaryRequest } from "./messages";
import { PostLoginResponse } from "./messages";
import { PostLoginRequest } from "./messages";
import { VerifyTokenResponse } from "./messages";
import { VerifyTokenRequest } from "./messages";
import { GetRankingResponse } from "./messages";
import { GetRankingRequest } from "./messages";
import { UpdateContestResponse } from "./messages";
import { UpdateContestRequest } from "./messages";
import { GetContestResponse } from "./messages";
import { GetContestRequest } from "./messages";
import { ListContestsResponse } from "./messages";
import { ListContestsRequest } from "./messages";
import { CreateContestResponse } from "./messages";
import { CreateContestRequest } from "./messages";
import { GetLatestSubmitResponse } from "./messages";
import { GetLatestSubmitRequest } from "./messages";
import { ListSubmitsResponse } from "./messages";
import { ListSubmitsRequest } from "./messages";
import { GetSubmitResponse } from "./messages";
import { GetSubmitRequest } from "./messages";
import { PostSubmitResponse } from "./messages";
import { PostSubmitRequest } from "./messages";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service backend.BackendService
 */
export const BackendService = new ServiceType("backend.BackendService", [
    { name: "PostSubmit", options: {}, I: PostSubmitRequest, O: PostSubmitResponse },
    { name: "GetSubmit", serverStreaming: true, options: {}, I: GetSubmitRequest, O: GetSubmitResponse },
    { name: "ListSubmits", options: {}, I: ListSubmitsRequest, O: ListSubmitsResponse },
    { name: "GetLatestSubmit", options: {}, I: GetLatestSubmitRequest, O: GetLatestSubmitResponse },
    { name: "CreateContest", options: {}, I: CreateContestRequest, O: CreateContestResponse },
    { name: "ListContests", options: {}, I: ListContestsRequest, O: ListContestsResponse },
    { name: "GetContest", options: {}, I: GetContestRequest, O: GetContestResponse },
    { name: "UpdateContest", options: {}, I: UpdateContestRequest, O: UpdateContestResponse },
    { name: "GetRanking", options: {}, I: GetRankingRequest, O: GetRankingResponse },
    { name: "VerifyToken", options: {}, I: VerifyTokenRequest, O: VerifyTokenResponse },
    { name: "PostLogin", options: {}, I: PostLoginRequest, O: PostLoginResponse }
]);
/**
 * @generated ServiceType for protobuf service backend.HealthcheckService
 */
export const HealthcheckService = new ServiceType("backend.HealthcheckService", [
    { name: "PingUnary", options: {}, I: PingUnaryRequest, O: PingUnaryResponse },
    { name: "PingServerSideStreaming", serverStreaming: true, options: {}, I: PingServerSideStreamingRequest, O: PingServerSideStreamingResponse }
]);
