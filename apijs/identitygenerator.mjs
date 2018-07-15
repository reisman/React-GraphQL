import uuid from 'uuid';

const createId = () => {
    return uuid.v4();
};

export default createId;