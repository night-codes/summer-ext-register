$(function () {
	$.tools.forceClick('.sign-in', function () { // вход
		$.wbox.open('Вход в кабинет вебмастера', window.tplRet('login'));
	});
	$.tools.forceClick('.sign-up', function () { // рега
		$.wbox.open('Регистрация вебмастера', window.tplRet('register'));
	});

	$('.sign-in-form').ajaxFormSender({
		success: function (result) {
			window.location.replace(panelPath + "home");
			return true;
		}
	});

	$('.sign-up-form').ajaxFormSender({
		success: function (result) {
			alert("Вы успешно зарегистрированы! С вами свяжется служба поддержки.")
			window.location.replace(panelPath + "home");
			return true;
		}
	});
});
