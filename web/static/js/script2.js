// document.write("<h2>Это document.write из файла</h2>");
// alert('Это alert из файла');
// function message() {
//   alert('Это alert из функции message из файла');
// }

    var myVar,  fio_inp, fio_div, fio_btn, tn_inp, tn_div, tn_btn;

    function fioFunc() {
      myVar = setTimeout(clearFIO, 1000);
    }
    function tnFunc() {
      myVar = setTimeout(clearTN, 1000);
    }
    function clearFIO() {
      document.getElementById('in_fio').value = '';
      // fio_inp = document.getElementById('in_fio')
      // fio_inp.value = '';
    }
    function clearTN() {
      document.getElementById('in_tn').value = '';
      // tn_inp = document.getElementById('in_tn')
      // tn_inp.value = '';
    }
// Цвет текста по клику в другом поле
function changeTN() {
  // fio_inp = document.getElementById('in_fio')
  // fio_div = document.getElementById('div_fio')
  // fio_btn = document.getElementById('btn_fio')
  // tn_inp = document.getElementById('in_tn')
  // tn_div = document.getElementById('div_tn')
  // tn_btn = document.getElementById('btn_tn')
  document.getElementById('in_fio').style.color = '#E1E1E1';
  document.getElementById('in_fio').style['border-color'] = '#E1E1E1';
  document.getElementById('div_fio').style.color = '#E1E1E1';
  document.getElementById('btn_fio').style['background-image'] = 'url(../img/find_light.gif)';
  document.getElementById('btn_fio').style.pointerEvents = "none";
  document.getElementById('in_tn').style.color = 'hsl(0, 13%, 55%)';
  document.getElementById('in_tn').style['border-color'] = '#C4C4C4';
  document.getElementById('div_tn').style.color = '#1F1F1F';
  document.getElementById('btn_tn').style['background-image'] = 'url(../img/find.gif)';
  document.getElementById('btn_tn').style.pointerEvents = "auto";
}

function changeFIO() {
  document.getElementById('in_tn').style.color = '#E1E1E1';
  document.getElementById('in_tn').style['border-color'] = '#E1E1E1';
  document.getElementById('div_tn').style.color = '#E1E1E1';
  document.getElementById('btn_tn').style['background-image'] = 'url(../img/find_light.gif)';
  document.getElementById('btn_tn').style.pointerEvents = "none";
  document.getElementById('in_fio').style.color = 'hsl(0, 13%, 55%)';
  document.getElementById('in_fio').style['border-color'] = '#C4C4C4';
  document.getElementById('div_fio').style.color = '#1F1F1F';
  document.getElementById('btn_fio').style['background-image'] = 'url(../img/find.gif)';
  document.getElementById('btn_fio').style.pointerEvents = "auto";
}
function changeFIO_btn() {
  document.getElementById('btn_fio').style['background-image'] = 'url(../img/find-focus.gif)';
}
function rechangeFIO_btn() {
  document.getElementById('btn_fio').style['background-image'] = 'url(../img/find.gif)';
}
function changeTN_btn() {
  document.getElementById('btn_tn').style['background-image'] = 'url(../img/find-focus.gif)';
}
function rechangeTN_btn() {
  document.getElementById('btn_tn').style['background-image'] = 'url(../img/find.gif)';
}
function changeOther() {
    if (fio_inp.value = '') {
      fio_inp.style.color = '#E1E1E1';
      fio_inp.style['border-color'] = '#E1E1E1';
      fio_div.style.color = '#E1E1E1';
      fio_btn.style['background-image'] = 'url(../img/find_light.gif)';
      fio_btn.style.pointerEvents = "none";
    } else if (tn_inp.value = '') {
        tn_inp.style.color = '#E1E1E1';
        tn_inp.style['border-color'] = '#E1E1E1';
        tn_div.style.color = '#E1E1E1';
        tn_btn.style['background-image'] = 'url(../img/find_light.gif)';
        tn_btn.style.pointerEvents = "none";
    } else {
      fio_inp.style.color = 'hsl(0, 13%, 55%)';
      fio_inp.style['border-color'] = '#C4C4C4';
      fio_div.style.color = '#1F1F1F';
      fio_btn.style['background-image'] = 'url(../img/find.gif)';
      fio_btn.style.pointerEvents = "auto";
      tn_inp.style.color = 'hsl(0, 13%, 55%)';
      tn_inp.style['border-color'] = '#C4C4C4';
      tn_div.style.color = '#1F1F1F';
      tn_btn.style['background-image'] = 'url(../img/find.gif)';
      tn_btn.style.pointerEvents = "auto";
    }

}


// первый вариант - очистка обоих полей

    function myFunction() {
      myVar = setTimeout(clearInput, 1000);
    }
    function clearInput() {
      // alert('Это alert из функции clearInput в head');
      document.getElementById('in_fio').value = '';
      document.getElementById('in_tn').value = '';
      // var frm = document.getElementsByName('name')[0];
      // frm.submit(); // Submit
      // frm.reset(); // Reset
      // frm.value = '';
      // return false; // Prevent page refresh
    }

