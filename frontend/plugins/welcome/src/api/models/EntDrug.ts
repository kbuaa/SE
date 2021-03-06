/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Prescription
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDrugEdges,
    EntDrugEdgesFromJSON,
    EntDrugEdgesFromJSONTyped,
    EntDrugEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrug
 */
export interface EntDrug {
    /**
     * DrugName holds the value of the "Drug_Name" field.
     * @type {string}
     * @memberof EntDrug
     */
    drugName?: string;
    /**
     * 
     * @type {EntDrugEdges}
     * @memberof EntDrug
     */
    edges?: EntDrugEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDrug
     */
    id?: number;
}

export function EntDrugFromJSON(json: any): EntDrug {
    return EntDrugFromJSONTyped(json, false);
}

export function EntDrugFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrug {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugName': !exists(json, 'Drug_Name') ? undefined : json['Drug_Name'],
        'edges': !exists(json, 'edges') ? undefined : EntDrugEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDrugToJSON(value?: EntDrug | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Drug_Name': value.drugName,
        'edges': EntDrugEdgesToJSON(value.edges),
        'id': value.id,
    };
}


