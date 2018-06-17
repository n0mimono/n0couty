import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';
import { UserSummary, initUserSummary, UserItemSummary, initUserItemSummary  } from './common_user'
import * as Common from './common'

// constants
export interface User {
    isLoading: boolean
    summary?: UserSummary
}
export interface Item {
    isLoading: boolean
    summary?: UserItemSummary
}

// actions
const actionCreator = actionCreatorFactory()

export const actions = {
    init: actionCreator('DETAIL_INIT'),
    updateUser: actionCreator<User>('USER_UPDATE_USER'),
    updateItem: actionCreator<Item>('USER_UPDATE_ITEM'),
}

export interface Actions {
    init: () => Action<undefined>
}

// states
export interface State {
    user: User
    item: Item
}

const initState: State = {
    user: {
        summary: initUserSummary,
        isLoading: true,
    },
    item: {
        summary: initUserItemSummary,
        isLoading: true,
    }
}

// reducers
export const Reducer = reducerWithInitialState(initState)
    .case(actions.init, (state) => {
        return {
            ...state,
        }
    })
    .case(actions.updateUser, (state, user) => {
        return {
            ...state,
            user: {
                isLoading: user.isLoading,
                summary: user.summary == undefined ? state.user.summary : { ...user.summary }
            }
        }
    })
    .case(actions.updateItem, (state, item) => {
        return {
            ...state,
            item: {
                isLoading: item.isLoading,
                summary: item.summary == undefined ? state.item.summary : { ...item.summary }
            }
        }
    })
    .case(Common.scoutActions.starChange, (state, check) => {
        let user = { ... state.user }
        user.summary.scout.starred = check.on
        return {
            ...state,
            user: user,
        }
    })