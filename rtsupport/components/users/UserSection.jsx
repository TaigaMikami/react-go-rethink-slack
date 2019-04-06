import React, {Component} from 'react';
import PropTypes from 'prop-types';
import UserForm from './UserForm.jsx';
import UserList from './UserList.jsx';

class UserSection extends Component {
  render() {
    return(
      <div className='support card'>
        <div className="card-header bg-primary text-white">
          <strong>Users</strong>
        </div>
        <div className="card-body channels">
          <UserList {...this.props}/>
          <UserForm {...this.props}/>
        </div>
      </div>
    )
  }
}

UserSection.propTypes = {
  users: PropTypes.array.isRequired,
  setUser: PropTypes.func.isRequired,
}

export default UserSection
