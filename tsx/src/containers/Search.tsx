import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState, history } from '../store'
import * as Search from '../modules/search'
import { Header, SubHeader, Block, Loading,
    UserList, SearchForm } from './Commons'

import * as utility from '../utility'
import { UserSummary } from '../modules/common_user';
import * as Common from '../modules/common'

// container component

type Props = Search.State & Search.Actions & Common.ScouterProps

class Component extends React.Component<Props> {
    componentDidMount() {
        this.props.init()
    }

    render() {
        let props = this.props
        return (
            <div>
                <Header>
                    Search <small>検索</small>
                </Header>
                <SubHeader>
                    Request <small>要件</small>
                </SubHeader>
                <Block>
                    <SearchForm isLoading={props.isLoading} onClick={q => this.props.getUserList(q)} />
                </Block>
                <SubHeader>
                    Results <small>検索結果</small>
                </SubHeader>
                <Block>
                    {
                        props.isLoading ? <Loading /> :
                        <UserList onDetail={id => history.push('/users/' + id)}
                            users={utility.toArray(props.users)}
                            onStar={props.checkStar} /> 
                    }
                </Block>
            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.search }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {
            dispatch(Search.actions.init())
        },
        getUserList: (q: Search.SearchQuery) => {
            // fetch
            let query = 'max=20' + '&simple=' + q.simple
            fetch('/api/ml/word/similarity_users?' + query, {
                method: 'POST',
                credentials: "same-origin",
                headers: { 'content-type': 'application/json' },
                body: JSON.stringify({ doc: q.doc }),
            })
            .then(r => r.json())
            .then(r => {
                console.log(r)
                let userList = r.similars
                dispatch(Search.actions.updateUserList(utility.toArray(userList)))
                dispatch(Search.actions.updateLoading(false))

                let users: UserSummary[] = []
                for (let i = 0; i < userList.length; i++) {
                    let query = 'id=' + userList[i].user_id
                    fetch('/api/users?' + query, {
                        method: 'GET',
                        credentials: "same-origin",
                    })
                    .then(r => r.json())
                    .then(r => {
                        let summary = r.summary
                        users = [ ...users, summary]
                        dispatch(Search.actions.updateSummary(utility.toArray(users)))
                    })
                    .catch(e => console.log(e))
                }
            })
            .catch(e => {
                dispatch(Search.actions.updateLoading(false))
            })

            // loading
            dispatch(Search.actions.updateLoading(true))
        },
        checkStar: (check: Common.UserStarCheck) => {
            let query = 'id=' + check.userId + '&star=' + check.on
            fetch('/api/users/scout?' + query, {
                method: 'POST',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                dispatch(Common.scoutActions.starChange(check))
            })
            .catch(e => console.log(e))
        },
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Component)
