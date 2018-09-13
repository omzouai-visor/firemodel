// DO NOT EDIT - Code generated by firemodel.

import firebase from 'firebase';

export interface DocumentSnapshot<DataType = firebase.firestore.DocumentData>
  extends firebase.firestore.DocumentSnapshot {
  data(options?: firebase.firestore.SnapshotOptions): DataType | undefined;
}
export interface QueryDocumentSnapshot<
  DataType = firebase.firestore.DocumentData
> extends firebase.firestore.QueryDocumentSnapshot {
  data(options?: firebase.firestore.SnapshotOptions): DataType | undefined;
}
export interface QuerySnapshot<DataType = firebase.firestore.DocumentData>
  extends firebase.firestore.QuerySnapshot {
  readonly docs: QueryDocumentSnapshot<DataType>[];
}
export interface DocumentSnapshotExpanded<
  DataType = firebase.firestore.DocumentData
> {
  exists: firebase.firestore.DocumentSnapshot['exists'];
  ref: firebase.firestore.DocumentSnapshot['ref'];
  id: firebase.firestore.DocumentSnapshot['id'];
  metadata: firebase.firestore.DocumentSnapshot['metadata'];
  data: DataType;
}
export interface QuerySnapshotExpanded<
  DataType = firebase.firestore.DocumentData
> {
  metadata: {
    hasPendingWrites: firebase.firestore.QuerySnapshot['metadata']['hasPendingWrites'];
    fromCache: firebase.firestore.QuerySnapshot['metadata']['fromCache'];
  };
  size: firebase.firestore.QuerySnapshot['size'];
  empty: firebase.firestore.QuerySnapshot['empty'];
  docs: {
    [docId: string]: DocumentSnapshotExpanded<DataType>;
  };
}
export interface DocumentReference<DataType> extends firebase.firestore.DocumentReference {
  data(options?: firebase.firestore.SnapshotOptions): DataType | undefined;
}
export interface CollectionReference<
  DataType = firebase.firestore.DocumentData
> extends firebase.firestore.CollectionReference {
  get(options?: firebase.firestore.GetOptions): Promise<QuerySnapshot<DataType>>;
}
export interface Collection<DataType = firebase.firestore.DocumentData> {
  [id: string]: DocumentSnapshotExpanded<DataType>;
}

export namespace example {
  type URL = string;

  export interface IFile {
    url: URL;
    mimeType: string;
    name: string;
  }

  /** TODO: Add documentation to Direction. */
  export enum Direction {
    /** TODO: Add documentation to left. */
    left = "LEFT",
    /** TODO: Add documentation to right. */
    right = "RIGHT",
    /** TODO: Add documentation to up. */
    up = "UP",
    /** TODO: Add documentation to down. */
    down = "DOWN",
  }

  /** A Test is a test model. */
  export interface ITestModel {
    /** TODO: Add documentation to nested. */
    nested: CollectionReference<ITestModel>;
    /** The name. */
    name?: string;
    /** The age. */
    age?: number;
    /** The number pi. */
    pi?: number;
    /** The birth date. */
    birthdate?: firebase.firestore.Timestamp;
    /** True if it is good. */
    isGood?: boolean;
    /** TODO: Add documentation to data. */
    data?: firebase.firestore.Blob;
    /** TODO: Add documentation to friend. */
    friend?: DocumentReference<ITestModel>;
    /** TODO: Add documentation to location. */
    location?: firebase.firestore.GeoPoint;
    /** TODO: Add documentation to colors. */
    colors?: string[];
    /** TODO: Add documentation to meta. */
    meta?: { [key: string]: any; };
    /** TODO: Add documentation to direction. */
    direction?: Direction;
    /** TODO: Add documentation to test_file. */
    testFile?: IFile;
    /** TODO: Add documentation to url. */
    url?: URL;
  }
}