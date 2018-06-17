import { createStore, combineReducers, applyMiddleware } from 'redux'
import { routerReducer, routerMiddleware } from 'react-router-redux'
import { createHashHistory, createBrowserHistory } from 'history'

import * as Home from './modules/home'
import * as Crawl from './modules/crawl'
import * as Users from './modules/users'
import * as Detail from './modules/detail'
import * as Search from './modules/search'
import * as Stats from './modules/stats'
import * as Common from './modules/common'

export type AppState = {
    home: Home.State,
    crawl: Crawl.State,
    users: Users.State,
    detail: Detail.State,
    search: Search.State,
    stats: Stats.State,

    page: Common.PageState,
    crawler: Common.CrawlerState,

    routing: any,
}

export const history = createBrowserHistory()
const middleware = routerMiddleware(history)

const store = createStore(
    combineReducers<AppState>({
        home: Home.Reducer,
        crawl: Crawl.Reducer,
        users: Users.Reducer,
        detail: Detail.Reducer,
        search: Search.Reducer,
        stats: Stats.Reducer,

        page: Common.PageReducer,
        crawler: Common.CrawlerReducer,

        routing: routerReducer,
    }),
    applyMiddleware(middleware)
)

export default store
