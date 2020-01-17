import React, { Fragment, Component } from 'react'

import { Layout } from './Layout'
import { checkLogin } from './main'

export class Vm extends Component {

    render() {
        return (
            <Fragment>
                <Layout>
                    <div className="row">
                        <h4>VM</h4>
                    </div>
                </Layout>
            </Fragment>
        )
    }

    componentDidMount() {
        checkLogin()
    }

}

export default Vm

