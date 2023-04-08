import request from "../request.js";

export const getArticle = (params) => {
    return request('get', `/article/paper`, params)
}