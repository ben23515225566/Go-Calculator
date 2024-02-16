const oneBtn = document.querySelector("#number_1")
const twoBtn = document.querySelector("#number_2")
const threeBtn = document.querySelector("#number_3")
const fourBtn = document.querySelector("#number_4")
const fiveBtn = document.querySelector("#number_5")
const sixBtn = document.querySelector("#number_6")
const sevenBtn = document.querySelector("#number_7")
const eightBtn = document.querySelector("#number_8")
const nineBtn = document.querySelector("#number_9")
const zeroBtn = document.querySelector("#number_0")

const resultElement = document.querySelector("#result")
const clearBtn = document.querySelector("#clear")

let displayVal = "0"

let numBtns = document.querySelectorAll(".numberBtn")
console.log("loading...")

function updateResult(clickObj){
    // clickObj表示點擊的事件，而clickObj.target表示點擊事件的DOM對象，在這裡也就是<button>
    var btnText = clickObj.target.innerHTML
    if(displayVal === "0"){
        displayVal = ""
    }
    displayVal += btnText
    resultElement.value = displayVal
}
function clear(clickObj) {
    var btnText = clickObj.target.innerHTML
    displayVal = "0"
    resultElement.value = displayVal
}
for(let i = 0; i < numBtns.length; i++){
    numBtns[i].addEventListener("click", updateResult, false)
}
clearBtn.addEventListener("click", clear, false)