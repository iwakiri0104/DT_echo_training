function check()
{
    //②次にその関数内でテキストエリアの値を取得。
    txt = document.form1.text1.value;
    //③その後、その文字列の文字数を取得。
    n = txt.length;
    //④最後にその文字数が制限文字数以上だった場合、アラートを実行。
    if (n > 10) alert("10文字以内で入力してください");
}
