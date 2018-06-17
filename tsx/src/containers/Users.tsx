import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState, history } from '../store'
import * as Users from '../modules/users'
import { Header, SubHeader, Block, Loading,
    UserList, UserListBtns } from './Commons'

import * as utility from '../utility'
import * as Common from '../modules/common'

// container component

type Props = Users.State & Users.Actions & Common.ScouterProps

class Component extends React.Component<Props> {
    componentWillMount() {
        let parser = new URL(window.location.href)

        let page = this.props.page
        if (parser.searchParams.has('page')) {
            page = parser.searchParams.get('page')
        }

        let onlyStarred = this.props.onlyStarred
        if (parser.searchParams.has('onlyStarred')) {
            onlyStarred = parser.searchParams.get('paonlyStarredge') == 'true'
        }

        this.props.getList({ page: page, onlyStarred: onlyStarred})
        this.props.init()
    }

    render() {
        let props = this.props
        return (
            <div>
                <Header>
                    Users <small>ユーザ情報</small>
                </Header>
                <SubHeader>
                    List <small>ユーザ一覧</small>
                </SubHeader>
                <Block>
                    <UserListBtns
                        current={props.onlyStarred ? 1 : 0}
                        isLoading={props.isLoading}
                        onClick={n => props.getList({ page: '1', onlyStarred: n == 1 })}
                    />
                    {
                        props.isLoading ? <Loading /> :
                        <UserList users={utility.toArray(props.users)}
                            onNext={q => props.getList({ page: q, onlyStarred: props.onlyStarred })}
                            onDetail={id => history.push('/users/' + id)}
                            page={props.page} prev={props.prev} next={props.next}
                            onStar={props.checkStar}/>
                    }
                </Block>
            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.users }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {
            dispatch(Users.actions.init())
        },
        getList: (q: Users.Query) => {
            // fetch
            let query = 'page=' + q.page + '&onlyStarred=' + q.onlyStarred
            fetch('/api/users/list?' + query, {
                method: 'GET',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                dispatch(Users.actions.updateSummary({
                    users: r.users,
                    page: r.page,
                    prev: r.prev,
                    next: r.next,
                }))
                dispatch(Users.actions.updateLoading(false))
            })
            .catch(e => console.log(e))

            // loading
            dispatch(Users.actions.updateLoading(true))

            // switch
            dispatch(Users.actions.updateOnlyStarred(q.onlyStarred))

            // repush
            //let parser = new URL(window.location.href)
            //let path = parser.pathname + '?page=' + q
            //history.push(path)
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
