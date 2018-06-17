import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState } from '../store'
import * as Crawl from '../modules/crawl'
import * as Common from '../modules/common'

import { Header, SubHeader, Block,
    CrawlButton, CrawlProgress, CrawlState } from './Commons'

// container component

type Props = Crawl.State & Crawl.Actions & Common.CrawlerState

class Component extends React.Component<Props> {
    componentDidMount() {
        this.props.init()
    }

    render() {
        let props = this.props
        let cs = props.cs

        return (
            <div>
                <Header>
                    Crawl <small>クローラ管理</small>
                </Header>
                <SubHeader>
                    Controll <small>クローラの操作</small>
                </SubHeader>
                <Block>
                    <CrawlProgress now={cs.now} max={cs.max} active={cs.run} interval={cs.interval} />
                    <CrawlButton isStop={!cs.run} isLoading={props.isLoading} onClick={() => props.startCrawl(!cs.run)} />
                </Block>
                <SubHeader>
                    State <small>クローラの状態</small>
                </SubHeader>
                <Block>
                    <CrawlState active={cs.run} now={cs.now} max={cs.max} id={cs.qiitaId}
                        description={cs.description} score={cs.score} message={cs.message} />
                </Block>
            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.crawl, ...appState.crawler }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {            
            fetch('/api/crawl', {
                method: 'GET',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                dispatch(Common.crawlerActions.updateCrawlState(r))
            })
        },
        updateCrawlState: (cs: Common.CrawlState) => {            
            dispatch(Common.crawlerActions.updateCrawlState(cs))            
        },
        startCrawl: (start: boolean) => {
            dispatch(Common.crawlerActions.updateCrawlLoading(true))

            let method = start ? 'POST' : 'PUT'
            fetch('/api/crawl', {
                method: method,
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                dispatch(Common.crawlerActions.updateCrawlState(r))
                dispatch(Common.crawlerActions.updateCrawlLoading(false))
            })
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Component)
