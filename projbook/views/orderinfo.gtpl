<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>订单系统</title>
</head>

<body>
<form id="form1" name="form1" method="post" action="/list">
订单号：
<label for="orderId"></label>
<input type="text" name="orderId" id="orderId" />
订单类型：
<label for="ordertype"></label>
<select name="ordertype" id="ordertype">
  <option value="order_jd">京东订单</option>
  <option value="order_sunning">苏宁订单</option>
  <option value="order_mao">猫超订单</option>
</select>
提交人：
<label for="order_person"></label>
<input type="text" name="order_person" id="order_person" />
<input type="submit" name="submit" id="submit" value="提交" />
</form>
<p>&nbsp;</p>
<table width="780" height="30" border="1">
  <tr>
    <th scope="col">订单号</th>
    <th scope="col">订单类型</th>
    <th scope="col">订单提交人</th>
    <th scope="col">提交日期</th>
    <th scope="col">订单状态</th>
  </tr>
  {{ range .}}
  <tr>
    <td> {{.OrderId }} </td>
    <td> {{.OrderType }} </td>
    <td> {{.OrderUser}} </td>
    <td> {{.CreateTime}}</td>
    <td> 
    <select name="status" id="status">
  <option value="order_submmit">已提交</option>
  <option value="order_finish">已结款</option>
</select>
</td>
  </tr>
  {{ end }}
</table>
</body>
</html>
