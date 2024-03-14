package query

const GetOrders = `
SELECT o.id, o.order_date, o.ship_date, o.postal_code,
cus.id, cus.name, sg.id, sg.name,
sm.id, sm.name,
ct.id, ct.name, st.id, st.name, co.id, co.name,
rg.id, rg.name
FROM orders o
INNER JOIN cities ct ON o.city_id = ct.id
INNER JOIN states st ON ct.state_id = st.id
INNER JOIN countries co ON co.id = st.country_id
INNER JOIN regions rg ON o.region_id = rg.id
INNER JOIN customers cus ON o.customer_id = cus.id
INNER JOIN segments sg ON cus.segment_id = sg.id
INNER JOIN ship_modes sm ON o.ship_mode_id = sm.id;
`

const GetOrderDetails = `
SELECT
od.id, od.sales, od.quantity, od.discount, od.profit ,od.order_id,
pd.id, pd.name,
sc.id, sc.name,
ctg.id, ctg.name
FROM order_details od
INNER JOIN products pd ON od.product_id = pd.id
INNER JOIN sub_categories sc ON pd.sub_category_id = sc.id
INNER JOIN categories ctg ON sc.category_id = ctg.id;
`

// const GetOrders = `
// SELECT o.id, o.order_date, o.ship_date, o.postal_code,
// cus.id, cus.name, sg.id, sg.name,
// sm.id, sm.name,
// ct.id, ct.name, st.id, st.name, co.id, co.name,
// rg.id, rg.name,
// od.id, od.sales, od.quantity, od.discount, od.profit,
// pd.id, pd.name,
// sc.id, sc.name,
// ctg.id, ctg.name
// FROM orders o
// INNER JOIN cities ct ON o.city_id = ct.id
// INNER JOIN states st ON ct.state_id = st.id
// INNER JOIN countries co ON co.id = st.country_id
// INNER JOIN regions rg ON o.region_id = rg.id
// INNER JOIN customers cus ON o.customer_id = cus.id
// INNER JOIN segments sg ON cus.segment_id = sg.id
// INNER JOIN ship_modes sm ON o.ship_mode_id = sm.id
// INNER JOIN order_details od ON o.id = od.order_id
// INNER JOIN products pd ON od.product_id = pd.id
// INNER JOIN sub_categories sc ON pd.sub_category_id = sc.id
// INNER JOIN categories ctg ON sc.category_id = ctg.id;
// `

const GET_SAME_USERNAME_OR_EMAIL = `
	select count(*) as jumlah FROM users
	WHERE username=? OR email=?;
`
const INSERT_USER = `
	INSERT INTO users(id,username, password, email, first_name, last_name, role_id)
	VALUES(?,?,?,?,?,?,?);
`

const GET_USER_BY_USERNAME_AND_PASSWORD = `
	SELECT id, username, email, first_name, last_name, role_id FROM users
	WHERE username = ? AND password = ?;
`

const GET_USER_BY_ID = `
	SELECT id, username, email, first_name, last_name, role_id FROM users
	WHERE id = ?;
`

const GET_SALES_SUM = `
select sum(od.sales) as sales,month(o.order_date) as month,year(o.order_date) as year from orders o
inner join order_details od on od.order_id = o.id
group by month(o.order_date), year(o.order_date)
order by year(o.order_date) desc, month(o.order_date) desc
limit ?;`

const GET_TOTAL_PRODUCT = `
select count(*) as sum from products;
`
const GET_MOST_BOUGHT_CATEGORY = `
select c.name from order_details od 
inner join products p on p.id = od.product_id 
inner join sub_categories sc on sc.id = p.sub_category_id 
inner join categories c on c.id = sc.category_id
group by c.name
order by count(od.id) desc limit 1;`

const GET_TOP_TRANSACTION = `
select c.id as customer_id, p.name as product_name, o.order_date ,od.sales from order_details od 
inner join orders o on o.id = od.order_id 
inner join products p on p.id = od.product_id 
inner join customers c on c.id = o.customer_id
order by sales desc
limit ?;`

const GET_ROLE_ID = `
SELECT id FROM roles
WHERE name=?;
`

const GET_CODE = `
SELECT code,role_id FROM codes
WHERE code=? AND email=?;
`

const INVITE_ADMIN = `
INSERT INTO codes(email,role_id,code) VALUES (?,?,?);`

const GET_SALES_DATA = `
SELECT * FROM (
	select month(o.order_date) as 'Order Month', year(o.order_date) as 'Order Year', p.numeric_id as 'Product ID', sum(od.sales) as Sales  from orders o
	inner join order_details od on od.order_id = o.id 
	inner join products p on p.id = od.product_id
	group by p.numeric_id,year(o.order_date) , month(o.order_date)
	) as query
	where query.` + "`Order Month`" + ` = ? and query.` + "`Order Year`" + ` = ? and query.` + "`Product ID`" + ` = ?
	order by query.` + "`Order Month`" + `;
`
const GET_PRODUCT_BY_ID = `
	SELECT p.id, p.name, p.numeric_id, p.sub_category_id, sc.name ,c.id ,c.name  FROM products p
	INNER join sub_categories sc ON sc.id = p.sub_category_id 
	INNER join categories c ON c.id = sc.category_id 
	WHERE numeric_id = ?;
`
const GET_PRODUCT_SUMMARY = `
	SELECT
	count(*) OVER(),
	pd.id, pd.name,
	sc.name,
	ctg.name,
	sum(od.sales) as total_sales
	FROM order_details od
	INNER JOIN products pd ON od.product_id = pd.id
	INNER JOIN sub_categories sc ON pd.sub_category_id = sc.id
	INNER JOIN categories ctg ON sc.category_id = ctg.id
	group by product_id
	LIMIT ? OFFSET ?;
`

const GET_PRODUCTS = `
	SELECT numeric_id, name FROM products;
`
