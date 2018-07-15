import React from 'react';
import { Card, CardText, CardBody, CardTitle } from 'reactstrap';
import moment from 'moment';

const Message = ({text, publishedat, author }) => {
    return (
        <div>
            <Card>
                <CardBody>
                    <CardTitle>Message from {author} ({moment(publishedat).format('LTS')})</CardTitle>
                    <CardText>{text}</CardText>
                </CardBody>
            </Card>
        </div>
    );
};

export default Message;