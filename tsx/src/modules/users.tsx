import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';
import { UserSummary } from './common_user'
import * as Common from './common'

// constants

export interface Query {
    page: string
    onlyStarred: boolean
}

export interface SummaryList {
    users: UserSummary[]
    page: string
    prev: string
    next: string
}

// actions
const actionCreator = actionCreatorFactory()

export const actions = {
    init: actionCreator('USERS_INIT'),
    updateSummary: actionCreator<SummaryList>('USERS_UPDATE_SUMMARY'),
    updateLoading: actionCreator<boolean>('USERS_UPDATE_LOADING'),
    updateOnlyStarred: actionCreator<boolean>('USERS_UPDATE_ONLY_STARRED'),
}

export interface Actions {
    init: () => Action<undefined>
    getList: (query: Query) => Action<Query>
}

// states
export interface State {
    isLoading: boolean
    users: UserSummary[]
    page: string
    prev: string
    next: string
    onlyStarred: boolean
}

const initState: State = {
    isLoading: true,
    users: [],
    page: "1",
    prev: "0",
    next: "0",
    onlyStarred: false,
}

// reducers
export const Reducer = reducerWithInitialState(initState)
    .case(actions.init, (state) => {
        return {
            ...state,
        }
    })
    .case(actions.updateSummary, (state, sl) => {
        return {
            ...state,
            users: [ ...sl.users ],
            page: sl.page,
            prev: sl.prev,
            next: sl.next,
        }
    })
    .case(actions.updateLoading, (state, isLoading) => {
        return {
            ...state,
            isLoading: isLoading,
        }
    })
    .case(actions.updateOnlyStarred, (state, onlyStarred) => {
        return {
            ...state,
            onlyStarred: onlyStarred,
        }
    })
    .case(Common.scoutActions.starChange, (state, check) => {
        let users = [ ...state.users ]
        for (let i = 0; i < users.length; i++) {
            if (users[i].user.id != check.userId) continue
            let user = { ... users[i] }
            user.scout.starred = check.on
            users[i] = user
        }

        return {
            ...state,
            users: users,
        }
    })
