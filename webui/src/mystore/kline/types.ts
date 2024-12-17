export enum API {
  GetKPointsForLine = '/v1/get/kpoints/for/line',
}

export interface KPoint {
  Nums: number[];
  Times: number[];
  FormatTimes: string[];
}

export interface EchartKPoints {
  CategoryItems: string[];
  Nums: number[][];
}

export interface GetKPointsForLineResponse {
  KPointType: string;
  KPoints: KPoint[];
  Limit: number;
  Offset: number;
  OriginalTime: number;
  TokenPairID: number;
  Total: number;
}
