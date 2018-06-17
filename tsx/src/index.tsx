import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider, connect } from 'react-redux'
import { Route, Switch } from 'react-router'
import { ConnectedRouter } from 'react-router-redux'
import store, { history } from './store'

import { Navigator } from './containers/Commons'
import Home from './containers/Home'
import Crawl from './containers/Crawl'
import Users from './containers/Users'
import Detail from './containers/Detail'
import Search from './containers/Search'
import Stats from './containers/Stats'

ReactDOM.render(
    <Provider store={store}>
        <ConnectedRouter history={history}>
            <Navigator>
                <Switch>
                    <Route exact path="/"><Home /></Route>
                    <Route exact path="/crawl"><Crawl /></Route>
                    <Route exact path="/users" ><Users /></Route>
                    <Route exact path="/users/:userId"><Detail /></Route>
                    <Route exact path="/search"><Search /></Route>
                    <Route exact path="/stats"><Stats /></Route>
                    <Route><h1>Not Found</h1></Route>
                </Switch>
            </Navigator>
        </ConnectedRouter>
    </Provider>,
    document.querySelector('.app')
)
