import { actionCreatorFactory } from 'typescript-fsa'
import { reducerWithInitialState } from 'typescript-fsa-reducers'
import { Action } from 'typescript-fsa';
import { LOCATION_CHANGE } from 'react-router-redux';
import { NodeStringDecoder } from 'string_decoder';

// constants
export const PathIndeces = ["/", "/crawl", "/users", "/users/who", "/search", "/stats"]

export interface CrawlState {
    run: boolean
    stop: boolean
    max: number
    now: number
    qiitaId: string
    description: string
    score: number
    interval: number
    message: string
}

export interface UserStarCheck {
    userId: number
    on: boolean
}

// actions
const actionCreator = actionCreatorFactory()

export const pageActions = {
    locationChange: actionCreator<any>(LOCATION_CHANGE)
}
export interface PageActions {
    move: (path: string) => Action<string>
}

export const crawlerActions = {
    updateCrawlState: actionCreator<CrawlState>('COMMON_UPDATE_CRAWL_STATE'),
    updateCrawlLoading: actionCreator<boolean>('COMMON_UPDATE_CRAWL_LOADING')
}
export interface RtCrawlerActions {
    updateCrawlState: (cs: CrawlState) => Action<undefined>
}

export const scoutActions = {
    starChange: actionCreator<UserStarCheck>('COMMON_STAR_CHANGE')
}
export interface ScoutActions {
    checkStar: (check: UserStarCheck) => Action<UserStarCheck>
}

// states
export interface PageState {
    path: string
}
const initPageState: PageState = {
    path: "/"
}

export interface CrawlerState {
    cs: CrawlState
    isLoading: boolean
}
const initCrawlerState: CrawlerState = {
    cs: {
        run: false,
        stop: true,
        max: 100,
        now: 0,
        qiitaId: '',
        description: '',
        score: 0,
        interval: 0,
        message: '',
    },
    isLoading: false
}

// props
export type PageProps = PageActions & PageState
export type RtCrawlerProps = RtCrawlerActions & CrawlerState
export type ScouterProps = ScoutActions

// reducers
export const PageReducer = reducerWithInitialState(initPageState)
    .case(pageActions.locationChange, (state, payload) => {
        return {
            ...state, path: payload.pathname
        }
    })

export const CrawlerReducer = reducerWithInitialState(initCrawlerState)
    .case(crawlerActions.updateCrawlState, (state, cs) => {
        return {
            ...state, cs: { ...cs }
        }
    })
    .case(crawlerActions.updateCrawlLoading, (state, isLoading) => {
        return {
            ...state, isLoading: isLoading
        }
    })
