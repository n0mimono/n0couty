import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';

// constants


// actions
const actionCreator = actionCreatorFactory()

export const actions = {
    init: actionCreator('HOME_INIT')
}

export interface Actions {
    init: () => Action<undefined>
}

// states
export interface State {
}

const initState: State = {
}

// reducers
export const Reducer = reducerWithInitialState(initState)
    .case(actions.init, (state) => {
        return {
            ...state,
        }
    })
