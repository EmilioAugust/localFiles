const fileInput = document.getElementById("fileInput")
const selectBtn = document.getElementById("selectBtn")
const dropZone = document.getElementById("dropZone")
const filesList = document.getElementById("filesList")

selectBtn.onclick = () => fileInput.click()

fileInput.onchange = () => {
    uploadFile(fileInput.files[0])
}

dropZone.addEventListener("dragover", e => {
    e.preventDefault()
})

dropZone.addEventListener("drop", e => {
    e.preventDefault()
    const file = e.dataTransfer.files[0]
    uploadFile(file)
})

async function uploadFile(file){

    const progressContainer = document.getElementById("progressContainer")
    const progressBar = document.getElementById("progressBar")

    progressContainer.style.display = "block"
    progressBar.style.width = "0%"

    const formData = new FormData()
    formData.append("file", file)

    const xhr = new XMLHttpRequest()

    xhr.open("POST", "/upload")

    xhr.upload.onprogress = function(e){
        if(e.lengthComputable){
            const percent = (e.loaded / e.total) * 100
            progressBar.style.width = percent + "%"
        }
    }

    xhr.onload = function(){
        progressBar.style.width = "100%"

        setTimeout(()=>{
            progressContainer.style.display = "none"
        },500)

        loadFiles()
    }

    xhr.send(formData)
}

function formatFileSize(bytes) {

    const mb = bytes / (1024 * 1024)
    const gb = bytes / (1024 * 1024 * 1024)

    if (gb >= 1) {
    return gb.toFixed(2) + " GB"
    }

    if (mb >= 1) {
    return mb.toFixed(2) + " MB"
    }

    const kb = bytes / 1024
    return kb.toFixed(2) + " KB"

}

async function loadFiles(){
    const res = await fetch("/files")

    const files = await res.json()

    filesList.innerHTML = ""

    files.forEach(file=>{

    const el = document.createElement("div")
    el.className="file"

    el.innerHTML = `
        <div class="file-info">
            <div class="file-name">${file.name}</div>
            <div class="file-size">${formatFileSize(file.size)}</div>
        </div>
        <div class="file-actions">
            <button onclick="deleteFile('${file.id}')">
                Delete
            </button>
            <button onclick="downloadFile('${file.id}')">
                Download
            </button>
        </div>
    `;

    filesList.appendChild(el)

    })

}

async function deleteFile(id) {

    const confirmDelete = confirm("Delete this file?")

    if (!confirmDelete) return

    await fetch("/delete/" + id, {
        method: "DELETE"
    })

    loadFiles()

}

function downloadFile(id){
    window.open("/download/" + id)
}

const uaDiv = document.getElementById("userAgent")
uaDiv.textContent = navigator.userAgent

loadFiles()

generateQR()