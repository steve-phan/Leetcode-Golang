const convertToInt = (str: string) => {
  let int = 0;
  const baseCharCode = 'a'.charCodeAt(0);

  for (let i = 0; i < str.length; i++) {
    int |= 10 << (str.charCodeAt(i) - baseCharCode);
  }

  return int;
};
