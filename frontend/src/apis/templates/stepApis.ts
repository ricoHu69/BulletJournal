import {doDelete, doFetch, doPost, doPut} from '../api-helper';

export const getSteps = () => {
    return doFetch('/api/steps')
        .then((res) => res.json())
        .catch((err) => {
            throw Error(err.message);
        });
}

export const getStep = (stepId: number) => {
    return doFetch(`/api/public/steps/${stepId}`)
        .then((res) => res.json())
        .catch((err) => {
            throw Error(err.message);
        });
}

export const createStep = (
    name: string, nextStepId?: number) => {
    const postBody = JSON.stringify({
        name: name,
        nextStepId: nextStepId
    });
    return doPost('/api/steps', postBody)
        .then(res => res.json())
        .catch(err => {
            throw Error(err.message);
        });
}

export const updateStep = (stepId: number, name?: string, nextStepId?: number) => {
    const putBody = JSON.stringify({
        name: name,
        nextStepId: nextStepId
    });
    return doPut(`/api/steps/${stepId}`, putBody)
        .then(res => res.json())
        .catch(err => {
            throw Error(err.message);
        });
}

export const deleteStep = (stepId: number) => {
    return doDelete(`/api/steps/${stepId}`)
        .then(res => res)
        .catch((err) => {
            throw Error(err.message);
        });
}