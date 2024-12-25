const mysql = require("mysql2/promise");
const { DB_HOST, DB_DATABASE, DB_USER, DB_PWD, AUTH_SECRET } = process.env;

const pool = mysql.createPool({
	host: DB_HOST,
	user: DB_USER,
	password: DB_PWD,
	database: DB_DATABASE,
});

const generalResolvers = {
	getUser: async ({ id }) => {
		const [rows] = await pool.query(
			"SELECT id, first_name, last_name FROM users WHERE id = ?",
			[id]
		);
		if (rows.length === 0) {
			throw new Error("User not found");
		}
		return rows[0];
	},
};

const sensitiveResolvers = {
	getAuthInfo: async ({ id }) => {
		const [rows] = await pool.query(
			"SELECT id, password_hash FROM users WHERE id = ?",
			[id]
		);
		if (rows.length === 0) {
			throw new Error("User not found");
		}
		return rows[0];
	},
};

const secret_auth = (req, res, next) => {
	const authHeader = req.headers["authorization"];
	if (!authHeader || authHeader !== `Bearer ${AUTH_SECRET}`) {
		return res.status(403).json({ message: "Forbidden" });
	}
	next();
};

module.exports = { generalResolvers, sensitiveResolvers, secret_auth };
