function fibonacci(n) {
    if (n <= 0) {
      return [];
    } else if (n === 1) {
      return [0];
    } else if (n === 2) {
      return [0, 1];
    }
    
    let fibSequence = [0, 1];
    for (let i = 2; i < n; i++) {
      fibSequence.push(fibSequence[i - 1] + fibSequence[i - 2]);
    }
    
    return fibSequence;
  }
  
  const n = 1000000;
  const fibSequence = fibonacci(n);
  console.log('generated fibonacci sequence of length %d', fibSequence.length);

