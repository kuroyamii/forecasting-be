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
