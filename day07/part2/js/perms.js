function makePerms(low, high) {
  const digits = [];
  for (let i = low; i <= high; i++) {
    digits.push(i);
  }

  const allPerms = [];

  function inner(digitsArr, perm = []) {
    if (perm.length === 5) {
      allPerms.push(perm);
    } else {
      digitsArr.forEach((digit, index) => {
        inner(digitsArr.slice(0, index).concat(digitsArr.slice(index + 1)), perm.concat(digit))
      })
    }
  }

  inner(digits);
  return allPerms;
}

console.log(makePerms(0, 4));
