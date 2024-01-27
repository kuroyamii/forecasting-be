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
	INSERT INTO users(id,username, password, email, full_name)
	VALUES(?,?,?,?,?);
`

const GET_USER_BY_USERNAME_AND_PASSWORD = `
	SELECT id, username, email, full_name FROM users
	WHERE username = ? AND password = ?;
`

const GET_USER_BY_ID = `
	SELECT id, username, email, full_name FROM users
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
