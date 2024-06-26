/* tslint:disable */
/* eslint-disable */
/**
 * zrok
 * zrok client access
 *
 * The version of the OpenAPI document: 0.3.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface CreateIdentityRequest
 */
export interface CreateIdentityRequest {
    /**
     * 
     * @type {string}
     * @memberof CreateIdentityRequest
     */
    name?: string;
}

/**
 * Check if a given object implements the CreateIdentityRequest interface.
 */
export function instanceOfCreateIdentityRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateIdentityRequestFromJSON(json: any): CreateIdentityRequest {
    return CreateIdentityRequestFromJSONTyped(json, false);
}

export function CreateIdentityRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateIdentityRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function CreateIdentityRequestToJSON(value?: CreateIdentityRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
    };
}

