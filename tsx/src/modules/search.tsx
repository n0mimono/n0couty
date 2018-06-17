import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';
import { UserSummary } from './common_user'
import * as Common from './common'

// constants

interface ScoredUser {
    id: number
    score: number
}

// actions
const actionCreator = actionCreatorFactory()

export const actions = {
    init: actionCreator('SEARCH_INIT'),
    updateUserList: actionCreator<ScoredUser[]>('SEARCH_UPDATE_USER_LIST'),
    updateLoading: actionCreator<boolean>('SEARCH_UPDATE_LOADING'),
    updateSummary: actionCreator<UserSummary[]>('SEARCH_UPDATE_SUMMARY'),
    updateFormValue: actionCreator<string>('SEARCH_UPDATE_FORM_VALUE'),
}

export interface Actions {
    init: () => Action<undefined>
    getUserList: (query: string) => Action<string>
    onFormChange: (query: string) => Action<string>
}

// states
export interface State {
    isLoading: boolean
    userList: ScoredUser[]
    users: UserSummary[]
    formValue: string
}

const initState: State = {
    isLoading: false,
    userList: [],
    users: [],
    formValue: "",
}

// reducers
export const Reducer = reducerWithInitialState(initState)
    .case(actions.init, (state) => {
        return {
            ...state,
        }
    })
    .case(actions.updateUserList, (state, userList) => {
        return {
            ...state,
            userList: [ ...userList ],
        }
    })
    .case(actions.updateLoading, (state, isLoading) => {
        return {
            ...state,
            isLoading: isLoading,
        }
    })
    .case(actions.updateSummary, (state, users) => {
        return {
            ...state,
            users: [ ...users ],
        }
    })
    .case(actions.updateFormValue, (state, value) => {
        return {
            ...state,
            formValue: value,
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