import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { AppState } from '../store'
import * as Stats from '../modules/stats'
import { Header, SubHeader, Block, Loading,
    StatsSummary } from './Commons'

// container component

type Props = Stats.State & Stats.Actions

class Component extends React.Component<Props> {
    componentDidMount() {
        this.props.init()
        this.props.getSummary(false)
    }

    render() {
        let props = this.props
        return (
            <div>
                <Header>
                    Statistics <small>統計</small>
                </Header>
                <SubHeader>
                    Summary <small>サマリー</small>
                </SubHeader>
                <Block>
                    {
                        props.isLoadingSummary ? <Loading /> :
                        <StatsSummary
                            summary={props.summary}
                            isLoading={props.isLoadingSummary}
                            onUpdate={() => props.getSummary(true)}
                         />
                    }
                </Block>
            </div>
        )
    }
}

// connect

function mapStateToProps(appState: AppState) {
    return { ...appState.stats }
}

function mapDispatchToProps(dispatch: Dispatch<void>) {
    return {
        init: () => {
            dispatch(Stats.actions.init())
        },
        getSummary: (update: boolean) => {
            // fetch
            fetch('/api/ml/summary', {
                method: update ? 'POST' : 'GET',
                credentials: "same-origin",
            })
            .then(r => r.json())
            .then(r => {
                let summary = r
                dispatch(Stats.actions.updateSummary(summary))
                dispatch(Stats.actions.updateLoadingSummary(false))
            })
            .catch(e => console.log(e))
          
            // loading
            dispatch(Stats.actions.updateLoadingSummary(true))
        },
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Component)
