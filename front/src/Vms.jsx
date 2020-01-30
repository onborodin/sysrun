import React, { Component, Fragment } from 'react'
import autobind from 'autobind-decorator'
import axios from 'axios'

import { clone } from 'lodash'

import { Layout } from './Layout'
import { checkLogin } from './main'

import { VmStart } from './VmStart'
import { VmStop } from './VmStop'
import { VmShutdown } from './VmShutdown'

import { Pager } from './Pager'

export class Vms extends Component {

    constructor(props) {
        super(props)

        this.state = {
            vms: [],
            alertMessage: ""
        }
    }

    @autobind
    listVms() {
        axios.get('/api/v1/vm/list')
            .then((res) => {
                if (res.data.error != null) {
                    if (!res.data.error) {
                        this.setState({
                            vms: res.data.result,
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
    renderTable() {
        return this.state.vms.map((vm, index) => {
            const { id, name, status } = vm
            const theVm = vm
            return (
                <tr key={id}>
                    <td>{id}</td>
                    <td>{name}</td>
                    <td>{status}</td>
                    <td><VmStart vmId={id} vm={theVm} startCallback={this.listVms} /></td>
                    <td><VmShutdown vmId={id} vm={theVm} shutdownCallback={this.listVms} /></td>
                    <td><VmStop vmId={id} vm={theVm} stopCallback={this.listVms} /></td>
                </tr>
            )
        })
    }

    render() {
        return (
            <Fragment>
                <Layout>
                    <div className="container1">
                        <div className="row justify-content-between mb-3">
                            <h5>
                                <i className="fas fa-server"></i>
                                <span> VMs </span>
                                <i className="fas fa-sync fa-xs"  onClick={this.listVms}></i>
                            </h5>

                        </div>
                    </div>

                    <table className="table table-striped table-hover table-sm">

                        <thead className="thead-light">
                            <tr>
                                <th>#</th>
                                <th>name</th>
                                <th>status</th>

                                <th>
                                    <i className="fas fa-play"></i>
                                </th>
                                <th>
                                    <i className="fas fa-power-off"></i>
                                </th>

                                <th>
                                    <i className="fas fa-stop"></i>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {this.renderTable()}
                        </tbody>
                    </table>

                </Layout>
            </Fragment>
        )
    }

    componentDidMount() {
        checkLogin("user")
        this.listVms()
    }
}

export default Vms
