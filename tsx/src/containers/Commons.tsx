import * as React from 'react'

import { Dispatch } from 'redux'
import { connect } from 'react-redux'
import { AppState, history } from '../store'

import { UserSummary, UserItemSummary, UserItemWithTags, initUserItemSummary } from '../modules/common_user'
import * as Common from '../modules/common'

import {
    Navbar, Nav, NavItem,
    Grid, Row, Col,
    PageHeader,
    Button, Form, FormGroup, FormControl, ControlLabel,
    ProgressBar, Checkbox,
    ListGroup, ListGroupItem,
    Glyphicon, Badge, Label, ButtonGroup
} from 'react-bootstrap'

import {
    XAxis, YAxis, CartesianGrid, Tooltip, Legend,
    BarChart, Bar, LineChart, Line
} from 'recharts'

import styles from '../css/style.css'
import * as utility from '../utility'

///////////////////////////////////////////////////////////////////////
// 
// Navigation
// 
///////////////////////////////////////////////////////////////////////

export const Navigator: React.SFC<{}> = (props) => {
    return (
        <div>
            <NavigatorBar />
            <Grid>
                <Row>
                    <Col xs={2}>
                        <ConnectedLeftMenu />
                    </Col>
                    <Col xs={10} className={styles.navTop}>
                        {props.children}
                    </Col>
                </Row>
            </Grid>
            <ConnectedRtCrawler />
        </div>
    )
}

const NavigatorBar: React.SFC<{}> = (props) => {
    return (
        <Navbar fixedTop>
            <Navbar.Header>
                <Navbar.Brand>
                    <a href="#">n0couty</a>
                </Navbar.Brand>
            </Navbar.Header>
        </Navbar>
    )
}

const LeftMenu: React.SFC<Common.PageProps & Common.CrawlerState> = (props) => {
    let key = Common.PathIndeces.indexOf(props.path)
    let label = props.cs.run ? "success" : "primary"

    return (
        <Nav bsStyle="pills" stacked activeKey={key} className={styles.navLeft}
            onSelect={e => props.move(Common.PathIndeces[Number(e)])}>
            <NavItem eventKey={0}>
                Home
            </NavItem>
            <NavItem eventKey={1}>
                Crawl <Label bsStyle={label}>{props.cs.now}</Label>
            </NavItem>
            <NavItem eventKey={2}>
                Users
            </NavItem>
            <NavItem eventKey={4}>
                Search
            </NavItem>
            <NavItem eventKey={5}>
                Statistics
            </NavItem>
        </Nav>
    )
}
const ConnectedLeftMenu = (() => {
    function mapStateToProps(appState: AppState) {
        return { ...appState.page, ...appState.crawler }
    }
    function mapDispatchToProps(dispatch: Dispatch<void>) {
        return {
            move: (path: string) => {
                history.push(path)
            },
        }
    }
    return connect(mapStateToProps, mapDispatchToProps)(LeftMenu)
})()

class RtCrawler extends React.Component<Common.RtCrawlerProps> {
    onClose: () => void

    componentDidMount() {
        let parser = new URL(window.location.href)
        let url = "ws://" + parser.host + "/socket/"
        
        let socket = new WebSocket(url)
        socket.onclose = ev => {
        }
        socket.onmessage = ev => {
            this.props.updateCrawlState(JSON.parse(ev.data))
        }
        this.onClose = () => {
            socket.close()
        }
    }

    componentWillUnmount() {
        this.onClose()
    }

    render() {
        return (
            <div></div>
        )
    }
}
const ConnectedRtCrawler = (() => {
    function mapStateToProps(appState: AppState) {
        return { ...appState.crawler }
    }
    function mapDispatchToProps(dispatch: Dispatch<void>) {
        return {
            updateCrawlState: (cs: Common.CrawlState) => {
                dispatch(Common.crawlerActions.updateCrawlState(cs))
            },
        }
    }
    return connect(mapStateToProps, mapDispatchToProps)(RtCrawler)
})()

export const Header: React.SFC<{}> = (props) => {
    return (
        <div>
            <PageHeader>
                {props.children}
            </PageHeader>
        </div>
    )
}

export const SubHeader: React.SFC<{}> = (props) => {
    return (
        <div>
            <h2>
                {props.children}
            </h2>
        </div>
    )
}

export const Block: React.SFC<{}> = (props) => {
    return (
        <div>
            <div className={styles.block}>
                {props.children}
            </div>
        </div>
    )
}

export const Loading: React.SFC<{}> = (props) => {
    return (
        <div className={styles.loading}>
            <div className={styles.panel}>
                <div className={styles.dot1}></div>
                <div className={styles.dot2}></div>
                <div className={styles.dot3}></div>
            </div>
        </div>
    )
}

interface SocialLinkLogoImgProps {
    serviceId: number
}

export const SocialLinkLogoImg: React.SFC<SocialLinkLogoImgProps> = (props) => {
    let htmls = [
        <i></i>,
        <i className="fab fa-2x fa-github-square"></i>,
        <i className="fab fa-2x fa-twitter-square"></i>,
        <i className="fab fa-2x fa-facebook-square"></i>,
        <i className="fab fa-2x fa-linkedin"></i>,
        <i className="fab fa-2x fa-google-plus-square"></i>
    ]
    let colors = [
        '#000000',
        '#171515',
        '#2AA3F0',
        '#4469B0',
        '#127BB4',
        '#DC524A',
    ]
    return (
        <span style={{ color: colors[props.serviceId], padding:'1px' }}>
            {htmls[props.serviceId]}
        </span>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// Crawler
// 
///////////////////////////////////////////////////////////////////////

interface ProgressProps {
    now: number
    max: number
    active: boolean
    interval: number
}

export const CrawlProgress: React.SFC<ProgressProps> = (props) => {
    let per = Math.round(1000 * props.now / props.max) / 10
    let secs = (props.interval / 1000000000) * (props.max - props.now)
    let hours = Math.round(secs / 60 / 60)

    return (
        <div className={styles.crawlProgress}>>
            <div className={styles.panel}>
                <div className={styles.left}>
                    <div>
                        {
                            props.active ? "データを更新しています... (" + props.now + "/" + props.max + ")" :
                                "データの更新を中断しています (" + props.now + "/" + props.max + ")"
                        }
                    </div>
                </div>
                <div className={styles.right}>
                    <div>
                        推定残り時間: {hours} 時間
                    </div>
                </div>
            </div>
            <ProgressBar
                bsStyle="success" striped active={props.active}
                min={0} max={props.max} now={props.now} label={per + "\%"} />
        </div>
    )
}

interface OnOffButtonProps {
    isStop: boolean
    isLoading: boolean
    onClick: () => void
}

export const CrawlButton: React.SFC<OnOffButtonProps> = (props) => {
    return (
        <div className={styles.crawlBtn}>
            <Button bsStyle="danger" disabled={props.isLoading} onClick={_ => props.onClick()}>
                {props.isStop ? "クロール開始" : "クロール停止"}
            </Button>
        </div>
    )
}

interface CrawStateProps {
    active: boolean
    now: number
    max: number
    id: string
    description: string
    score: number
    message: string
}

export const CrawlState: React.SFC<CrawStateProps> = (props) => {
    return (
        <ListGroup>
            <ListGroupItem header="active" bsStyle={props.active ? "success" : "info"}>
                {props.active ? "クローラは起動中です" : "クローラは停止中です"}
            </ListGroupItem>
            <ListGroupItem header="now">{props.now}</ListGroupItem>
            <ListGroupItem header="max">{props.max}</ListGroupItem>
            <ListGroupItem header="id">{props.id}</ListGroupItem>
            <ListGroupItem header="description">{props.description}</ListGroupItem>
            <ListGroupItem header="score">{props.score}</ListGroupItem>
            <ListGroupItem header="message">{props.message}</ListGroupItem>
        </ListGroup>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// Users
// 
///////////////////////////////////////////////////////////////////////

interface UserListProps {
    users: UserSummary[]
    page?: string
    prev?: string
    next?: string
    onNext?: (q: string) => void
    onDetail?: (id: number) => void
    onStar: (check: Common.UserStarCheck) => void
}

export const UserList: React.SFC<UserListProps> = (props) => {
    let ms = []
    for (let i = 0; i < props.users.length; i++) {
        ms.push(
            <ListGroupItem key={i}>
                <UserPage user={props.users[i]} onDetail={props.onDetail} onStar={props.onStar} />
            </ListGroupItem>
        )
    }
    return (
        <div>
            {
                props.page == undefined ? "" :
                    <UserListPage
                        page={props.page} prev={props.prev} next={props.next}
                        onNext={props.onNext} />
            }
            <ListGroup>
                {ms}
            </ListGroup>
            {
                props.page == undefined ? "" :
                    <UserListPage
                        page={props.page} prev={props.prev} next={props.next}
                        onNext={props.onNext} />
            }
        </div>
    )
}

interface UserPageProps {
    user: UserSummary
    onDetail?: (id: number) => void
    onStar: (check: Common.UserStarCheck) => void
}

export const UserPage: React.SFC<UserPageProps> = (props) => {
    let user = props.user.user
    let stat = props.user.stat
    let links = utility.toArray(props.user.links)
    let langs = utility.toArray(props.user.langStats)
    let scout = props.user.scout

    return (
        <Grid className={styles.userPage}>
            <Row className={styles.lot}>
                <Col xs={4} className={styles.header}>
                    {user.name}
                </Col>
                <Col xs={1}>
                    <Glyphicon glyph="asterisk" /> {user.id}
                </Col>
                {
                    !user.ban ? "" :
                        <Col xs={1}>
                            <Glyphicon glyph="warning-sign" /> BAN
                        </Col>
                }
                <Col xs={2}>
                    {links.map((l, i) =>
                        <a key={i} href={l.url} target="_blank"><SocialLinkLogoImg serviceId={l.serviceId} /></a>
                    )}
                </Col>
                {
                    props.onDetail == undefined ? "" :
                        <Col xs={1}>
                            <Button bsStyle="primary" onClick={_ => props.onDetail(user.id)}>
                                > Detail
                            </Button>
                        </Col>
                }
                <Col xs={1}>
                    <UserStar isStarred={scout.starred} onClick={on => props.onStar({ userId: user.id, on: on })} />
                </Col>
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row className={styles.lot}>
                <Col xs={3} xsPull={0}>
                    <Glyphicon glyph="user" /> <a href={"https://qiita.com/" + user.qiitaId} target="_blank">{user.qiitaId}</a>
                </Col>
                <Col xs={3} xsPull={1}>
                    <Glyphicon glyph="envelope" /> <a href={"mailto:" + user.mail}> {user.mail} </a>
                </Col>
                <Col xs={6} xsPull={2}>
                    <Glyphicon glyph="link" />  <a href={user.link} target="_blank"> {user.link} </a>
                </Col>
            </Row>
            <Row className={styles.lot}>
                <Col xs={3} xsPull={0}>
                    <Glyphicon glyph="oil" /> {user.organization}
                </Col>
                <Col xs={3} xsPull={1}>
                    <Glyphicon glyph="flag" /> {user.place}
                </Col>
                <Col xs={3} xsPull={2}>
                    <Glyphicon glyph="grain" /> {user.qiitaOrganization}
                </Col>
            </Row>
            <Row className={styles.lot}>
                <Col xs={9}>
                    <Glyphicon glyph="comment" /> {user.description}
                </Col>
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row className={styles.lot}>
                <Col xs={3}>
                    <Glyphicon glyph="arrow-up" /> 投稿 <Badge>{stat.items}</Badge>
                </Col>
                <Col xs={3} xsPull={1}>
                    <Glyphicon glyph="thumbs-up" /> いいね <Badge>{stat.contributions}</Badge>
                </Col>
                <Col xs={3} xsPull={2}>
                    <Glyphicon glyph="resize-small" /> フォロワー <Badge>{stat.followers}</Badge>
                </Col>
                <Col xs={3} xsPull={3}>
                    <Glyphicon glyph="resize-full" /> フォロー <Badge>{stat.followees}</Badge>
                </Col>
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row className={styles.lot}>
                <Col>
                    {langs.map((l, i) => <Badge key={i}>{l.name} {l.quantity}</Badge>)}
                </Col>
            </Row>
        </Grid>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// User List
// 
///////////////////////////////////////////////////////////////////////

interface UserListBtnsPageProps {
    current: number
    isLoading: boolean
    onClick: (n: number) => void
}

export const UserListBtns: React.SFC<UserListBtnsPageProps> = (props) => {
    let style = (cur: number, index: number) => cur == index ? "success" : "default"
    return (
        <ButtonGroup>
            <Button bsStyle={style(props.current, 0)} disabled={props.isLoading}
                onClick={_ => props.onClick(0)}>すべて</Button>
            <Button bsStyle={style(props.current, 1)} disabled={props.isLoading}
                onClick={_ => props.onClick(1)}>お気に入りのみ</Button>
        </ButtonGroup>
    )
}

interface UserListPageProps {
    page: string
    prev: string
    next: string
    onNext: (q: string) => void
}

export const UserListPage: React.SFC<UserListPageProps> = (props) => {
    let prev = props.prev == "0" ? "" : props.prev + " <"
    let next = props.next == "0" ? "" : "> " + props.next

    return (
        <div className={styles.userPager}>
            <div className={styles.item}>
                <a href="#" onClick={() => props.onNext(props.prev)}> {prev} </a>
            </div>
            <div className={styles.item}>
                {props.page}
            </div>
            <div className={styles.item}>
                <a href="#" onClick={() => props.onNext(props.next)}> {next} </a>
            </div>
        </div>
    )
}

interface UserStarProps {
    isStarred: boolean
    onClick: (on: boolean) => void
}

const UserStar: React.SFC<UserStarProps> = (props) => {
    let style = props.isStarred ? "warning" : "default"
    return (
        <div>
            <Button bsStyle={style} onClick={_ => props.onClick(!props.isStarred)}>
                <Glyphicon glyph="star" /> お気に入り
            </Button>
        </div>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// User Detail
// 
///////////////////////////////////////////////////////////////////////

interface UserPageItemsProps {
    item: UserItemSummary
    qiitaId: string
}

export const UserPageItems: React.SFC<UserPageItemsProps> = (props) => {
    let items = props.item.items
    let recents = props.item.recents
    let populars = props.item.populars

    return (
        <div>
            <h4>最近の投稿</h4>
            <ListGroup>
                {
                    recents.map((r, i) => <ListGroupItem key={i}>
                        <UserPageItemsUnit
                            qiitaId={props.qiitaId}
                            item={items.filter(item => item.body.id == r.itemId)[0]}
                        />
                    </ListGroupItem>)
                }
            </ListGroup>
            <h4>人気の投稿</h4>
            <ListGroup>
                {
                    populars.map((r, i) => <ListGroupItem key={i}>
                        <UserPageItemsUnit
                            qiitaId={props.qiitaId}
                            item={items.filter(item => item.body.id == r.itemId)[0]}
                        />
                    </ListGroupItem>)
                }
            </ListGroup>
        </div>
    )
}

interface UserPageItemsUnitProps {
    qiitaId: string
    item: UserItemWithTags
}

const UserPageItemsUnit: React.SFC<UserPageItemsUnitProps> = (props) => {
    let item = props.item
    return (
        <Grid className={styles.userPage}>
            <Row className={styles.lot}>
                <Col xs={9} className={styles.header2}>
                    <a href={"https://qiita.com/" + props.qiitaId + "/items/" + item.body.articleId} target="_blank"> {item.body.title} </a>
                </Col>
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row className={styles.lot}>
                <Col xs={3} xsPull={0}>
                    <Glyphicon glyph="thumbs-up" /> いいね <Badge>{item.body.contributions}</Badge>
                </Col>
                <Col xs={3} xsPull={1}>
                    <Glyphicon glyph="comment" /> コメント <Badge>{item.body.comments}</Badge>
                </Col>
                <Col xs={3} xsPull={2}>
                    <Glyphicon glyph="time" /> 投稿日 {item.body.date.split('T')[0]}
                </Col>
            </Row>
            <Row>
                <Col xs={9} className={styles.line} />
            </Row>
            <Row className={styles.lot}>
                <Col>
                    {item.tags.map((t, i) => <Badge key={i}>{t.name}</Badge>)}
                </Col>
            </Row>
        </Grid>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// Stats
// 
///////////////////////////////////////////////////////////////////////

interface StatsSummaryProps {
    summary: any
    isLoading: boolean
    onUpdate: () => void
}

export const StatsSummary: React.SFC<StatsSummaryProps> = (props) => {
    let summary = props.summary
    if (summary == undefined) {
        return (<div></div>)
    }
    let sums = summary.sums
    let corr = summary.corr

    let barData = [
        { key: "total", count: summary.total },
        ...Object.keys(sums).map(key => {
            return { key: key, count: sums[key] }
        })
    ]
    let lineData = Object.keys(corr).map(key => {
        return { ...corr[key], key: key }
    })

    let socialSums = summary.social_sums
    let socialCorr = summary.social_corr

    let socialBar = [
        ...Object.keys(socialSums).map(key => {
            return { key: key, count: socialSums[key] }
        })
    ]
    let socialLine = Object.keys(socialCorr).map(key => {
        return { ...socialCorr[key], key: key }
    })

    return (
        <div>
            <BarChart width={700} height={300} data={barData} margin={{ top: 10, right: 20, left: 20, bottom: 5 }}>
                <XAxis dataKey="key" /> <YAxis /> <Tooltip /> <Legend />
                <Bar dataKey="count" fill="#88D" />
            </BarChart>
            <LineChart width={700} height={500} data={lineData} margin={{ top: 40, right: 20, left: 20, bottom: 80 }}>
                <XAxis dataKey="key" /> <YAxis /> <Tooltip /> <Legend /> <CartesianGrid strokeDasharray="3 3" />
                <Line type="monotone" dataKey="name" stroke="#88D" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="description" stroke="#8D8" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="mail" stroke="#D88" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="link" stroke="#DD8" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="items" stroke="#D8D" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="contributions" stroke="#C76" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="followers" stroke="#67C" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="followees" stroke="#888" activeDot={{ r: 8 }} />
            </LineChart>

            <BarChart width={700} height={300} data={socialBar} margin={{ top: 10, right: 20, left: 20, bottom: 5 }}>
                <XAxis dataKey="key" /> <YAxis /> <Tooltip /> <Legend />
                <Bar dataKey="count" fill="#88D" />
            </BarChart>
            <LineChart width={700} height={500} data={socialLine} margin={{ top: 40, right: 20, left: 20, bottom: 80 }}>
                <XAxis dataKey="key" /> <YAxis /> <Tooltip /> <Legend /> <CartesianGrid strokeDasharray="3 3" />
                <Line type="monotone" dataKey="github" stroke="#88D" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="twitter" stroke="#8D8" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="facebook" stroke="#D88" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="linkedin" stroke="#DD8" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="google" stroke="#D8D" activeDot={{ r: 8 }} />
            </LineChart>
        </div>
    )
}

///////////////////////////////////////////////////////////////////////
// 
// Search
// 
///////////////////////////////////////////////////////////////////////

interface SearchFormProps {
    onClick: (q: any) => void
    isLoading: boolean
}

export class SearchForm extends React.Component<SearchFormProps> {
    text: string = ""
    useSimple: boolean = true

    onChangeText(t: string) {
        this.text = t
    }
    onChangeUseSimple(b: boolean) {
        this.useSimple = b
    }
    getQuery() {
        return {
            doc: this.text,
            simple: this.useSimple,
        }
    }

    render() {
        let props = this.props

        let placeholder = "Railsでつくっています。\nRubyを使える人を探しています。"
        return (
            <Form horizontal className={styles.searchForm}>
            <FormGroup>
                <Col xs={9}>
                    <FormControl componentClass="textarea" placeholder={placeholder}
                        multiple={true} rows={5}
                        onChange={e => this.onChangeText((e.target as any).value)} />
                </Col>
            </FormGroup>
            <FormGroup>
                <Col xs={6} className={styles.btnPanel}>
                    <Button disabled={props.isLoading}
                        onClick={_ => props.onClick(this.getQuery())}>検索する</Button>
                    <div className={styles.checkbox}>
                        <Checkbox defaultChecked={this.useSimple}
                            onClick={e => this.onChangeUseSimple(!this.useSimple)}>高速版</Checkbox>
                    </div>
                </Col>
            </FormGroup>
        </Form>
        )
    }
}
