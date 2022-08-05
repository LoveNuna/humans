/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface HumanFeeBalance {
  index?: string;
  chainName?: string;
  balance?: string;
  decimal?: string;
}

export interface HumanKeysignVoteData {
  index?: string;
  txHash?: string;
  pubKey?: string;
  voter?: string;
  txTime?: string;
}

export type HumanMsgKeysignVoteResponse = object;

export interface HumanMsgObservationVoteResponse {
  code?: string;
  msg?: string;
}

export interface HumanMsgRequestTransactionResponse {
  code?: string;
  msg?: string;
}

export type HumanMsgUpdateBalanceResponse = object;

export interface HumanObserveVote {
  index?: string;
  creator?: string;
  txhash?: string;
  from?: string;
  to?: string;
  amount?: string;
  txtime?: string;
  chainId?: string;
  used?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type HumanParams = object;

export interface HumanPoolBalanap {
  index?: string;
  transactionData?: string;
  originChain?: string;
  originAddress?: string;
  targetChain?: string;
  targetAddress?: string;
  amount?: string;
  time?: string;
  creator?: string;
  status?: string;
  confirmedBlockHash?: string;
  signedKey?: string;
  fee?: string;
}

export interface HumanPoolBalance {
  index?: string;
  chainName?: string;
  balance?: string;
  decimal?: string;
}

export interface HumanQueryAllFeeBalanceResponse {
  feeBalance?: HumanFeeBalance[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HumanQueryAllKeysignVoteDataResponse {
  keysignVoteData?: HumanKeysignVoteData[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HumanQueryAllObserveVoteResponse {
  observeVote?: HumanObserveVote[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HumanQueryAllPoolBalanapResponse {
  poolBalanap?: HumanPoolBalanap[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HumanQueryAllPoolBalanceResponse {
  poolBalance?: HumanPoolBalance[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface HumanQueryGetFeeBalanceResponse {
  feeBalance?: HumanFeeBalance;
}

export interface HumanQueryGetKeysignVoteDataResponse {
  keysignVoteData?: HumanKeysignVoteData;
}

export interface HumanQueryGetObserveVoteResponse {
  observeVote?: HumanObserveVote;
}

export interface HumanQueryGetPoolBalanapResponse {
  poolBalanap?: HumanPoolBalanap;
}

export interface HumanQueryGetPoolBalanceResponse {
  poolBalance?: HumanPoolBalance;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface HumanQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: HumanParams;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title human/fee_balance.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeBalanceAll
   * @summary Queries a list of FeeBalance items.
   * @request GET:/human/human/fee_balance
   */
  queryFeeBalanceAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HumanQueryAllFeeBalanceResponse, RpcStatus>({
      path: `/human/human/fee_balance`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeBalance
   * @summary Queries a FeeBalance by index.
   * @request GET:/human/human/fee_balance/{index}
   */
  queryFeeBalance = (index: string, params: RequestParams = {}) =>
    this.request<HumanQueryGetFeeBalanceResponse, RpcStatus>({
      path: `/human/human/fee_balance/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryKeysignVoteDataAll
   * @summary Queries a list of KeysignVoteData items.
   * @request GET:/human/human/keysign_vote_data
   */
  queryKeysignVoteDataAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HumanQueryAllKeysignVoteDataResponse, RpcStatus>({
      path: `/human/human/keysign_vote_data`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryKeysignVoteData
   * @summary Queries a KeysignVoteData by index.
   * @request GET:/human/human/keysign_vote_data/{index}
   */
  queryKeysignVoteData = (index: string, params: RequestParams = {}) =>
    this.request<HumanQueryGetKeysignVoteDataResponse, RpcStatus>({
      path: `/human/human/keysign_vote_data/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryObserveVoteAll
   * @summary Queries a list of ObserveVote items.
   * @request GET:/human/human/observe_vote
   */
  queryObserveVoteAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HumanQueryAllObserveVoteResponse, RpcStatus>({
      path: `/human/human/observe_vote`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryObserveVote
   * @summary Queries a ObserveVote by index.
   * @request GET:/human/human/observe_vote/{index}
   */
  queryObserveVote = (index: string, params: RequestParams = {}) =>
    this.request<HumanQueryGetObserveVoteResponse, RpcStatus>({
      path: `/human/human/observe_vote/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/human/human/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<HumanQueryParamsResponse, RpcStatus>({
      path: `/human/human/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolBalanapAll
   * @summary Queries a list of PoolBalanap items.
   * @request GET:/human/human/pool_balanap
   */
  queryPoolBalanapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HumanQueryAllPoolBalanapResponse, RpcStatus>({
      path: `/human/human/pool_balanap`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolBalanap
   * @summary Queries a PoolBalanap by index.
   * @request GET:/human/human/pool_balanap/{index}
   */
  queryPoolBalanap = (index: string, params: RequestParams = {}) =>
    this.request<HumanQueryGetPoolBalanapResponse, RpcStatus>({
      path: `/human/human/pool_balanap/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolBalanceAll
   * @summary Queries a list of PoolBalance items.
   * @request GET:/human/human/pool_balance
   */
  queryPoolBalanceAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<HumanQueryAllPoolBalanceResponse, RpcStatus>({
      path: `/human/human/pool_balance`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolBalance
   * @summary Queries a PoolBalance by index.
   * @request GET:/human/human/pool_balance/{index}
   */
  queryPoolBalance = (index: string, params: RequestParams = {}) =>
    this.request<HumanQueryGetPoolBalanceResponse, RpcStatus>({
      path: `/human/human/pool_balance/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
