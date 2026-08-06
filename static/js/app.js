const button = document.getElementById("hello-btn");
const panel = document.getElementById("response");

button.addEventListener("click", async () => {
  button.disabled = true;
  panel.classList.remove("error");
  panel.textContent = "Calling /api/hello…";

  try {
    const res = await fetch("/api/hello");
    if (!res.ok) {
      throw new Error(`HTTP ${res.status}`);
    }
    const data = await res.json();
    panel.textContent = JSON.stringify(data, null, 2);
  } catch (err) {
    panel.classList.add("error");
    panel.textContent = `Request failed: ${err.message}`;
  } finally {
    button.disabled = false;
  }
});
