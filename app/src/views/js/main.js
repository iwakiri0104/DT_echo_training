const textarea = document.getElementById('textarea');
const title = document.getElementById('title');
const submit = document.getElementById('btn');

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

/*文字数制限アラート　内容*/
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

/*ラジオボタンアクション*/
const btn_form = document.getElementById("target")
const btn_id = document.getElementById('btn_id');
btn_id.addEventListener("click", () => {
    let btnParent = btn_id.parentNode
    if (btn_form.checked) {
        btnParent.style.backgroundColor = "#DBF1F4"
    }
})

/*バリデーション*/
window.addEventListener('DOMContentLoaded', () => {
    //エラーメッセージ
    const errMsgTitle = document.querySelector('.err-msg-title');
    const errMsgContent = document.querySelector('.err-msg-content');

    // 「投稿」ボタンの要素にクリックイベントを設定
    submit.addEventListener('click', (e) => {
        // タイトル入力欄の空欄チェック
        if (!title.value) {
            // デフォルトアクションをキャンセル
            e.preventDefault();
            errMsgTitle.classList.add('form-invalid');
            errMsgTitle.textContent = '・タイトルは必須です';
            title.classList.add('input-invalid');
        }
        // 内容入力欄の空欄チェック
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


/*マウスオーバーリアクション*/
function closePopUp() {
    fukidashi.style.display = "none";
}
function openPopUp() {
    closePopUp();
    fukidashi.style.display = "block";
}