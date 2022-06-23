const textarea = document.getElementById('textarea');
const title = document.getElementById('title');
const submit = document.getElementById('btn');
const errMsgTitle = document.querySelector('.err-msg-title');
const errMsgContent = document.querySelector('.err-msg-content');


/*文字数制限アラート　タイトル*/
title.addEventListener('keyup', (e) => {
    if (e.target.value.length > 10) {
        alert("10文字以内でお願いします");
        submit.disabled = true;
    }
    ;
    if (e.target.value.length <= 10) {
        submit.disabled = false;
    }
    ;
});
/*文字数制限アラート　文章*/
textarea.addEventListener('keyup', (e) => {
    if (e.target.value.length > 300) {
        alert("300文字以内でお願いします");
        submit.disabled = true;
    }
    ;
    if (e.target.value.length <= 300) {
        submit.disabled = false;
    }
    ;
});

/*文字数カウント*/
function viewStrLen() {
    const len = document.getElementById("textarea").value.length;
    document.getElementById("strLen").innerText = len + "文字";
}

/*ラジオボタン*/
const btn_id = document.getElementById("btn_id")
btn_id.addEventListener("click", () => {
    let btnParent = btn_id.parentNode
    if (btn_id.checked) {
        btnParent.style.backgroundColor = "#DBF1F4"
    }
})


//イベントリスナーでid属性に対して処理してる
//idは1つしかもてない
//idを全部違うやつをつける
//クラスの名前を同じにすれば

window.addEventListener('DOMContentLoaded', () => {
    // 「送信」ボタンの要素にクリックイベントを設定する
    submit.addEventListener('click', (e) => {
        // 入力欄の空欄チェック
        if (!title.value) {
            // デフォルトアクションをキャンセル
            e.preventDefault();
            errMsgTitle.classList.add('form-invalid');
            errMsgTitle.textContent = '・タイトルは必須です';
            title.classList.add('input-invalid');
        }
        if (!textarea.value) {
            e.preventDefault();
            errMsgContent.classList.add('form-invalid');
            errMsgContent.textContent = '・内容は必須です';
            textarea.classList.add('input-invalid');
        }

        return;
        // エラーメッセージのテキストに空文字を代入
        errMsgContent.textContent = '';
        errMsgTitle.textContent = '';
        // クラスを削除
        textarea.classList.remove('input-invalid');
        title.classList.remove('input-invalid');

    }, false);
}, false);


function closePopUp() {
    fukidashi.style.display = "none";
}
function openPopUp() {
    closePopUp();
    fukidashi.style.display = "block";
}