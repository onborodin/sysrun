import React, { Fragment, Component } from 'react'
import $ from 'jquery'
import { trim, ltrim } from 'validator'
import autobind from 'autobind-decorator'
import axios from 'axios'

export class VmStart extends Component {

    constructor(props) {
        super(props)
        this.state = {
            formIsValid: false,
            alertMessage: "",
            confirm: false
        }
    }

    @autobind
    showForm() {
        $("#start-vm-" + this.props.vmId).modal('show')
    }

    @autobind
    hideForm() {
        $("#start-vm-" + this.props.vmId).modal('hide')
    }

    @autobind
    startVm() {
        axios.post('/api/v1/vm/start', {
                id: this.props.vmId,
        }).then((res) => {
            if (res.data.error != null) {
                if (!res.data.error) {
                    this.hideForm()
                    this.props.startCallback()
                } else {
                    this.setState({
                        alertMessage: "Backend error: " + res.data.message
                    })
                }
            }
        }).catch((err) => {
            this.setState({
                alertMessage: "Communication error"
            })

        })
    }

    @autobind
    onSubmit(event) {
        event.preventDefault()
        if (!this.state.formIsValid) {
            return
        }
        this.startVm()
    }

    @autobind
    onConfirm(event) {
        this.setState({ confirm: event.target.checked }, () => { this.validateForm() })

    }

    validateForm() {
        if (this.state.confirm) {
            this.setState({ formIsValid: true })
        } else {
            this.setState({ formIsValid: false })
        }
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
    makeId(base) {
        return 'vm-start-' + base + '-' + this.props.vmId
    }

    render() {
        return (
            <Fragment>

                    <i className="fas fa-play" onClick={this.showForm}></i>


                <div className="modal fade" id={"start-vm-" + this.props.vmId} tabIndex="-1" role="dialog"  ref="form">
                    <div className="modal-dialog" role="document">
                        <div className="modal-content">

                            <form acceptCharset="UTF-8" onSubmit={this.onSubmit} ref="form">

                                <div className="modal-header">
                                    <h5 className="modal-title">Start vm [{this.props.vm.name}]</h5>
                                    <button type="button" className="close" onClick={this.hideForm}>
                                        <span>&times;</span>
                                    </button>
                                </div>

                                <div className="modal-body">

                                    <div className="form-group form-check">
                                        <input id={this.makeId("confirm")} className="form-check-input"
                                                    type="checkbox" onChange={this.onConfirm} defaultChecked={this.state.confirm} />
                                        <label className="form-check-label" htmlFor={this.makeId("confirm")}> I agree</label>
                                    </div>

                                </div>

                                {this.showAlert()}

                                <div className="modal-footer">
                                    <button type="button" className="btn btn-sm btn-secondary" onClick={this.hideForm}>Close</button>
                                    <button type="submit" className={this.state.formIsValid ? "btn btn-sm btn-primary" : "btn btn-sm btn-primary disabled"} onClick={this.onSubmit}>Start</button>
                                </div>

                            </form>

                        </div>
                    </div>
                </div>
            </Fragment>
        )
    }
}

export default VmStart
