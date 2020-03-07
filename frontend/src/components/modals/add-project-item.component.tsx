import React from 'react';
import { Modal, Input, Form, Button, Select, Avatar, Tooltip } from 'antd';
import {
  PlusOutlined
} from '@ant-design/icons';
import { connect } from 'react-redux';
import { GroupsWithOwner } from '../../features/group/interfaces';
import { createProjectByName } from '../../features/project/actions';
import { updateGroups } from '../../features/group/actions';
import { ProjectType, toProjectType } from '../../features/project/constants';
import { IState } from '../../store';
import { Project } from '../../features/project/interfaces';
import { History } from 'history';

import './modals.styles.less';

//props of groups
type GroupProps = {
  groups: GroupsWithOwner[];
  updateGroups: () => void;
};

type ModalState = {
  isShow: boolean;
};

class AddProjectItem extends React.Component<GroupProps,
  ModalState
> {
  componentDidMount() {
    this.props.updateGroups();
  }

  state: ModalState = {
    isShow: false,
  };

  showModal = () => {
    this.setState({ isShow: true });
  };

  onCancel = () => {
    this.setState({ isShow: false });
  };

  render() {
    const { groups: groupsByOwner } = this.props;

    const modal = (
    <Modal
      title='Create New BuJo Item'
      visible={this.state.isShow}
      onCancel={this.onCancel}
      footer={[
        <Button key='cancel' onClick={this.onCancel}>
          Cancel
        </Button>,
        <Button
          key='create'
          type='primary'
        >
          Create
        </Button>
      ]}
    >
    </Modal>);
      return (
        <div>
          <Tooltip placement="bottom" title='Create New BuJo Item' >
            <PlusOutlined className='rotateIcon' onClick={this.showModal}/>
          </Tooltip>
          {modal}
        </div>
      );
  }
}

const mapStateToProps = (state: IState) => ({
  groups: state.group.groups,
  project: state.project.project,
});

export default connect(mapStateToProps, { updateGroups, createProjectByName })(
  AddProjectItem
);
