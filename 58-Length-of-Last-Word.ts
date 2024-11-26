function lengthOfLastWord(s: string): number {
    const splitted = s.split(" ");
    const newSplitted = splitted.filter((element) => element.length > 0);

    return newSplitted[newSplitted.length - 1].length;
}
