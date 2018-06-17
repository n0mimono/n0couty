import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState } from '../store'
import * as Detail from '../modules/detail'
import { Header, SubHeader, Block, Loading,
    UserPage, UserPageItems } from './Commons'

import * as utility from '../utility'
import * as Common from '../modules/common'

// container component

type Props = Detail.State & Detail.Actions & Common.ScouterProps

class Component extends React.Component<Props> {
    componentDidMount() {
        this.props.init()
    }

    render() {
        let props = this.props
    
        return (
            <div>
                <Header>
                    User Detail <small>ユーザ詳細</small>
                </Header>
                <SubHeader>
                    Summary <small>基本情報</small>
                </SubHeader>
                <Block>
                    {
                        props.user.isLoading ? <Loading /> :
                        <UserPage user={props.user.summary} onStar={props.checkStar} />
                    }
                </Block>
                <SubHeader>
                    Items <small>投稿</small>
                </SubHeader>
                <Block>
                    {
                        props.item.isLoading ? <Loading /> :
                        <UserPageItems item={props.item.summary}
                            qiitaId={props.user.summary.user.qiitaId}
                        />
                    }
                </Block>

            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.detail }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {
            dispatch(Detail.actions.init())

            let parser = new URL(window.location.href)
            let q = parser.pathname.split('/')[2]

            // fetch
            let query = 'id=' + q
            fetch('/api/users?' + query, {
                method: 'GET',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                let summary = r.summary
                dispatch(Detail.actions.updateUser({ isLoading: false, summary: summary}))
            })
            .catch(e => console.log(e))

            // loading
            dispatch(Detail.actions.updateUser({ isLoading: true }))

            // fetch
            fetch('/api/users/items?' + query, {
                method: 'GET',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                let summary = r.summary
                dispatch(Detail.actions.updateItem({ isLoading: false, summary: summary}))
            })
            .catch(e => console.log(e))

            // loading
            dispatch(Detail.actions.updateItem({ isLoading: true }))
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
