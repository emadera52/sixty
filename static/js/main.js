(function($){

	// Avoid embed this site in an iframe of other WebSite
	if (top.location != location) {
		top.location.href = location.href;
	}

	// btn checked box toggle
	$(document).on('click', '.btn-checked', function(){
		var $e = $(this);
		var $i = $e.siblings('[name='+$e.data('name')+']');
		if($e.hasClass('active')) {
			$i.val('true');
		} else {
			$i.val('false');
		}
		$e.blur();
	});

	// change locale and reload page
	$(document).on('click', '.lang-changed', function(){
		var $e = $(this);
		var lang = $e.data('lang');
		$.cookie('lang', lang, {path: '/', expires: 365});
		window.location.reload();
	});

	(function(){
		var v = $.cookie('JsStorage');
		if(v){
			var values = v.split(':::');
			if(values.length > 1){
				$.jStorage[values[0]].apply(this, values.splice(1));
			}
			$.removeCookie('JsStorage', {path: '/'});
		}
	})();

	(function(){

		$.fn.mdFilter = function(){
			var $e = $(this);
			$e.find('img').each(function(_,img){
				var $img = $(img);
				$img.addClass('img-responsive');
			});

			var $pre = $e.find('pre > code').parent();
			$pre.addClass("prettyprint");
			prettyPrint();
		};

	})();


	$(function(){
        // Encode url.
        var $doc = $('.docs-markdown');
        $doc.find('a').each(function () {
            var node = $(this);
            var link = node.attr('href');
            var index = link.indexOf('#');
            if (link.indexOf('http') === 0 && link.indexOf(window.location.hostname) == -1) {
                return;
            }
            if (index < 0 || index + 1 > link.length) {
                return;
            }
            var val = link.substring(index + 1, link.length);
            val = encodeURIComponent(decodeURIComponent(val).toLowerCase().replace(/\s+/g, '-'));
            node.attr('href', link.substring(0, index) + '#' + val);
        });

        // Set anchor.
        $doc.find('h1, h2, h3, h4, h5, h6').each(function () {
            var node = $(this);
            if (node.hasClass('ui')) {
                return;
            }
            var val = encodeURIComponent(node.text().toLowerCase().replace(/\s+/g, "-"));
            node = node.wrap('<div id="' + val + '" class="anchor-wrap" ></div>');
            node.append('<a class="anchor" href="#' + val + '"><span class="octicon octicon-link"></span></a>');
        });
    });

	$(function(){
		// on dom ready

		$('[data-show=tooltip]').each(function(k, e){
			var $e = $(e);
			$e.tooltip({placement: $e.data('placement'), title: $e.data('tooltip-text')});
			$e.tooltip('show');
		});

		$('[rel=select2]').select2();

		$('.markdown').mdFilter();

        if($.jPanelMenu && $('[data-toggle=jpanel-menu]').size() > 0) {
        var jpanelMenuTrigger = $('[data-toggle=jpanel-menu]');

            var jPM = $.jPanelMenu({
                animated: false,
                menu: jpanelMenuTrigger.data('target'),
                direction: 'left',
                trigger: '.'+ jpanelMenuTrigger.attr('class'),
                excludedPanelContent: '.jpanel-menu-exclude',
                openPosition: '180px',
                afterOpen: function() {
                    jpanelMenuTrigger.addClass('open');
                    $('html').addClass('jpanel-menu-open');
                },
                afterClose: function() {
                    jpanelMenuTrigger.removeClass('open');
                    $('html').removeClass('jpanel-menu-open');
                }
            });

			//jRespond settings
			var jRes = jRespond([{
				label: 'small',
				enter: 0,
				exit: 1010
			}]);

			//turn jPanel Menu on/off as needed
			jRes.addFunc({
			breakpoint: 'small',
				enter: function() {
					jPM.on();
				},
				exit: function() {
					jPM.off();
				}
			});
		}

	});

})(jQuery);