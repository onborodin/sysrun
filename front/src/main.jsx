
import './app.scss'
import 'bootstrap'

import React from 'react'
import ReactDOM from 'react-dom'

import { Router, Route, Switch } from 'react-router-dom'
import { createBrowserHistory } from 'history'

import { Provider } from 'mobx-react';
import * as Cookies from 'js-cookie'

import { decorate, observable, action } from "mobx"

import NotFound from './NotFound'
import Login from './Login'
import Users from './Users'
import Vms from './Vms'

export const history = createBrowserHistory()

const cookieName = "session"

class Store {
    @observable username
    @observable limit

    constructor() {
        this.username = ""
        this.isadmin = false
    }

    @action login = (username, isadmin) => {
        this.username = username
        this.isadmin = isadmin
    }
    @action logout = () => {
        this.username = ""
        this.isadmin = false
        Cookies.remove(cookieName)
        history.push("/login")
    }
}

export const store = new Store()

export function checkLogin(accessLevel) {
    let cookie = Cookies.get(cookieName)
    if (store.username == "" || cookie == null) {
        history.push("/login")
    }
    if (accessLevel == "admin" && store.isadmin == false) {
        history.push("/login")
    }
}

ReactDOM.render(
    <Router history={history}>
        <Switch>
            <Route exact path="/login" component={Login} />
            <Route exact path="/users" component={Users} />
            <Route exact path="/" component={Vms} />
            <Route path="*" component={NotFound} />
        </Switch>
    </Router>,
    document.getElementById('root')
)
