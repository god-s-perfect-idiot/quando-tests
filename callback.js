function processCallbacks(callback) {
    for (let i = 0; i < 100000; i++) {
      callback(i);
    }
  }
  
  function callbackFunction(index) {
    console.log("Callback index:", index);
  }
  
  processCallbacks(callbackFunction);