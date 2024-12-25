const mysql = require("mysql2/promise");
const { v4: uuidv4 } = require('uuid');
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
    createUser: async ({ first_name, last_name, password_hash }) => {
        const id = uuidv4(); // niet 'uuid()' in de sql query want dan heb ik die hier niet voor de select query
        await pool.query(
          'INSERT INTO users (id, first_name, last_name, password_hash) VALUES (?, ?, ?, ?)',
          [id, first_name, last_name, password_hash]
        );
        const [newUser] = await pool.query('SELECT id, first_name, last_name FROM users WHERE id = ?', [id]);
        return newUser[0];
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
