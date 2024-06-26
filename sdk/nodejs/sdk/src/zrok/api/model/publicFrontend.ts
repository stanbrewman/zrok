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

import { RequestFile } from './models';

export class PublicFrontend {
    'token'?: string;
    'zId'?: string;
    'urlTemplate'?: string;
    'publicName'?: string;
    'createdAt'?: number;
    'updatedAt'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "token",
            "baseName": "token",
            "type": "string"
        },
        {
            "name": "zId",
            "baseName": "zId",
            "type": "string"
        },
        {
            "name": "urlTemplate",
            "baseName": "urlTemplate",
            "type": "string"
        },
        {
            "name": "publicName",
            "baseName": "publicName",
            "type": "string"
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "number"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return PublicFrontend.attributeTypeMap;
    }
}

