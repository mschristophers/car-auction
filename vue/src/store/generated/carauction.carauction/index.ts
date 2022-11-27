import { Client, registry, MissingWalletError } from 'car-auction-client-ts'

import { Auction } from "car-auction-client-ts/carauction.carauction/types"
import { Bid } from "car-auction-client-ts/carauction.carauction/types"
import { Params } from "car-auction-client-ts/carauction.carauction/types"


export { Auction, Bid, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				
				_Structure: {
						Auction: getStructure(Auction.fromPartial({})),
						Bid: getStructure(Bid.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: carauction.carauction initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		async sendMsgMakeAuction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CarauctionCarauction.tx.sendMsgMakeAuction({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMakeAuction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMakeAuction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddBid({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.CarauctionCarauction.tx.sendMsgAddBid({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddBid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddBid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgMakeAuction({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CarauctionCarauction.tx.msgMakeAuction({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMakeAuction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMakeAuction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddBid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CarauctionCarauction.tx.msgAddBid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddBid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddBid:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
