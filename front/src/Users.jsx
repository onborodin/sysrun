import React, { Component, Fragment } from 'react'
import autobind from 'autobind-decorator'
import axios from 'axios'

import { clone } from 'lodash'

import { Layout } from './Layout'
import { checkLogin } from './main'

import { UserCreate } from './UserCreate'
import { UserUpdate } from './UserUpdate'
import { UserDelete } from './UserDelete'
import { Pager } from './Pager'

export class Users extends Component {

    constructor(props) {
        super(props)

        this.state = {
            users: [],
            offset: 0,
            limit: 5,
            total: 0,
            alertMessage: ""
        }
    }

    @autobind
    listUsers() {
        axios.post('/api/v2/user/list', {
                limit: this.state.limit,
                offset: this.state.offset
        }).then((res) => {
            if (res.data.error != null) {
                let users = []
                if (res.data.result.users != null) {
                    users = res.data.result.users
                }

                if (!res.data.error) {
                    this.setState({
                        users: users,
                        total: res.data.result.total,
                        offset: res.data.result.offset,
                        limit: res.data.result.limit
                    })
                } else {
                    this.setState({
                        alertMessage: "Backend error"
                    })
                }
            }
        }).catch((err) => {
            this.setState({
                alertMessage: "Communication error"
            })
        })
    }

    showAlert() {
        if (this.state.alertMessage != "") {
            return (
                <div className="alert alert-warning border mx-4" role="alert">
                  <div className="text-center">{this.state.alertMessage}</div>
                </div>
            )
        }
    }

    @autobind
    changeOffset(newOffset) {
        this.setState({ offset: newOffset }, () => { this.listUsers() })
    }

    @autobind
    onChangeLimit(event) {
        const newLimit = Number(event.target.value)
        var newOffset = Math.floor(this.state.offset / newLimit) * newLimit
        this.setState({ limit: newLimit, offset: newOffset }, () => { this.listUsers() })
    }


    @autobind
    renderTable() {
        return this.state.users.map((user, index) => {
            const { id, username, isadmin } = user
            const theUser = user
            return (
                <tr key={id}>
                    <td>{index + this.state.offset + 1}</td>
                    <td>{username}</td>
                    <td>{this.renderIsAdmin(isadmin)}</td>
                    <td><UserUpdate user={theUser} updateCallback={this.listUsers} /></td>
                    <td><UserDelete userId={id} user={theUser} updateCallback={this.listUsers} /></td>
                </tr>
            )
        })
    }

    @autobind
    renderIsAdmin(isadmin) {
        if (isadmin) {
            return <i className="fas fa-user-plus"></i>;
        } else {
            return <i className="fas fa-microchip"></i>;
        }
    }

    render() {
        return (
            <Fragment>
                <Layout>
                    <div className="container1">
                        <div className="row justify-content-between mb-3">

                            <div className="col-2 d-inline-flex">
                                    <h5>Users</h5>

                                    <div className="btn btn-primary btn-sm mx-3" onClick={this.listUsers}>
                                        <i className="fas fa-sync fa-xs"></i>
                                    </div>
                            </div>

                            <div className="col-1">
                                <UserCreate updateCallback={this.listUsers}/>
                            </div>

                        </div>
                    </div>

                    <div className="d-inline-flex mb-1">
                        <select className="custom-select" id="page-limit" value={this.state.limit} onChange={this.onChangeLimit}>
                            <option value="5">5</option>
                            <option value="10">10</option>
                            <option value="25">25</option>
                            <option value="50">50</option>
                        </select>
                    </div>

                    <table className="table table-striped table-hover table-sm">

                        <thead className="thead-light">
                            <tr>
                                <th>#</th>
                                <th>login name</th>
                                <th>mode</th>

                                <th>
                                    <i className="far fa-edit"></i>
                                </th>
                                <th>
                                    <i className="far fa-trash-alt"></i>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {this.renderTable()}
                        </tbody>
                    </table>

                    <Pager total={this.state.total} limit={this.state.limit} offset={this.state.offset} callback={this.changeOffset} />

                </Layout>
            </Fragment>
        )
    }

    componentDidMount() {
        checkLogin("admin")
        this.listUsers()
    }
}

export default Users
