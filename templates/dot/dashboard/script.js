
$(function () {
	var $tbl = $('#maintable>tbody');

	// Filter - sent each time when list of users is loaded
	var filter = { 
	}

	// load list of users
	function update() {
		$tbl.listLoad({
			url: ajaxUrl + 'getAll',
			method: 'POST',
			noitemsTpl: 'noitems',
			itemTpl: 'item',
			data: filter,
		});
	}
	update();

	// "New item" button
	var $newAdmin = $.tools.addButton({
		html: '<span class="fa fa-plus"></span> Add record',
		onClick: function () {
			$.wbox.open('Add new record', window.tplRet('form-add'));
		}
	});

	// Submit form "New item"
	$('#add-form').ajaxFormSender({
		url: ajaxUrl + 'add',
		success: function (result) {
			$tbl.tplPrepend('item', result.data);
			$('#noitems').hide().remove();
			$tbl.children('tr[data-id=' + result.data.id + ']').children().highlight(500);
			return true;
		}
	});

	// "Edit" button pressed
	$('.edit').ajaxActionSender({
		url: ajaxUrl + 'get',
		method: 'POST',
		success: function (result) {
			$.wbox.open('Change record', window.tplRet('form-edit', result.data));
		}
	});

	// Submit form "Edit"
	$('#edit-form').ajaxFormSender({
		url: ajaxUrl + 'edit',
		success: function (result) {
			result.data.teasersCount = '...';
			$tbl.children('tr[data-id=' + result.data.id + ']').tplReplace('item', result.data);
			$tbl.children('tr[data-id=' + result.data.id + ']').children().highlight(500);
			return true;
		}
	});

	// "Remove" button pressed
	$('.remove').ajaxActionSender({
		url: ajaxUrl + 'delete',
		method: 'POST',
		remove: true // remove from list if success
	});
});
