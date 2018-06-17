import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState } from '../store'
import * as Home from '../modules/home'
import { Header, SubHeader, Block } from './Commons'

// container component

type Props = Home.State & Home.Actions

class Component extends React.Component<Props> {
    componentDidMount() {
        this.props.init()
    }

    render() {
        return (
            <div>
                <Header>
                    Home <small>ホーム</small>
                </Header>
                <SubHeader>
                    Crawl <small>クローラ管理</small>
                </SubHeader>
                <Block>
                    クローラの管理を行います。
                </Block>
                <SubHeader>
                    Users <small>ユーザ一覧</small>
                </SubHeader>
                <Block>
                    ユーザの一覧を表示します。
                </Block>
                <SubHeader>
                    Search <small>検索</small>
                </SubHeader>
                <Block>
                    ユーザを検索します。
                </Block>
                <SubHeader>
                    Statistics <small>統計</small>
                </SubHeader>
                <Block>
                    統計情報を表示します。
                </Block>
            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.home }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {
            dispatch(Home.actions.init())
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Component)
