/**
 * blog列表
 */
let form, layer;
layui.use(['form', 'layer'], function () {
    form = layui.form;
    layer = layui.layer;


});

function star(id) {
    layer.msg('no no no,' + id, {icon: 5})
}

function comment(id) {

}