export function shuffleArray<T>(arr: T[]) {
    const arrCopy = [...arr]
    for (let i = 0; i < arrCopy.length; i++) {
        const j = Math.floor(Math.random() * (i + 1));
        [arrCopy[i], arrCopy[j]] = [arrCopy[j], arrCopy[i]];
    }
    return arrCopy;
}



