"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("app/customers/page",{

/***/ "(app-pages-browser)/./app/components/tables/SimpleTable.tsx":
/*!***********************************************!*\
  !*** ./app/components/tables/SimpleTable.tsx ***!
  \***********************************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/jsx-dev-runtime.js\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-OA6OURRG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GEJVU65N.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-DRZNIHMG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GIQFRSD6.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-MGVPL3OH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-J4QO5GAJ.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-T2WCTPDH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-B5H2YLEF.mjs\");\n/* harmony import */ var _iconify_react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @iconify/react */ \"(app-pages-browser)/./node_modules/@iconify/react/dist/iconify.mjs\");\n/* __next_internal_client_entry_do_not_use__ default auto */ \n\n\nconst SimpleTable = (param)=>{\n    let { columns, data, size = \"md\", showEdit = false, showDelete = false } = param;\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__.TableContainer, {\n        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__.Table, {\n            size: size,\n            children: [\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__.Thead, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                        children: [\n                            columns.map((col, index)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                    children: col.name\n                                }, index, false, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 38,\n                                    columnNumber: 15\n                                }, undefined)),\n                            (showEdit || showDelete) && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {}, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 44,\n                                columnNumber: 15\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 36,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 35,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__.Tbody, {\n                    children: data.map((dataVal, dataIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                            children: [\n                                columns.map((col, colIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__.Td, {\n                                        children: dataVal[col.key]\n                                    }, colIndex, false, {\n                                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                        lineNumber: 54,\n                                        columnNumber: 17\n                                    }, undefined)),\n                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__.Td, {\n                                    children: [\n                                        showEdit && showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                                            className: \"flex\",\n                                            children: [\n                                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                    icon: \"lucide:edit\"\n                                                }, void 0, false, {\n                                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                    lineNumber: 63,\n                                                    columnNumber: 21\n                                                }, undefined),\n                                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                    icon: \"wpf:delete\"\n                                                }, void 0, false, {\n                                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                    lineNumber: 64,\n                                                    columnNumber: 21\n                                                }, undefined)\n                                            ]\n                                        }, void 0, true, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 62,\n                                            columnNumber: 19\n                                        }, undefined),\n                                        showEdit && !showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                            icon: \"lucide:edit\"\n                                        }, void 0, false, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 69,\n                                            columnNumber: 21\n                                        }, undefined),\n                                        !showEdit && showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                            icon: \"wpf:delete\"\n                                        }, void 0, false, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 73,\n                                            columnNumber: 21\n                                        }, undefined)\n                                    ]\n                                }, void 0, true, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 60,\n                                    columnNumber: 15\n                                }, undefined)\n                            ]\n                        }, dataIndex, true, {\n                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                            lineNumber: 51,\n                            columnNumber: 13\n                        }, undefined))\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 48,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__.Tfoot, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                        children: [\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                children: \"To convert\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 84,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                children: \"into\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 85,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                isNumeric: true,\n                                children: \"multiply by\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 86,\n                                columnNumber: 13\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 83,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 82,\n                    columnNumber: 9\n                }, undefined)\n            ]\n        }, void 0, true, {\n            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n            lineNumber: 34,\n            columnNumber: 7\n        }, undefined)\n    }, void 0, false, {\n        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n        lineNumber: 33,\n        columnNumber: 5\n    }, undefined);\n};\n_c = SimpleTable;\n/* harmony default export */ __webpack_exports__[\"default\"] = (SimpleTable);\nvar _c;\n$RefreshReg$(_c, \"SimpleTable\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevSignature = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevSignature) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports signature on update so we can compare the boundary\n                // signatures. We avoid saving exports themselves since it causes memory leaks (https://github.com/vercel/next.js/pull/53797)\n                module.hot.dispose(function (data) {\n                    data.prevSignature =\n                        self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports);\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevSignature !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevSignature, self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports))) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevSignature !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKGFwcC1wYWdlcy1icm93c2VyKS8uL2FwcC9jb21wb25lbnRzL3RhYmxlcy9TaW1wbGVUYWJsZS50c3giLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7O0FBWXlCO0FBQ2E7QUFVdEMsTUFBTVMsY0FBMEM7UUFBQyxFQUMvQ0MsT0FBTyxFQUNQQyxJQUFJLEVBQ0pDLE9BQU8sSUFBSSxFQUNYQyxXQUFXLEtBQUssRUFDaEJDLGFBQWEsS0FBSyxFQUNuQjtJQUVDLHFCQUNFLDhEQUFDUCw0REFBY0E7a0JBQ2IsNEVBQUNQLG1EQUFLQTtZQUFDWSxNQUFNQTs7OEJBQ1gsOERBQUNYLG1EQUFLQTs4QkFDSiw0RUFBQ0csZ0RBQUVBOzs0QkFDQU0sUUFBUUssR0FBRyxDQUFDLENBQUNDLEtBQXdCQyxzQkFDcEMsOERBQUNaLGdEQUFFQTs4Q0FDQVcsSUFBSUUsSUFBSTttQ0FERkQ7Ozs7OzRCQUtUSixDQUFBQSxZQUFZQyxVQUFTLG1CQUNyQiw4REFBQ1QsZ0RBQUVBOzs7Ozs7Ozs7Ozs7Ozs7OzhCQUlULDhEQUFDSCxtREFBS0E7OEJBQ0hTLEtBQUtJLEdBQUcsQ0FBQyxDQUFDSSxTQUFjQywwQkFFdkIsOERBQUNoQixnREFBRUE7O2dDQUNBTSxRQUFRSyxHQUFHLENBQUMsQ0FBQ0MsS0FBd0JLLHlCQUVwQyw4REFBQ2YsZ0RBQUVBO2tEQUNBYSxPQUFPLENBQUNILElBQUlNLEdBQUcsQ0FBQzt1Q0FEVkQ7Ozs7OzhDQU1YLDhEQUFDZixnREFBRUE7O3dDQUNBTyxZQUFZQyw0QkFDWCw4REFBQ1M7NENBQUlDLFdBQVU7OzhEQUNiLDhEQUFDaEIsZ0RBQUlBO29EQUFDaUIsTUFBSzs7Ozs7OzhEQUNYLDhEQUFDakIsZ0RBQUlBO29EQUFDaUIsTUFBSzs7Ozs7Ozs7Ozs7O3dDQUlkWixZQUFZLENBQUNDLDRCQUNWLDhEQUFDTixnREFBSUE7NENBQUNpQixNQUFLOzs7Ozs7d0NBR2QsQ0FBQ1osWUFBWUMsNEJBQ1YsOERBQUNOLGdEQUFJQTs0Q0FBQ2lCLE1BQUs7Ozs7Ozs7Ozs7Ozs7MkJBdEJWTDs7Ozs7Ozs7Ozs4QkErQmIsOERBQUNqQixtREFBS0E7OEJBQ0osNEVBQUNDLGdEQUFFQTs7MENBQ0QsOERBQUNDLGdEQUFFQTswQ0FBQzs7Ozs7OzBDQUNKLDhEQUFDQSxnREFBRUE7MENBQUM7Ozs7OzswQ0FDSiw4REFBQ0EsZ0RBQUVBO2dDQUFDcUIsU0FBUzswQ0FBQzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQU0xQjtLQXBFTWpCO0FBc0VOLCtEQUFlQSxXQUFXQSxFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwcC9jb21wb25lbnRzL3RhYmxlcy9TaW1wbGVUYWJsZS50c3g/YWI0ZSJdLCJzb3VyY2VzQ29udGVudCI6WyIndXNlIGNsaWVudCc7XG5cbmltcG9ydCB7IFNpbXBsZVRhYmxlQ29sdW1uIH0gZnJvbSAnLi9TaW1wbGVUYWJsZS50eXBlcyc7XG5pbXBvcnQge1xuICBUYWJsZSxcbiAgVGhlYWQsXG4gIFRib2R5LFxuICBUZm9vdCxcbiAgVHIsXG4gIFRoLFxuICBUZCxcbiAgVGFibGVDb250YWluZXIsXG59IGZyb20gJ0BjaGFrcmEtdWkvcmVhY3QnXG5pbXBvcnQgeyBJY29uIH0gZnJvbSAnQGljb25pZnkvcmVhY3QnO1xuXG5pbnRlcmZhY2UgU2ltcGxlVGFibGVQcm9wcyB7XG4gIGNvbHVtbnM6IFNpbXBsZVRhYmxlQ29sdW1uW107XG4gIGRhdGE6IGFueVtdO1xuICBzaXplPzogc3RyaW5nO1xuICBzaG93RWRpdD86IGJvb2xlYW47XG4gIHNob3dEZWxldGU/OiBib29sZWFuO1xufVxuXG5jb25zdCBTaW1wbGVUYWJsZTogUmVhY3QuRkM8U2ltcGxlVGFibGVQcm9wcz4gPSAoe1xuICBjb2x1bW5zLFxuICBkYXRhLFxuICBzaXplID0gJ21kJyxcbiAgc2hvd0VkaXQgPSBmYWxzZSxcbiAgc2hvd0RlbGV0ZSA9IGZhbHNlXG59KSA9PiB7XG5cbiAgcmV0dXJuIChcbiAgICA8VGFibGVDb250YWluZXI+XG4gICAgICA8VGFibGUgc2l6ZT17c2l6ZX0+IFxuICAgICAgICA8VGhlYWQ+XG4gICAgICAgICAgPFRyPlxuICAgICAgICAgICAge2NvbHVtbnMubWFwKChjb2w6IFNpbXBsZVRhYmxlQ29sdW1uLCBpbmRleDogbnVtYmVyKSA9PiAoXG4gICAgICAgICAgICAgIDxUaCBrZXk9e2luZGV4fT5cbiAgICAgICAgICAgICAgICB7Y29sLm5hbWV9XG4gICAgICAgICAgICAgIDwvVGg+XG4gICAgICAgICAgICApKX1cblxuICAgICAgICAgICAgeyhzaG93RWRpdCB8fCBzaG93RGVsZXRlKSAmJlxuICAgICAgICAgICAgICA8VGg+PC9UaD5cbiAgICAgICAgICAgIH1cbiAgICAgICAgICA8L1RyPlxuICAgICAgICA8L1RoZWFkPlxuICAgICAgICA8VGJvZHk+XG4gICAgICAgICAge2RhdGEubWFwKChkYXRhVmFsOiBhbnksIGRhdGFJbmRleDogbnVtYmVyKSA9PiAoXG5cbiAgICAgICAgICAgIDxUciBrZXk9e2RhdGFJbmRleH0+XG4gICAgICAgICAgICAgIHtjb2x1bW5zLm1hcCgoY29sOiBTaW1wbGVUYWJsZUNvbHVtbiwgY29sSW5kZXg6IG51bWJlcikgPT4gKFxuXG4gICAgICAgICAgICAgICAgPFRkIGtleT17Y29sSW5kZXh9PlxuICAgICAgICAgICAgICAgICAge2RhdGFWYWxbY29sLmtleV19XG4gICAgICAgICAgICAgICAgPC9UZD5cblxuICAgICAgICAgICAgICApKX1cbiAgICAgICAgICAgICAgXG4gICAgICAgICAgICAgIDxUZD5cbiAgICAgICAgICAgICAgICB7c2hvd0VkaXQgJiYgc2hvd0RlbGV0ZSAmJiAoXG4gICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT0nZmxleCc+XG4gICAgICAgICAgICAgICAgICAgIDxJY29uIGljb249XCJsdWNpZGU6ZWRpdFwiIC8+ICAgICAgICAgICAgICAgICAgIFxuICAgICAgICAgICAgICAgICAgICA8SWNvbiBpY29uPVwid3BmOmRlbGV0ZVwiIC8+XG4gICAgICAgICAgICAgICAgICA8L2Rpdj5cbiAgICAgICAgICAgICAgICApfVxuXG4gICAgICAgICAgICAgICAge3Nob3dFZGl0ICYmICFzaG93RGVsZXRlICYmIChcbiAgICAgICAgICAgICAgICAgICAgPEljb24gaWNvbj1cImx1Y2lkZTplZGl0XCIgLz4gICAgICAgICAgICAgICAgICAgXG4gICAgICAgICAgICAgICAgKX1cblxuICAgICAgICAgICAgICAgIHshc2hvd0VkaXQgJiYgc2hvd0RlbGV0ZSAmJiAoXG4gICAgICAgICAgICAgICAgICAgIDxJY29uIGljb249XCJ3cGY6ZGVsZXRlXCIgLz5cbiAgICAgICAgICAgICAgICApfVxuICAgICAgICAgICAgICA8L1RkPlxuXG4gICAgICAgICAgICA8L1RyPlxuXG4gICAgICAgICAgKSl9XG5cbiAgICAgICAgPC9UYm9keT5cbiAgICAgICAgPFRmb290PlxuICAgICAgICAgIDxUcj5cbiAgICAgICAgICAgIDxUaD5UbyBjb252ZXJ0PC9UaD5cbiAgICAgICAgICAgIDxUaD5pbnRvPC9UaD5cbiAgICAgICAgICAgIDxUaCBpc051bWVyaWM+bXVsdGlwbHkgYnk8L1RoPlxuICAgICAgICAgIDwvVHI+XG4gICAgICAgIDwvVGZvb3Q+XG4gICAgICA8L1RhYmxlPlxuICAgIDwvVGFibGVDb250YWluZXI+XG4gICk7XG59XG5cbmV4cG9ydCBkZWZhdWx0IFNpbXBsZVRhYmxlO1xuIl0sIm5hbWVzIjpbIlRhYmxlIiwiVGhlYWQiLCJUYm9keSIsIlRmb290IiwiVHIiLCJUaCIsIlRkIiwiVGFibGVDb250YWluZXIiLCJJY29uIiwiU2ltcGxlVGFibGUiLCJjb2x1bW5zIiwiZGF0YSIsInNpemUiLCJzaG93RWRpdCIsInNob3dEZWxldGUiLCJtYXAiLCJjb2wiLCJpbmRleCIsIm5hbWUiLCJkYXRhVmFsIiwiZGF0YUluZGV4IiwiY29sSW5kZXgiLCJrZXkiLCJkaXYiLCJjbGFzc05hbWUiLCJpY29uIiwiaXNOdW1lcmljIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///(app-pages-browser)/./app/components/tables/SimpleTable.tsx\n"));

/***/ })

});