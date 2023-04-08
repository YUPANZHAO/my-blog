import request from "../request.js";

export const getArticleList = (params) => {
    return request('get', `/article/list`, params)
}