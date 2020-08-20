dragons = {};

(async () => {
  const button = document.getElementById("run");
  button.disabled = true;

  // Initialize
  const go = new Go();
  let { module, instance } = await WebAssembly.instantiateStreaming(
    fetch("dragons.wasm"),
    go.importObject
  );
  button.disabled = false;

  // Run
  button.onclick = async () => {
    button.disabled = true;

    setTimeout(() => {
      // Call go functions...
      parsed = dragons.parse("_x_,___,___");
      console.log("Parsed:", parsed);
    }, 100);

    await go.run(instance);
    instance = await WebAssembly.instantiate(module, go.importObject);
  };
})();
