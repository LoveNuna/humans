// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestTransaction } from "./types/human/tx";
import { MsgObservationVote } from "./types/human/tx";
import { MsgUpdateBalance } from "./types/human/tx";
import { MsgKeysignVote } from "./types/human/tx";


const types = [
  ["/human.human.MsgRequestTransaction", MsgRequestTransaction],
  ["/human.human.MsgObservationVote", MsgObservationVote],
  ["/human.human.MsgUpdateBalance", MsgUpdateBalance],
  ["/human.human.MsgKeysignVote", MsgKeysignVote],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRequestTransaction: (data: MsgRequestTransaction): EncodeObject => ({ typeUrl: "/human.human.MsgRequestTransaction", value: MsgRequestTransaction.fromPartial( data ) }),
    msgObservationVote: (data: MsgObservationVote): EncodeObject => ({ typeUrl: "/human.human.MsgObservationVote", value: MsgObservationVote.fromPartial( data ) }),
    msgUpdateBalance: (data: MsgUpdateBalance): EncodeObject => ({ typeUrl: "/human.human.MsgUpdateBalance", value: MsgUpdateBalance.fromPartial( data ) }),
    msgKeysignVote: (data: MsgKeysignVote): EncodeObject => ({ typeUrl: "/human.human.MsgKeysignVote", value: MsgKeysignVote.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
