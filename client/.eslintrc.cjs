/** @type {import("@types/eslint").Linter.BaseConfig} */
module.exports = {
	root: true,
	extends: [
		"eslint:recommended",
		"plugin:@typescript-eslint/recommended",
		"plugin:svelte/recommended",
		"prettier",
	],
	parser: "@typescript-eslint/parser",
	plugins: ["@typescript-eslint", "simple-import-sort"],
	parserOptions: {
		sourceType: "module",
		ecmaVersion: 2020,
		extraFileExtensions: [".svelte"],
		warnOnUnsupportedTypeScriptVersion: false,
	},
	env: {
		browser: true,
		es2017: true,
		node: true,
	},
	rules: {
		"@typescript-eslint/ban-ts-comment": "off",
		"@typescript-eslint/consistent-type-imports": "error",
		"@typescript-eslint/no-unused-vars": [
			"warn",
			{ ignoreRestSiblings: true, varsIgnorePattern: "^(_|\\$\\$)", argsIgnorePattern: "^_" },
		],
		"@typescript-eslint/no-explicit-any": "off",
		"no-fallthrough": "off",
		"no-undef": "off",
		"no-unused-vars": "off",
		"no-extra-boolean-cast": "off",
		// Doesn't play nice with using stores...
		"svelte/valid-compile": "off",
		"simple-import-sort/imports": "warn",
		"simple-import-sort/exports": "warn",
	},
	overrides: [
		{
			files: ["*.svelte"],
			parser: "svelte-eslint-parser",
			parserOptions: {
				parser: "@typescript-eslint/parser",
			},
		},
	],
};
