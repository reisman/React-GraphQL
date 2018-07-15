import moment from 'moment';
import createId from './identitygenerator';

const sample_authors = [
    {id: "1", name: "Hans"},
    {id: "2", name: "Peter"},
    {id: "3", name: "Sam"},
    {id: "4", name: "Tom"},
    {id: "5", name: "Joe"},
];

const sample_messages = [
    {id: "1", author: sample_authors[4], publishedat: moment().subtract(1, 'seconds').toDate(), text: "TextA"},
    {id: "2", author: sample_authors[2], publishedat: moment().subtract(2, 'seconds').toDate(), text: "TextB"},
    {id: "3", author: sample_authors[2], publishedat: moment().subtract(3, 'seconds').toDate(), text: "TextC"},
    {id: "4", author: sample_authors[3], publishedat: moment().subtract(4, 'seconds').toDate(), text: "TextE"},
    {id: "5", author: sample_authors[4], publishedat: moment().subtract(5, 'seconds').toDate(), text: "TextF"},
];

export const getAuthors = () => sample_authors;
export const findAuthor = (id) => sample_authors.find(a => a.id === id);
export const createAuthor = (name) => {
    const id = createId();
    const author = {id, name};
    sample_authors.push(author);
    return author;
};

export const getMessages = () => sample_messages;
export const findMessage = (id) => sample_messages.find(m => m.id === id);
export const createMessage = (text, author) => {
    const id = createId();
    const publishedat = moment().toDate();
    const msg = {id, author, publishedat, text};
    sample_messages.push(msg);
    return msg;
};

