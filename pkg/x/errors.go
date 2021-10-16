package x

// Error messages.
var Errors = map[string]string{
	`file_not_x`:                    `this is not x source file: `,
	`invalid_token`:                 `undefined code content`,
	`invalid_syntax`:                `invalid syntax`,
	`no_entry_point`:                `entry point (main) function is not defined`,
	`exist_name`:                    `name is already exist`,
	`brace_not_closed`:              `brace is opened but not closed`,
	`function_body_not_exist`:       `function body is not declared`,
	`incompatible_type`:             `incompatible value type`,
	`operator_overflow`:             `operator overflow`,
	`incompatible_datatype`:         `data type are not compatible`,
	`operator_notfor_string`:        `this operator is not defined for string type`,
	`operator_notfor_rune`:          `this operator is not defined for rune type`,
	`operator_notfor_null`:          `this operator is not defined for null type`,
	`operator_notfor_bool`:          `this operator is not defined for boolean type`,
	`operator_notfor_any`:           `this operator is not defined for any type`,
	`operator_notfor_float`:         `this operator is not defined for float type(s)`,
	`operator_notfor_int`:           `this operator is not defined for integer type(s)`,
	`operator_notfor_uint`:          `this operator is not defined for unsigned integer type(s)`,
	`operator_notfor_pointer`:       `this operator is not defined for pointer type`,
	`name_not_defined`:              `name is not defined`,
	`not_function_call`:             `value is not function`,
	`parameter_exist`:               `parameter is already exist in this name`,
	`argument_overflow`:             `argument overflow`,
	`entrypoint_have_return`:        `entry point is cannot have return type`,
	`entrypoint_have_parameters`:    `entry point is cannot have parameter(s)`,
	`require_return_value`:          `return statements of non-void functions should have return value`,
	`void_function_return_value`:    `void functions is cannot returns any value`,
	`bitshift_must_unsigned`:        `bit shifting value is must be unsigned`,
	`logical_not_bool`:              `logical expression is have only boolean type values`,
	`const_value_update`:            `value is cannot update of constants`,
	`type_not_support_value_update`: `type is not support value update`,
	`invalid_type`:                  `invalid data type`,
	`invalid_attribute`:             `invalid attribute for type`,
	`invalid_numeric_range`:         `arithmetic value overflow`,
	`invalid_data_unary`:            `invalid data type for unary operator`,
	`invalid_operator`:              `invalid operator`,
	`invalid_data_plus`:             `invalid data type for plus operator`,
	`invalid_data_tilde`:            `invalid data type for tilde operator`,
	`invalid_data_logical_not`:      `invalid data type for logical not operator`,
	`invalid_data_star`:             `invalid data type for star operator`,
	`invalid_data_amper`:            `invalid data type for amper operator`,
	`invalid_escape_sequence`:       `invalid escape sequence`,
	`invalid_const_data_type`:       `invalid data type for constant`,
	`invalid_type_source`:           `invalid data type source for type alias`,
	`missing_autotype_value`:        `auto-type declarations should have a initializer`,
	`missing_type`:                  `data type missing`,
	`missing_value`:                 `value is not given`,
	`missing_argument`:              `missing argument(s)`,
	`missing_block_comment`:         `missing block comment close`,
	`missing_semicolon`:             `missing statement terminator at end of statement`,
	`missing_rune_end`:              `rune is not finished`,
	`missing_return`:                `missing return at end of function`,
	`missing_string_end`:            `string is not finished`,
	`missing_const_value`:           `constants must have value specification`,
	`null_for_autotype`:             `null is cannot use with auto-type definations`,
	`void_for_autotype`:             `void data is cannot use for auto-type definations`,
	`rune_empty`:                    `rune is cannot empty`,
	`rune_overflow`:                 `rune is should be single`,
}
