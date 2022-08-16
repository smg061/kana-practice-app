export interface IPriorityQueue<T> {
    swap(idx1: number, idx2: number): T[]
    bubbleUp(): void;
    enqueue(value: T): T[]
    bubbleDown(): void;
    dequeue(): T | undefined
}

export interface Node<T> {
    value: T,
    priority: number,
}
class PQNode<T> implements Node<T> {
    constructor(public priority: number, public value: T) { }
}

export class PriorityQueue<T> implements IPriorityQueue<Node<T>> {
    private _data: Node<T>[];
    constructor(initialData?: T[]) {
        this._data = [];
        if (initialData) {
            const nodes = initialData.map((datum, i) => new PQNode(i + 1, datum));
            this._data = nodes;
        }
    }
    swap(idx1: number, idx2: number): Node<T>[] {
        let temp = this._data[idx1];
        this._data[idx1] = this._data[idx2];
        this._data[idx1] = temp;
        return this._data
    }
    bubbleUp() {
        // get index of inserted element
        let idx = this._data.length - 1;
        while (idx > 0) {
            let parentIdx = Math.floor((idx - 1) / 2);
            // swap if priority is greater than parent
            if (this._data[parentIdx].priority > this._data[idx].priority) {
                this.swap(idx, parentIdx);
                idx = parentIdx;
            } else {
                break;
            }
        }
        return 0;
    }
    enqueue(value: Node<T>): Node<T>[] {
        this._data.push(value);
        this.bubbleUp();
        return this._data
    }
    bubbleDown() {
        let parentIdx = 0;
        const length = this._data.length;
        let elementPriority = this._data[0].priority;
        while (true) {
            let leftChildIdx = (2 * parentIdx) + 1;
            let rightChildIdx = (2 * parentIdx) + 2;
            let leftChildPriority: number = -1;
            let rightChildPriority: number = -1;
            let idxToSwap: number | null = null;
            if (leftChildIdx < length) {
                leftChildPriority = this._data[leftChildIdx].priority;
                if (leftChildPriority < elementPriority) {
                    idxToSwap = leftChildIdx;
                }
            }
            if (rightChildIdx < length) {
                rightChildPriority = this._data[rightChildIdx].priority;

                if (
                    (rightChildPriority < elementPriority && idxToSwap === null) ||
                    (rightChildPriority < leftChildPriority && idxToSwap !== null)
                ) {
                    idxToSwap = rightChildIdx;
                }

            }

            if (idxToSwap === null) {
                break;
            }
            // swap with planned element
            this.swap(parentIdx, idxToSwap);
            parentIdx = idxToSwap;
        }
    }
    dequeue(): Node<T> | undefined {
        this.swap(0, this._data.length - 1);
        let poppedNode = this._data.pop();
        // re-adjust the heap if length is greater than 1
        if (this._data.length) {
            this.bubbleDown();
        }
        return poppedNode;
    }
}