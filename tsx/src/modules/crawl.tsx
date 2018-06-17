import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';

import * as Common from './common'

// constants

// actions
const actionCreator = actionCreatorFactory()

export const actions = {
}

export interface Actions {
    init: () => Action<undefined>
    startCrawl: (start: boolean) => Action<boolean>
    updateCrawlState: (cs: Common.CrawlState) => Action<undefined>
}

// states
export interface State {
}

const initState: State = {
}

// reducers
export const Reducer = reducerWithInitialState(initState)
