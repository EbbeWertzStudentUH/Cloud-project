const mysql = require("mysql2/promise");
const { v4: uuidv4 } = require("uuid");
const { DB_HOST, DB_DATABASE, DB_USER, DB_PWD, SENSITIVE_ENDPOINT_API_SECRET } =
	process.env;

const pool = mysql.createPool({
	host: DB_HOST,
	user: DB_USER,
	password: DB_PWD,
	database: DB_DATABASE,
});

const generalResolvers = {
	user: async ({ id }) => {
		const [rows] = await pool.query(
			"SELECT id, first_name, last_name FROM users WHERE id = ?",
			[id]
		);
		if (rows.length === 0) {
			throw new Error("User not found");
		}
		console.log(rows);
		return rows[0];
	},
	users: async ({ ids }) => {
		const meerdereVraagtekens = ids.map(() => "?").join(", ");
		const [rows] = await pool.query(
			`SELECT id, first_name, last_name
			FROM users
			WHERE id IN (${meerdereVraagtekens});
			`,
			ids
		);
		console.log(rows);
		return rows;
	},
	friends: async ({ id }) => {
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
			FROM friends f
			JOIN users u ON u.id = f.friend_id
			WHERE f.user_id = ?;
			`,
			[id]
		);
		console.log(rows);
		return rows;
	},
	friendRequests: async ({ id }) => {
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
			FROM friend_requests f
			JOIN users u ON u.id = f.friend_id
			WHERE f.user_id = ?;
			`,
			[id]
		);
		console.log(rows);
		return rows;
	},
	addFriendRequest: async ({ user_id, friend_id }) => {
		await pool.query(
			"INSERT INTO friend_requests (user_id, friend_id) VALUES (?, ?)",
			[user_id, friend_id]
		);
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
			FROM friend_requests f
			JOIN users u ON u.id = f.friend_id
			WHERE f.user_id = ?;
			`,
			[user_id]
		);
		return rows;
	},
	addFriend: async ({ user_id, friend_id }) => {
		await pool.query(
			"INSERT INTO friends (user_id, friend_id) VALUES (?, ?)",
			[user_id, friend_id]
		);
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
			FROM friends f
			JOIN users u ON u.id = f.friend_id
			WHERE f.user_id = ?;
			`,
			[user_id]
		);
		return rows;
	},
	removeFriendRequest: async ({ user_id, friend_id }) => {
		await pool.query(
			"DELETE FROM friend_requests WHERE user_id = ? AND friend_id = ?",
			[user_id, friend_id]
		);
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
		  FROM friend_requests f
		  JOIN users u ON u.id = f.friend_id
		  WHERE f.user_id = ?;`,
			[user_id]
		);
		return rows;
	},
	removeFriend: async ({ user_id, friend_id }) => {
		await pool.query(
			"DELETE FROM friends WHERE user_id = ? AND friend_id = ?",
			[user_id, friend_id]
		);
		const [rows] = await pool.query(
			`SELECT u.id, u.first_name, u.last_name
		  FROM friends f
		  JOIN users u ON u.id = f.friend_id
		  WHERE f.user_id = ?;`,
			[user_id]
		);
		return rows;
	},
};

const sensitiveResolvers = {
	user: async ({ id, email }) => {
		let id_field = {};
		if (id) {
			id_field = { name: "id", val: id };
		} else if (email) {
			id_field = { name: "email", val: email };
		} else {
			throw new Error(
				"Provide either 'id' or 'email' to fetch user info"
			);
		}
		console.log(id_field);
		const [rows] = await pool.query(
			`SELECT id, first_name, last_name, password_hash, email FROM users WHERE ${id_field.name} = ?`,
			[id_field.val]
		);
		if (rows.length === 0) {
			throw new Error("User not found");
		}
		return rows[0];
	},
	createUser: async ({ first_name, last_name, password_hash, email }) => {
		const id = uuidv4(); // niet 'uuid()' in de sql query want dan heb ik die hier niet voor de select query
		await pool.query(
			"INSERT INTO users (id, first_name, last_name, password_hash, email) VALUES (?, ?, ?, ?, ?)",
			[id, first_name, last_name, password_hash, email]
		);
		const [newUser] = await pool.query(
			"SELECT id, first_name, last_name, email, password_hash FROM users WHERE id = ?",
			[id]
		);
		return newUser[0];
	},
};

const secret_auth = (req, res, next) => {
	const authHeader = req.headers["authorization"];
	if (
		!authHeader ||
		authHeader !== `Bearer ${SENSITIVE_ENDPOINT_API_SECRET}`
	) {
		return res.status(403).json({ message: "Forbidden" });
	}
	next();
};

module.exports = { generalResolvers, sensitiveResolvers, secret_auth };
