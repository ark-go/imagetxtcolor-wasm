self.importScripts("wasm_exec.js"); //
//!----------- Первый способ загрузки Go Wasm ----------------------
// const go = new Go();
// WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
//   (result) => {
//     go.run(result.instance);
//     console.log("run");
//   }
// );
//!------------ ^^^^^^^^^^^  -----------------------
//!----------- Второй способ загрузки Go Wasm -----------------------
async function run(fileUrl) {
  try {
    const file = await fetch(fileUrl);
    const buffer = await file.arrayBuffer();
    const go = new Go();
    const { instance } = await WebAssembly.instantiate(buffer, go.importObject);
    go.run(instance);
  } catch (err) {
    console.error(err);
  }
}
setTimeout(() => run("main.wasm"));
//!----------- ^^^^^^^^^^^^^ -----------------------
// Сообения из Chrome клиента
self.onmessage = function (evt) {
  var objData = evt.data; // payload
  var sMessagePurpose = objData.MessagePurpose;
  if (sMessagePurpose === "goCreateImages") {
    let ar = createImages(objData.txt);
    self.postMessage(ar);
  }
  if (sMessagePurpose === "getImageId") {
    let ar = getImageFromId(objData.id);
    self.postMessage(ar);
  }
};
