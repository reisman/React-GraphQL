import React from 'react';
import { Card, CardBody, CardTitle } from 'reactstrap';

const User = ({name}) => {
    return (
        <div>
            <Card>
                <CardBody>
                    <CardTitle>{name}</CardTitle>
                </CardBody>
            </Card>
        </div>
    );
};

export default User;