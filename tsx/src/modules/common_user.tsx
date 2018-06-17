import { unstable_renderSubtreeIntoContainer } from "react-dom";

export interface UserSummary {
    user: User
    links: {[key: number] :UserSocialLink}
    stat: UserStat
    langStats: {[key: number] :UserLanguageStat}
    scout: UserScout
}

export const initUserSummary: UserSummary = {
    user: {
        id: 0,
        qiitaId: "",
        name: "",
        imageUrl: "",
        description: "",
        mail: "",
        link: "",
        organization: "",
        place: "",
        qiitaOrganization: "",
        ban: false,
    },
    links: [],
    stat: {
        id: 0,
        userId: 0,
        items: 0,
        contributions: 0,
        followers: 0,
        followees: 0,
    },
    langStats: [],
    scout: {
        id: 0,
        userId: 0,
        starred: false,
    }
}

export interface User {
    id: number
    qiitaId: string
    name: string
    imageUrl: string
    description: string
    mail: string
    link: string
    organization: string
    place: string
    qiitaOrganization: string
    ban: boolean
}

export interface UserSocialLink {
    id: number
    userId: number
    serviceId: number
    url: string
}

export interface UserStat {
    id: number
    userId: number
    items: number
    contributions: number
    followers: number
    followees: number
}

export interface UserLanguageStat {
    id: number
    userId: number
    name: string
    quantity: number
}

export interface UserScout {
    id: number
    userId: number
    starred: boolean
}

export interface UserItemSummary {
	items: UserItemWithTags[]
	populars: UserPopularItem[]
	recents: UserRecentItem[]
}

export const initUserItemSummary: UserItemSummary = {
    items: [],
    populars: [],
    recents: []
}

export interface UserItemWithTags {
	body: UserItem
	tags: UserItemTag[]
}

export interface UserItem {
	id: number
	userId: number
	articleId: string
	contributions: number
	comments: number
	title: number
	date: string
}

export interface UserItemTag {
	id: number
	itemId: number
	name: string
}

export interface UserPopularItem {
	id: number
	userId: number
	itemId: number
}

export interface UserRecentItem {
	id: number
	userId: number
	itemId: number
}
