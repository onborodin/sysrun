
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
import Vm from './Vm'

export const history = createBrowserHistory()

const cookieName = "session"

class Store {
    @observable username
    @observable limit

    constructor() {
        this.username = ""
        this.fileLimit = 10
        this.filePattern = "*"
        this.bucketLimit = 10
        this.bucketPattern = "*"
    }

    @action login = (username) => {
        console.log('login:', username)
        this.username = username
    }
    @action logout = () => {
        this.username = ""
        Cookies.remove(cookieName)
        history.push("/login")
    }
}

export const store = new Store()

export function checkLogin() {
    console.log("check login")
    let cookie = Cookies.get(cookieName)
    if (store.username == "" || cookie == null) {
        //history.push("/login")
    }
}

ReactDOM.render(
    <Router history={history}>
        <Switch>
            <Route exact path="/login" component={Login} />
            <Route exact path="/users" component={Users} />
            <Route exact path="/" component={Vm} />
            <Route path="*" component={NotFound} />
        </Switch>
    </Router>,
    document.getElementById('root')
)
