import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';

// constants


// actions
const actionCreator = actionCreatorFactory()

export const actions = {
    init: actionCreator('STATS_INIT'),
    updateSummary: actionCreator<any>('STATS_UPDATE_SUMMARY'),
    updateLoadingSummary: actionCreator<boolean>('STATS_UPDATE_LOADING_SUMMARY')
}

export interface Actions {
    init: () => Action<undefined>
    getSummary: (update: boolean) => Action<boolean>
}

// states
export interface State {
    summary: any
    isLoadingSummary: any
}

const initState: State = {
    summary: undefined,
    isLoadingSummary: true,
}

// reducers
export const Reducer = reducerWithInitialState(initState)
    .case(actions.init, (state) => {
        return {
            ...state,
        }
    })
    .case(actions.updateSummary, (state, summary) => {
        return {
            ...state,
            summary: { ...summary }
        }
    })
    .case(actions.updateLoadingSummary, (state, isLoading) => {
        return {
            ...state,
            isLoadingSummary: isLoading,
        }
    })
